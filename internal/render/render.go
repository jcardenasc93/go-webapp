package render

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/jcardenasc93/go-webapp/internal/config"
	"github.com/jcardenasc93/go-webapp/internal/models"
	"github.com/justinas/nosurf"
)

var app *config.AppConfig

// SetupTemplates setup templates package based on AppConfig
func SetupTemplates(a *config.AppConfig) {
	app = a
}

// AddDefaultTempData adds default data to template data
func AddDefaultTempData(td *models.TemplateData, r *http.Request) *models.TemplateData {
	td.CSRFToken = nosurf.Token(r)
	return td
}

// RenderTemplate parses a given template in the ResponseWriter
func RenderTemplate(w http.ResponseWriter, r *http.Request, tmplName string, td *models.TemplateData) {
	var tmplCache map[string]*template.Template
	if app.UseCacheTemplates {
		// Get template cache from app config
		tmplCache = app.TemplateCache

	} else {
		// Create new instance on each call (read template from disk)
		tmplCache, _ = CreateTemplateCache()
	}

	// get template
	tmpl, ok := tmplCache[tmplName]
	if !ok {
		log.Fatal("Error accesing to cache map with the given key")
	}

	td = AddDefaultTempData(td, r)
	// render template
	// NOTE: here td is like context object passed in Django templates
	err := tmpl.Execute(w, td)
	if err != nil {
		log.Fatal(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	templateCache := map[string]*template.Template{}
	templsDir := "./templates"

	// Looks for all *pages.tmpl files in ./templates
	pages, err := filepath.Glob(templsDir + "/*.page.tmpl")
	if err != nil {
		log.Println("Error scaning templates directory:\n", err)
		return templateCache, err
	}
	for _, page := range pages {
		name := filepath.Base(page)
		tmplSetup, err := template.New(name).ParseFiles(page)
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
