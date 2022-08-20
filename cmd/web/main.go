package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

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
	initSession()
	// Initilize app based on AppConfig struct
	initApp()

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

func initApp() {
	app.Port = serverPort
	// Allows access session from any handler
	app.Session = sessionMan
	app.IsProduction = false
	// Setup templates config
	tmplCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache: ", err)
	} else {
		app.TemplateCache = tmplCache
		// In a dev env is ok disable cache
		app.UseCacheTemplates = false
		// Setup handlers config
		repo := handlers.NewRepo(&app)
		handlers.SetupHandlers(repo)
		// Setup templates
		render.SetupTemplates(&app)
	}
}

func initSession() {
	sessionMan = scs.New()
	// Set life time
	sessionMan.Lifetime = 12 * time.Hour
	// Persist session even when user leaves site
	sessionMan.Cookie.Persist = true
	sessionMan.Cookie.SameSite = http.SameSiteLaxMode
	// Securing allows to encrypt session but only through https so on dev env doesn't apply
	sessionMan.Cookie.Secure = app.IsProduction
}
