package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestApplication_GetAllDogBreedsJSON(t *testing.T) {
	// create a request
	req, _ := http.NewRequest("GET", "/api/dog-breeds", nil)
	// create a response
	rr := httptest.NewRecorder()
	// create a handler
	handler := http.HandlerFunc(testApp.GetAllDogBreeds)
	// serve a handler
	handler.ServeHTTP(rr, req)

	// check response
	if rr.Code != http.StatusOK {
		t.Errorf("Wrong response code %d, wanted 200", rr.Code)
	}
}

func TestApplication_GetAllCatBreeds(t *testing.T) {
	// create a request
	req, _ := http.NewRequest("GET", "/api/cat-breeds", nil)
	// create a response
	rr := httptest.NewRecorder()
	// create a handler
	handler := http.HandlerFunc(testApp.GetAllCatBreeds)
	// serve a handler
	handler.ServeHTTP(rr, req)

	// check response
	if rr.Code != http.StatusOK {
		t.Errorf("Wrong response code %d, wanted 200", rr.Code)
	}
}
