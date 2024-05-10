package main

import (
	"flag"
	"fmt"
	"go-breeders/configuration"
	"html/template"
	"log"
	"net/http"
	"time"
)

const port = ":4000"

type application struct {
	templateMap map[string]*template.Template
	config      appConfig
	App         *configuration.Application
	catService  *RemoteService
}

type appConfig struct {
	useCache bool
	dsn      string // Data Source Name
}

func main() {
	app := application{
		templateMap: make(map[string]*template.Template),
	}

	flag.BoolVar(&app.config.useCache, "cache", false, "Use template cache")
	flag.StringVar(&app.config.dsn, "dsn", "mariadb:myverysecretpassword@tcp(localhost:3306)/breeders?parseTime=true&tls=false&collation=utf8_unicode_ci&timeout=5s", "DSN")
	flag.Parse()

	// get database
	db, err := initMySQLDB(app.config.dsn)
	if err != nil {
		log.Panic(err)
	}

	// jsonBackend := &JSONBackend{}
	// jsonAdapter := &RemoteService{Remote: jsonBackend}

	XMLBackend := &XMLBackend{}
	xmlAdapter := &RemoteService{Remote: XMLBackend}

	app.App = configuration.New(db)
	app.catService = xmlAdapter

	srv := &http.Server{
		Addr:              port,
		Handler:           app.routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       30 * time.Second,
		ReadHeaderTimeout: 30 * time.Second,
		WriteTimeout:      30 * time.Second,
	}

	fmt.Println("Starting web application on port", port)
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
