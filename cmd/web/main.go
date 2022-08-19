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
	// Initilize app based on AppConfig struct
	app, err := initApp()
	if err != nil {
		return
	}

	// Session management
	sessionMan = initSession()
	// Allows access session from any handler
	app.Session = sessionMan
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
	app.IsProduction = false
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

func initSession() *scs.SessionManager {
	sm := scs.New()
	// Set life time
	sm.Lifetime = 12 * time.Hour
	// Persist session even when user leaves site
	sm.Cookie.Persist = true
	sm.Cookie.SameSite = http.SameSiteLaxMode
	// Securing allows to encrypt session but only through https so on dev env doesn't apply
	sm.Cookie.Secure = app.IsProduction

	return sm
}
