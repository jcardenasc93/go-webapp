package config

import (
	"html/template"
	"log"

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
}
