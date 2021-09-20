package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

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

func findBooking(bookings []Spot, id string) int {
	for i, booking := range bookings {
		if booking.ID == id {
			return i
		}
	}
	return -1
}
