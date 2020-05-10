package main

import (
	"net/http"
	"net/http/httptest"
	"io/ioutil"
	"testing"
)

func TestCounterValue(t *testing.T) {
	expected := "100"
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	c := Counter{}
	handler := http.HandlerFunc(c.CountHandler)
	for i := 0; i < 99; i++ {
		handler.ServeHTTP(httptest.NewRecorder(), req)
	}
	handler.ServeHTTP(rr, req)
	body, err := ioutil.ReadAll(rr.Body)
	if err != nil {
		t.Errorf("can't read body")
	}
	if string(body) != expected {
		t.Errorf("handler returned wrong count value")
	}
}
