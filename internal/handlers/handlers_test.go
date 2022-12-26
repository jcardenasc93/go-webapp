package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type postData struct {
	key   string
	value string
}

var tests = []struct {
	name               string
	url                string
	method             string
	params             []postData
	expectedStatusCode int
}{
	{
		"home",
		"/",
		"GET",
		[]postData{},
		http.StatusOK,
	},
	{
		"about",
		"/about",
		"GET",
		[]postData{},
		http.StatusOK,
	},
	{
		"majestic",
		"/majestic-suite",
		"GET",
		[]postData{},
		http.StatusOK,
	},
	{
		"comfortable",
		"/comfortable-place",
		"GET",
		[]postData{},
		http.StatusOK,
	},
	{
		"make-reservation",
		"/make-reservation",
		"GET",
		[]postData{},
		http.StatusOK,
	},
	{
		"booking",
		"/booking",
		"GET",
		[]postData{},
		http.StatusOK,
	},
	{
		"booking-summary",
		"/booking-summary",
		"GET",
		[]postData{},
		http.StatusOK,
	},
	{
		"contact",
		"/contact",
		"GET",
		[]postData{},
		http.StatusOK,
	},
	{
		"post-make-reservation",
		"/make-reservation",
		"POST",
		[]postData{
			{key: "first_name", value: "Juan"},
			{key: "last_name", value: "Smith"},
			{key: "email", value: "juan.smith@email.com"},
			{key: "phone", value: "+1555555555"},
		},
		http.StatusOK,
	},
	{
		"post-booking",
		"/booking",
		"POST",
		[]postData{
			{key: "start", value: "2022-01-01"},
			{key: "end", value: "2022-01-05"},
		},
		http.StatusOK,
	},
	{
		"post-booking-json",
		"/booking-json",
		"POST",
		[]postData{
			{key: "start", value: "2022-01-01"},
			{key: "end", value: "2022-01-05"},
		},
		http.StatusOK,
	},
}

func TestHandlers(t *testing.T) {
	routes := getRoutes()
	ts := httptest.NewTLSServer(routes)
	defer ts.Close()

	for _, tc := range tests {
		if tc.method == "GET" {
			resp, err := ts.Client().Get(ts.URL + tc.url)
			checkError(&err, t)
			if resp.StatusCode != tc.expectedStatusCode {
				t.Errorf("Error running %s:%s.\nExpected %d but got %d", tc.method, tc.url, tc.expectedStatusCode, resp.StatusCode)
			}
		} else {
			data := url.Values{}
			for _, p := range tc.params {
				data.Add(p.key, p.value)
			}
			resp, err := ts.Client().PostForm(ts.URL+tc.url, data)
			checkError(&err, t)
			if resp.StatusCode != tc.expectedStatusCode {
				t.Errorf("Error running %s:%s.\nExpected %d but got %d", tc.method, tc.url, tc.expectedStatusCode, resp.StatusCode)
			}
		}
	}
}

func checkError(err *error, t *testing.T) {
	if *err != nil {
		t.Fatal(err)
	}
}
