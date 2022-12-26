package render

import (
	"errors"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/jcardenasc93/go-webapp/internal/config"
	"github.com/jcardenasc93/go-webapp/internal/models"
	"github.com/justinas/nosurf"
)

var app *config.AppConfig
var functions = template.FuncMap{}
var templsDir = "./templates"

// SetupTemplates setup templates package based on AppConfig
func SetupTemplates(a *config.AppConfig) {
	app = a
}

// AddDefaultTempData adds default data to template data
func AddDefaultTempData(td *models.TemplateData, r *http.Request) *models.TemplateData {
	// Get user's msg
	td.FlashMsg = app.Session.PopString(r.Context(), "flash")
	td.ErrorMsg = app.Session.PopString(r.Context(), "error")
	td.Warning = app.Session.PopString(r.Context(), "warning")

	td.CSRFToken = nosurf.Token(r)
	return td
}

// RenderTemplate parses a given template in the ResponseWriter
func RenderTemplate(w http.ResponseWriter, r *http.Request, tmplName string, td *models.TemplateData) error {
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
		errMsg := "Error accesing to cache map with the given key"
		return errors.New(errMsg)
	}

	td = AddDefaultTempData(td, r)
	// render template
	// NOTE: here td is like context object passed in Django templates
	err := tmpl.Execute(w, td)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	templateCache := map[string]*template.Template{}

	// Looks for all *pages.tmpl files in ./templates
	pages, err := filepath.Glob(templsDir + "/*.page.tmpl")
	if err != nil {
		return templateCache, err
	}
	for _, page := range pages {
		name := filepath.Base(page)
		tmplSetup, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return templateCache, err
		}

		// Look for layouts
		matches, err := filepath.Glob(templsDir + "/*.layout.tmpl")
		if err != nil {
			return templateCache, err
		}
		if len(matches) > 0 {
			// Match each page template with required layouts
			tmplSetup, err = tmplSetup.ParseGlob(templsDir + "/*.layout.tmpl")
			if err != nil {
				return templateCache, err
			}
		}
		templateCache[name] = tmplSetup
	}

	return templateCache, nil
}
