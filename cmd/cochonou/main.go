package main

import (
	"context"
	"flag"
	std_os "os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/asdine/storm"
	"github.com/bsphere/le_go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	go_ovh "github.com/ovh/go-ovh/ovh"
	"github.com/subosito/gotenv"

	"github.com/genesor/cochonou"
	"github.com/genesor/cochonou/bolt"
	"github.com/genesor/cochonou/http"
	"github.com/genesor/cochonou/os"
	"github.com/genesor/cochonou/ovh"
)

func main() {
	isDev := flag.Bool("dev", false, "Run the service in developer mode")
	flag.Parse()

	// Load env variables from dev.env file in dev mode.
	if *isDev == true {
		gotenv.Load("dev.env")
	}

	db, err := storm.Open(os.GetEnvWithDefault("BOLT_DB_PATH", "cochonou_dev.db"))
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	store := &bolt.RedirectionStore{
		DB: db,
	}

	ovhClient, err := go_ovh.NewClient(
		os.MustGetEnv("OVH_ENDPOINT"),
		os.MustGetEnv("OVH_APP_KEY"),
		os.MustGetEnv("OVH_APP_SECRET"),
		os.MustGetEnv("OVH_CONSUMER_KEY"),
	)
	OVHDomainHandler := ovh.NewDomainHandler(
		os.MustGetEnv("COCHONOU_DOMAIN"),
		ovhClient,
	)
	storedDomainHandler := &cochonou.StoredDomainHandler{
		DomainHandler: OVHDomainHandler,
		Store:         store,
	}

	redirHandler := &http.RedirectionHandler{
		DomainHandler: storedDomainHandler,
		Store:         store,
	}

	servers := []*echo.Echo{}
	api := echo.New()
	web := echo.New()

	api.POST("/redirections", redirHandler.HandleCreate)
	api.GET("/redirections", redirHandler.HandleGetList)

	api.Server.Addr = os.GetEnvWithDefault("API_HTTP_ADDR", ":9494")
	servers = append(servers, api)

	web.Server.Addr = os.GetEnvWithDefault("WEB_HTTP_ADDR", ":9393")
	servers = append(servers, web)

	// Config and setup servers
	var wg sync.WaitGroup
	for _, server := range servers {
		// Add middlewares to servers
		server.Use(middleware.Logger())
		server.Use(middleware.Recover())

		logEntriesToken := std_os.Getenv("LOGENTRIES_TOKEN")
		if logEntriesToken != "" {
			writer, err := le_go.Connect(logEntriesToken)
			if err != nil {
				panic(err)
			}

			api.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
				Output: writer,
			}))
		}

		wg.Add(1)
		go func(server *echo.Echo) {
			defer wg.Done()
			server.Logger.Fatal(server.StartServer(server.Server))
		}(server)
	}

	c := make(chan std_os.Signal, 1)
	signal.Notify(c, std_os.Interrupt, syscall.SIGTERM)
	go func() {
		for _ = range c {
			for _, server := range servers {
				defer wg.Done()
				wg.Add(1)
				// Wait for interrupt signal to gracefully shutdown the server with
				// a timeout of 5 seconds.
				ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
				defer cancel()
				server.Shutdown(ctx)
			}
		}
	}()

	wg.Wait()
}

func startServer(e *echo.Echo) {
	e.Logger.Fatal(e.StartServer(e.Server))
}
