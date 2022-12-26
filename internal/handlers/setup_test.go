package handlers

import (
	"encoding/gob"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jcardenasc93/go-webapp/internal/config"
	"github.com/jcardenasc93/go-webapp/internal/models"
	"github.com/jcardenasc93/go-webapp/internal/render"
	"github.com/justinas/nosurf"
)

var app config.AppConfig
var sessionMan *scs.SessionManager
var functions = template.FuncMap{}

func getRoutes() http.Handler {
	// Includes complex types to store in session
	gob.Register(models.Reservation{})

	// Session management
	sessionMan = config.InitSession(sessionMan)
	// Initilize app
	app.InitApp("", sessionMan)

	// Setup templates config
	tmplCache, err := CreateTestTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache:", err)
	} else {
		app.TemplateCache = tmplCache
		// In a dev env is ok disable cache
		app.UseCacheTemplates = true
		// Setup templates
		render.SetupTemplates(&app)
	}

	// Setup config
	repo := NewRepo(&app)
	SetupHandlers(repo)

	// Get routes
	mux := chi.NewRouter()
	// Adds middlewares
	mux.Use(middleware.Recoverer)
	mux.Use(LoadSession)

	// Adds routes
	mux.Get("/", Repo.Home)
	mux.Get("/about", Repo.About)
	mux.Get("/contact", Repo.Contact)
	mux.Get("/majestic-suite", Repo.Majestic)
	mux.Get("/comfortable-place", Repo.Comfortable)

	mux.Post("/make-reservation", Repo.PostMakeReservation)
	mux.Get("/make-reservation", Repo.MakeReservation)

	mux.Get("/booking", Repo.Booking)
	mux.Post("/booking", Repo.PostBooking)
	mux.Post("/booking-json", Repo.BookingJSON)

	mux.Get("/booking-summary", Repo.BookingSummary)

	// File server
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))
	return mux
}

// GenCSRFToken creates middleware to generate CSRF token to secure POST requests
func GenCSRFToken(next http.Handler) http.Handler {
	tokenHandler := nosurf.New(next)
	tokenHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.IsProduction, // dev env not uses https
		SameSite: app.SameSite,
	})
	return tokenHandler
}

// LoadSession load and save session for each request
func LoadSession(next http.Handler) http.Handler {
	return sessionMan.LoadAndSave(next)
}

func CreateTestTemplateCache() (map[string]*template.Template, error) {
	templateCache := map[string]*template.Template{}
	templsDir := "./../../templates"

	// Looks for all *pages.tmpl files in ./templates
	pages, err := filepath.Glob(templsDir + "/*.page.tmpl")
	if err != nil {
		log.Println("Error scaning templates directory:\n", err)
		return templateCache, err
	}
	for _, page := range pages {
		name := filepath.Base(page)
		tmplSetup, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			log.Println("Error scaning templates directory:\n", err)
			return templateCache, err
		}

		// Look for layouts
		matches, err := filepath.Glob(templsDir + "/*.layout.tmpl")
		if err != nil {
			log.Println("Error scaning templates directory:\n", err)
			return templateCache, err
		}
		if len(matches) > 0 {
			// Match each page template with required layouts
			tmplSetup, err = tmplSetup.ParseGlob(templsDir + "/*.layout.tmpl")
			if err != nil {
				log.Println("Error scaning templates directory:\n", err)
				return templateCache, err
			}
		}
		templateCache[name] = tmplSetup
	}

	return templateCache, nil
}
