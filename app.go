package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
}

var books []Book

func (a *App) Initialize() {
	// Create mock data - TODO: impl database
	books = append(books, Book{ID: "1", Isbn: "58899", Title: "Book One", Author: &Author{Firstname: "Frank", Lastname: "Dorito"}})
	books = append(books, Book{ID: "2", Isbn: "58877", Title: "Book Two", Author: &Author{Firstname: "Berend", Lastname: "Spa"}})

	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(":8000", a.Router))
}

// Get all books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// Get single book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	// Loop through books and find with id
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{})
}

// Create a new book
func createBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(1000000)) // Mock, do not use in prod
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

// Update a book
func updateBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	json.NewEncoder(w).Encode(books)
}

// Delete a book
func deleteBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/api/books", getBooks).Methods("GET")
	a.Router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	a.Router.HandleFunc("/api/books", createBooks).Methods("POST")
	a.Router.HandleFunc("/api/books/{id}", updateBooks).Methods("PUT")
	a.Router.HandleFunc("/api/books/{id}", deleteBooks).Methods("DELETE")
}
