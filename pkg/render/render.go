package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

//NOTE: pkg var to help cache templates instead of read files from disk on each request
var templatesCache = make(map[string]*template.Template)

// This method parse a given template in the ResponseWriter
func RenderTemplate(w http.ResponseWriter, tm string) {
	var t *template.Template
	var err error
	// Look for template in cache var
	_, inCache := templatesCache[tm]
	if inCache {
		// Template exists
		log.Printf("Using cached template: %s", tm)
	} else {
		// Adds to cache var
		err = addTemplateCache(tm)
	}
	t = templatesCache[tm]
	err = t.Execute(w, nil)
	if err != nil {
		log.Printf("Error parsing template %s:\n", tm)
		log.Println(err)
	}
}

func addTemplateCache(t string) error {
	tmplFiles := []string{
		fmt.Sprintf("./templates/%s", t),
		fmt.Sprint("./templates/base.layout.tmpl"),
	}
	parsedTemplate, err := template.ParseFiles(tmplFiles...)
	if err != nil {
		return err
	}
	// Add parsed template to cache var
	templatesCache[t] = parsedTemplate
	return nil
}
