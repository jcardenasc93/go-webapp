package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/alexedwards/scs/v2"
	"github.com/jcardenasc93/go-webapp/pkg/config"
	"github.com/jcardenasc93/go-webapp/pkg/handlers"
	"github.com/jcardenasc93/go-webapp/pkg/render"
)

const serverPort = ":8000"

var app config.AppConfig
var sessionMan *scs.SessionManager

func main() {
	// Session management
	sessionMan = config.InitSession(sessionMan)
	// Initilize app
	app.InitApp(serverPort, sessionMan)

	setupTemplates(&app)
	setupHandlers(&app)

	log.Println(fmt.Sprintf("Starting server on %s", serverPort))

	//NOTE: Here starts the server
	srv := &http.Server{
		Addr:    app.Port,
		Handler: routing(&app),
	}
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func setupTemplates(app *config.AppConfig) {
	// Setup templates config
	tmplCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache: ", err)
	} else {
		app.TemplateCache = tmplCache
		// In a dev env is ok disable cache
		app.UseCacheTemplates = false
		// Setup templates
		render.SetupTemplates(app)
	}
}

func setupHandlers(app *config.AppConfig) {
	// Setup handlers config
	repo := handlers.NewRepo(app)
	handlers.SetupHandlers(repo)
}
