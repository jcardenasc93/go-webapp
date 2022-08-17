package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jcardenasc93/go-webapp/pkg/config"
	"github.com/jcardenasc93/go-webapp/pkg/handlers"
	"github.com/jcardenasc93/go-webapp/pkg/render"
)

const serverPort = ":8000"

func main() {
	// Initilize app based on AppConfig struct
	app, err := initApp()
	if err != nil {
		return
	}
	render.SetupTemplates(&app)

	log.Println(fmt.Sprintf("Starting server on %s", serverPort))

	//NOTE: Here starts the server
	srv := &http.Server{
		Addr:    app.Port,
		Handler: routing(&app),
	}
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func initApp() (config.AppConfig, error) {
	var app config.AppConfig
	app.Port = serverPort
	// Setup templates config
	tmplCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache: ", err)
		return app, err
	} else {
		// Setup handlers config
		repo := handlers.NewRepo(&app)
		handlers.SetupHandlers(repo)
		app.TemplateCache = tmplCache
		// In a dev env is ok disable cache
		app.UseCacheTemplates = false
	}

	return app, err
}
