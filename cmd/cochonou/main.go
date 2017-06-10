package main

import (
	"flag"
	std_os "os"

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

	e := echo.New()
	// Root level middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	logEntriesToken := std_os.Getenv("LOGENTRIES_TOKEN")
	if logEntriesToken != "" {
		writer, err := le_go.Connect(logEntriesToken)
		if err != nil {
			panic(err)
		}

		e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
			Output: writer,
		}))
	}

	e.POST("/redirections", redirHandler.HandleCreate)
	e.GET("/redirections", redirHandler.HandleGetList)
	e.Logger.Fatal(e.Start(os.GetEnvWithDefault("HTTP_ADDR", ":9494")))
}
