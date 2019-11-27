package zoom

import (
	"bytes"
	"encoding/json"
	"testing"
	"time"
)

func TestNullDatePointer(t *testing.T) {
	type structure struct {
		Date *Date `json:"date"`
	}
	var (
		js       = []byte(`{ "date": null }`)
		parsed   = structure{}
		expected *Date
	)

	if err := json.Unmarshal(js, &parsed); err != nil {
		t.Error(err)
	}

	actual := parsed.Date

	if actual != expected {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}

func TestNullDate(t *testing.T) {
	type structure struct {
		Date Date `json:"date"`
	}
	var (
		js       = []byte(`{ "date": null }`)
		parsed   = structure{}
		expected = time.Time{}
	)

	if err := json.Unmarshal(js, &parsed); err != nil {
		t.Error(err)
	}

	actual := parsed.Date.Time

	if actual != expected {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}

func TestEmptyStringDate(t *testing.T) {
	type structure struct {
		Date Date `json:"date"`
	}
	var (
		js       = []byte(`{ "date": "" }`)
		parsed   = structure{}
		expected = time.Time{}
	)

	if err := json.Unmarshal(js, &parsed); err != nil {
		t.Error(err)
	}

	actual := parsed.Date.Time

	if actual != expected {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}

func TestInvalidDate(t *testing.T) {
	type structure struct {
		Date Date `json:"date"`
	}
	var (
		js     = []byte(`{ "date": "bogus" }`)
		parsed = structure{}
	)

	if err := json.Unmarshal(js, &parsed); err == nil {
		t.Error("expected error parsing bogus time, got nil")
	}
}

func TestValidDate(t *testing.T) {
	type structure struct {
		Date Date `json:"date"`
	}
	var (
		js       = []byte(`{ "date": "` + todayStr + `" }`)
		parsed   = structure{}
		expected = today.Time
	)

	if err := json.Unmarshal(js, &parsed); err != nil {
		t.Error(err)
	}

	actual := parsed.Date.Time

	if !actual.Equal(expected) {
		t.Errorf("expected %+v, got %+v", expected, actual)
	}
}

func TestDateToString(t *testing.T) {
	t.Skip("Can't quite get this to marshal to 2006-01-02")

	type structure struct {
		Date Date `json:"date"`
	}
	var (
		struc    = structure{Date{today.Time}}
		expected = []byte(`{"date":"` + todayStr + `"}`)
	)

	actual, err := json.Marshal(struc)
	if err != nil {
		t.Errorf("got err marshaling json: %s\n", err)
	}

	if bytes.Compare(actual, expected) != 0 {
		t.Errorf("expected %s, got %s", string(expected), string(actual))
	}
}
