package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var a App

func TestMain(m *testing.M) {
	a.Initialize()
	code := m.Run()
	os.Exit(code)
}

func TestGetBook(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/books/1", nil)
	res := executeRequest(req)
	checkResponseCode(t, http.StatusOK, res.Code)

	var result Book
	json.Unmarshal(res.Body.Bytes(), &result)

	if result.ID != "1" {
		t.Errorf("Expected an two books in an array. Got %v", result.ID)
	}
	if result.Isbn != "58899" {
		t.Errorf("Expected an two books in an array. Got %v", result.Isbn)
	}
	if result.Title != "Book One" {
		t.Errorf("Expected an two books in an array. Got %v", result.Title)
	}
	if result.Author.Firstname != "Frank" {
		t.Errorf("Expected an two books in an array. Got %v", result.Author.Firstname)
	}
	if result.Author.Lastname != "Dorito" {
		t.Errorf("Expected an two books in an array. Got %v", result.Author.Lastname)
	}
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}
