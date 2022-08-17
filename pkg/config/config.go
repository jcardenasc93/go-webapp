package config

import (
	"html/template"
	"log"
)

// AppConfig holds the app config
type AppConfig struct {
	Port          string
	TemplateCache map[string]*template.Template
	UseCacheTemplates      bool
	InfoLog       *log.Logger
}
