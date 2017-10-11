package zoom

import (
	"bytes"
	"encoding/json"
	"testing"
	"time"
)

func TestNullTimePointer(t *testing.T) {
	type structure struct {
		Time *Time `json:"time"`
	}
	var (
		js       = []byte(`{ "time": null }`)
		parsed   = structure{}
		expected *Time
	)

	if err := json.Unmarshal(js, &parsed); err != nil {
		t.Error(err)
	}

	actual := parsed.Time

	if actual != expected {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}

func TestNullTime(t *testing.T) {
	type structure struct {
		Time Time `json:"time"`
	}
	var (
		js       = []byte(`{ "time": null }`)
		parsed   = structure{}
		expected = time.Time{}
	)

	if err := json.Unmarshal(js, &parsed); err != nil {
		t.Error(err)
	}

	actual := parsed.Time.Time

	if actual != expected {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}

func TestEmptyStringTime(t *testing.T) {
	type structure struct {
		Time Time `json:"time"`
	}
	var (
		js       = []byte(`{ "time": "" }`)
		parsed   = structure{}
		expected = time.Time{}
	)

	if err := json.Unmarshal(js, &parsed); err != nil {
		t.Error(err)
	}

	actual := parsed.Time.Time

	if actual != expected {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}

func TestInvalidTime(t *testing.T) {
	type structure struct {
		Time Time `json:"time"`
	}
	var (
		js     = []byte(`{ "time": "bogus" }`)
		parsed = structure{}
	)

	if err := json.Unmarshal(js, &parsed); err == nil {
		t.Error("expected error parsing bogus time, got nil")
	}
}

func TestValidTime(t *testing.T) {
	type structure struct {
		Time Time `json:"time"`
	}
	var (
		js       = []byte(`{ "time": "` + nowStr + `" }`)
		parsed   = structure{}
		expected = now.Time
	)

	if err := json.Unmarshal(js, &parsed); err != nil {
		t.Error(err)
	}

	actual := parsed.Time.Time

	if !actual.Equal(expected) {
		t.Errorf("expected %+v, got %+v", expected, actual)
	}
}

func TestTimeToString(t *testing.T) {
	type structure struct {
		Time Time `json:"time"`
	}
	var (
		struc    = structure{Time{now.Time}}
		expected = []byte(`{"time":"` + nowStr + `"}`)
	)

	actual, err := json.Marshal(struc)
	if err != nil {
		t.Errorf("got err marshaling json: %s\n", err)
	}

	if bytes.Compare(actual, expected) != 0 {
		t.Errorf("expected %s, got %s", string(expected), string(actual))
	}
}
