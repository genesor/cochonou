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

	store := &bolt.RedirectionStore{
		DB: db,
	}
	if store != nil {

	}

	redirHandler := &http.RedirectionHandler{}
	e.POST("/redirections", redirHandler.HandleCreate)
	e.Logger.Fatal(e.Start(os.GetEnvWithDefault("HTTP_ADDR", ":9494")))
}
