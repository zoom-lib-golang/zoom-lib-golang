package zoom

import (
	"bytes"
	"encoding/json"
	"net/url"
	"testing"
)

func TestNullURLPointer(t *testing.T) {
	type structure struct {
		URL *URL `json:"url"`
	}
	var (
		js       = []byte(`{ "url": null }`)
		parsed   = structure{}
		expected *URL
	)

	if err := json.Unmarshal(js, &parsed); err != nil {
		t.Error(err)
	}

	actual := parsed.URL

	if actual != expected {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}

func TestNullURL(t *testing.T) {
	type structure struct {
		URL URL `json:"url"`
	}
	var (
		js       = []byte(`{ "url": null }`)
		parsed   = structure{}
		expected = &url.URL{}
	)

	if err := json.Unmarshal(js, &parsed); err != nil {
		t.Error(err)
	}

	actual := parsed.URL.URL

	if *actual != *expected {
		t.Errorf("expected %+v, got %+v", *expected, *actual)
	}
}

func TestEmptyStringURL(t *testing.T) {
	type structure struct {
		URL URL `json:"url"`
	}
	var (
		js       = []byte(`{ "url": "" }`)
		parsed   = structure{}
		expected = &url.URL{}
	)

	if err := json.Unmarshal(js, &parsed); err != nil {
		t.Error(err)
	}

	actual := parsed.URL.URL

	if *actual != *expected {
		t.Errorf("expected %+v, got %+v", *expected, *actual)
	}
}

func TestInvalidURL(t *testing.T) {
	type structure struct {
		URL URL `json:"url"`
	}
	var (
		js     = []byte(`{ "url": "f&o://bar" }`)
		parsed = structure{}
	)

	if err := json.Unmarshal(js, &parsed); err == nil {
		t.Errorf("expected error parsing bogus url, got nil: %+v\n", *parsed.URL.URL)
	}
}

func TestValidURL(t *testing.T) {
	type structure struct {
		URL URL `json:"url"`
	}
	var (
		js       = []byte(`{ "url": "` + urlStr + `" }`)
		parsed   = structure{}
		expected = urlT.URL
	)

	if err := json.Unmarshal(js, &parsed); err != nil {
		t.Error(err)
	}

	actual := parsed.URL.URL

	if *actual != *expected {
		t.Errorf("expected %+v, got %+v", *expected, *actual)
	}
}

func TestURLToString(t *testing.T) {
	type structure struct {
		URL *URL `json:"url"`
	}
	var (
		struc    = structure{&urlT}
		expected = []byte(`{"url":"` + urlStr + `"}`)
	)

	actual, err := json.Marshal(struc)
	if err != nil {
		t.Errorf("got err marshaling json: %s\n", err)
	}

	if bytes.Compare(actual, expected) != 0 {
		t.Errorf("expected %s, got %s", string(expected), string(actual))
	}
}
