package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var bookings []Spot

func main() {
	r := mux.NewRouter()

	// Create Mock data
	bookings = append(bookings, Spot{ID: "1", BookedDate: time.Date(2021, 9, 4, 0, 0, 0, 0, time.Now().Location()), BookingStatus: "NB", BookedOn: time.Now(), BookedBy: "123123"})

	initializeRoutes(r)
	log.Fatal(http.ListenAndServe(":8000", r))
}

func initializeRoutes(r *mux.Router) {
	// General / Admin routes
	r.HandleFunc("/api/bookings", createBooking).Methods("POST")

	// Me / Employee routes
	r.HandleFunc("/api/me/bookings", getBookings).Methods("GET")
	r.HandleFunc("/api/me/bookings/{id}", getBooking).Methods("GET")
	r.HandleFunc("/api/me/bookings", createBooking).Methods("POST")
}

func getBookings(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, bookings)
}

func getBooking(w http.ResponseWriter, r *http.Request) {
	if !isAuthorized(*r) {
		respondWithError(w, http.StatusUnauthorized, "You are not authorized to book a spot.")
		return
	}

	// Extract the id from the url
	vars := mux.Vars(r)
	id, ok := vars["id"]

	i := findBooking(bookings, id)
	if !ok || i == -1 {
		respondWithError(w, http.StatusNotFound, "The booking was not found.")
		return
	}
	respondWithJSON(w, http.StatusOK, bookings[i])
}

func createBooking(w http.ResponseWriter, r *http.Request) {
	if !isAuthorized(*r) {
		respondWithError(w, http.StatusUnauthorized, "You are not authorized to book a spot.")
		return
	}

	const MAX_BOOKINGS = 15
	if len(bookings) >= MAX_BOOKINGS {
		respondWithError(w, http.StatusServiceUnavailable, "There are no spots available.")
		return
	}

	var newSpot Spot
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&newSpot); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	bookings = append(bookings, newSpot)
	respondWithJSON(w, http.StatusCreated, newSpot)
}
