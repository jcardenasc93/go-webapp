package config

import (
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
)

// AppConfig holds the app config
type AppConfig struct {
	Port              string
	TemplateCache     map[string]*template.Template
	UseCacheTemplates bool
	InfoLog           *log.Logger
	IsProduction      bool
	Session           *scs.SessionManager
	SameSite          http.SameSite
}

// InitApp acts like contructor
func (app *AppConfig) InitApp(serverPort string, sessionMan *scs.SessionManager) {
	app.Port = serverPort
	// Allows access session from any handler
	app.Session = sessionMan
	app.IsProduction = false
	app.SameSite = http.SameSiteLaxMode
	// Securing allows to encrypt session but only through https so on dev env doesn't apply
	app.Session.Cookie.Secure = app.IsProduction
	app.Session.Cookie.SameSite = app.SameSite
}

func InitSession(sessionMan *scs.SessionManager) *scs.SessionManager {
	sessionMan = scs.New()
	// Set life time
	sessionMan.Lifetime = 12 * time.Hour
	// Persist session even when user leaves site
	sessionMan.Cookie.Persist = true
	return sessionMan
}
