package forms

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/asaskevich/govalidator"
)

type Form struct {
	data   url.Values
	Errors formErrors
}

// New initializes a Form struct
func New(data url.Values) *Form {
	return &Form{
		data:   data,
		Errors: formErrors(map[string][]string{}),
	}
}

// Required iters over all given fields and checks that their values are not empty
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.data.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field is required")
		}
	}
}

// Has checks if given form field is not empty
func (f *Form) Has(field string) bool {
	v := f.data.Get(field)
	if v == "" {
		return false
	}
	return true
}

// MinLenght checks that given field value has the minimum given chars
func (f *Form) MinLenght(field string, minL int) {
	v := f.data.Get(field)
	if len(strings.TrimSpace(v)) < minL {
		f.Errors.Add(field, fmt.Sprintf("Value should be %v at least", minL))
	}
}

func (f *Form) IsEmail(field string) {
	if !govalidator.IsEmail(f.data.Get(field)) {
		f.Errors.Add(field, "Invalid email address")
	}
}

// Valid returns true if no errors found in form
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}
