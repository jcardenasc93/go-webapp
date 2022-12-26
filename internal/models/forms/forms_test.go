package forms

import (
	"net/url"
	"testing"
)

func TestForm_Valid(t *testing.T) {
	f := url.Values{}
	form := New(f)

	isValid := form.Valid()
	if !isValid {
		t.Error("Form validation error. Form is valid but checked as invalid")
	}
}

func TestForm_MinLengh(t *testing.T) {
	key, value := "field", "test"
	data := url.Values{}
	form := New(data)
	form.data.Add(key, value)

	form.MinLenght(key, 3)
	if !form.Valid() {
		t.Errorf("This form is valid, but got these error:\n%v", form.Errors)
	}

	form.MinLenght(key, 8)
	if form.Valid() {
		t.Error("This form is invalid, but got no errors")
	}
}

func TestForm_Has(t *testing.T) {
	key, value := "field", "test"
	data := url.Values{}
	form := New(data)
	form.data.Add(key, value)

	fieldExists := form.Has(key)
	if !fieldExists {
		t.Errorf("Field %s exists in form with value %s", key, value)
	}

	fieldExists = form.Has("no-existing")
	if fieldExists {
		t.Errorf("Field %s doesn't exist in form", "no-existing")
	}
}

func TestForm_IsEmail(t *testing.T) {
	key, value := "email", "test@valid.com"
	data := url.Values{}
	form := New(data)
	form.data.Add(key, value)

	form.IsEmail(key)
	if !form.Valid() {
		t.Errorf("This form is valid, but got these error:\n%v", form.Errors)
	}

	form.data.Set(key, "test.invalid.com")
	form.IsEmail(key)
	if form.Valid() {
		t.Error("This form has invalid email, but got no errors")
	}
}

func TestForm_Required(t *testing.T) {
	data := url.Values{}
	form := New(data)

	form.Required("a", "b", "c")
	if form.Valid() {
		t.Error("This form is invalid. But got no errors")
	}

	formData := url.Values{}
	formData.Add("a", "1")
	formData.Add("b", "2")
	formData.Add("c", "3")

	form = New(formData)
	form.Required("a", "b", "c")

	if !form.Valid() {
		t.Errorf("This form is valid. But got these errors:\n%v", form.Errors)
	}
}
