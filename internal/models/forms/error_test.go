package forms

import (
	"net/url"
	"testing"
)

func TestFormErrors_Get(t *testing.T) {
	key, value := "field", "test"
	data := url.Values{}
	form := New(data)
	form.data.Add(key, value)

	form.MinLenght(key, 6)
	e := form.Errors.Get(key)
	if e == "" {
		t.Error("Expected error but got empty one")
	}
}
