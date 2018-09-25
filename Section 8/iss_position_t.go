package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetISSLocation(t *testing.T) {
	req, err := http.NewRequest("GET", "http://api.open-notify.org/iss-now.json", nil)
	req.Header.Set("Content-Type", "application/json2")

	if err != nil {
		t.Fatal(err)
	}

	resp := httptest.NewRecorder()
	http.DefaultServeMux.ServeHTTP(resp, req)

	// test Content-Type, StatusCode
	if req.Header.Get("Content-Type") != "application/json" {
		t.Fatalf("Received non JSON response: %s\n", req.Header.Get("Content-Type"))

	}

	_, err = ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Fail()
	}
}
