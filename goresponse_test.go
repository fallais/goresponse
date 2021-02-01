package goresponse_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fallais/goresponse"
)

type FakeResponse struct {
	FakeFieldString string `json:"fake_field_string"`
	FakeFieldInt    int    `json:"fake_field_int"`
	FakeFieldBool   bool   `json:"fake_field_bool"`
}

func TestJSON(t *testing.T) {
	fr := FakeResponse{
		FakeFieldString: "yep",
		FakeFieldInt:    1,
		FakeFieldBool:   false,
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		goresponse.JSON(w, http.StatusOK, fr)
	}))
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		t.Error(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		t.Error(err)
	}

	// Check the status code
	if res.StatusCode != http.StatusOK {
		t.Error("wrong expected status code")
	}

	// Check the response body
	expected := `{"fake_field_string":"yep","fake_field_int":1,"fake_field_bool":false}`
	if string(body) != expected {
		t.Errorf("wrong expected body: %s", string(body))
	}
}
