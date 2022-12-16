package models

// TemplateData type holds necesary mappings for different types that will be usefull when send data to templates
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	Float     map[string]float32
	Complex   map[string]interface{}
	CSRFToken string
	FlashMsg  string
	ErrorMsg  string
	Warning   string
}
