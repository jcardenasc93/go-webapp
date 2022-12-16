package forms

type formErrors map[string][]string

// Add allows to add a new error with custom message for specific form field
func (fe *formErrors) Add(field string, msg string) {
	(*fe)[field] = append((*fe)[field], msg)
}

// Get retrieves the first error for the given field
func (fe *formErrors) Get(field string) string {
	errorList := (*fe)[field]
	if len(errorList) == 0 {
		return ""
	}
	return errorList[0]
}
