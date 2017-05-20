package main

import (
	"github.com/asdine/storm"
	"github.com/labstack/echo"

	"github.com/genesor/cochonou/bolt"
	"github.com/genesor/cochonou/http"
	"github.com/genesor/cochonou/os"
)

func main() {
	e := echo.New()

	db, err := storm.Open(os.GetEnvWithDefault("BOLT_DB_PATH", "cochonou_dev.db"))
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	// _, _ := db.Begin(true)

	store := &bolt.ImageRedirectionStore{
		DB: db,
	}
	if store != nil {

	}

	helloHandler := http.NewHelloHandler()
	e.GET("/", helloHandler.HandleHello)
	e.Logger.Fatal(e.Start(os.GetEnvWithDefault("HTTP_ADDR", ":9494")))
}
