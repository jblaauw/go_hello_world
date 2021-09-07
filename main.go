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

	// Routes
	r.HandleFunc("/api/bookings", getBookings).Methods("GET")

	// Listen for incoming requests + error logging
	log.Fatal(http.ListenAndServe(":8000", r))
}

func getBookings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bookings)
}
