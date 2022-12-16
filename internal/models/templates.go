package models

import "github.com/jcardenasc93/go-webapp/internal/models/forms"

// TemplateData type holds necesary mappings for different types that will be usefull when send data to templates
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	Float     map[string]float32
	Data      map[string]interface{}
	CSRFToken string
	FlashMsg  string
	ErrorMsg  string
	Warning   string
	Form      *forms.Form
}
