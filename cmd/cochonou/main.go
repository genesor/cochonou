package main

import (
	"github.com/asdine/storm"
	"github.com/labstack/echo"
	go_ovh "github.com/ovh/go-ovh/ovh"

	"github.com/genesor/cochonou"
	"github.com/genesor/cochonou/bolt"
	"github.com/genesor/cochonou/http"
	"github.com/genesor/cochonou/os"
	"github.com/genesor/cochonou/ovh"
)

func main() {
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
	}

	e := echo.New()
	e.POST("/redirections", redirHandler.HandleCreate)
	e.Logger.Fatal(e.Start(os.GetEnvWithDefault("HTTP_ADDR", ":9494")))
}
