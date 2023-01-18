package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"

	"github.com/alexedwards/scs/v2"
	"github.com/jcardenasc93/go-webapp/internal/config"
	"github.com/jcardenasc93/go-webapp/internal/config/helpers"
	"github.com/jcardenasc93/go-webapp/internal/handlers"
	"github.com/jcardenasc93/go-webapp/internal/models"
	"github.com/jcardenasc93/go-webapp/internal/render"
)

const serverPort = ":8000"

var app config.AppConfig
var sessionMan *scs.SessionManager

func main() {
	// Includes complex types to store in session
	gob.Register(models.Reservation{})

	err := run()
	if err != nil {
		log.Fatal("Cannot create template cache: ", err)
	}

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

func setupTemplates(app *config.AppConfig) error {
	// Setup templates config
	tmplCache, err := render.CreateTemplateCache()
	if err != nil {
		return err
	} else {
		app.TemplateCache = tmplCache
		// In a dev env is ok disable cache
		app.UseCacheTemplates = false
		// Setup templates
		render.SetupTemplates(app)
	}
	return nil
}

func setupHandlers(app *config.AppConfig) {
	// Setup handlers config
	repo := handlers.NewRepo(app)
	handlers.SetupHandlers(repo)
}

func run() error {
	// Session management
	sessionMan = config.InitSession(sessionMan)
	// Initilize app
	app.InitApp(serverPort, sessionMan)

	err := setupTemplates(&app)
	if err != nil {
		return err
	}
	setupHandlers(&app)
	helpers.NewHelpers(&app)

	return nil
}
