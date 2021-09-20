package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
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
	r.HandleFunc("/api/bookings", createBooking).Methods("POST")

	// Listen for incoming requests + error logging
	log.Fatal(http.ListenAndServe(":8000", r))
}

func getBookings(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, bookings)
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

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func isAuthorized(r http.Request) bool {
	reqToken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer ")
	return len(splitToken) == 2
	// TODO: implement OAuth
	// can get the token like: reqToken = splitToken[1]
}
