package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"time"

	"github.com/gorilla/mux"
)

var rrBookings *httptest.ResponseRecorder
var req *http.Request

func aRequestIsCreatedForTheEndpoint(method, endpoint string) error {
	var _ error
	if requestBody == nil {
		req, _ = http.NewRequest(method, endpoint, nil)
	} else {
		req, _ = http.NewRequest(method, endpoint, requestBody)
	}

	return nil
}

func theUserIsVerifiedAs(expectedRole string) error {
	// TODO: implement token verification
	req.Header.Set("Authorization", "Bearer RsT5OjbzRn430zqMLgV3Ia")
	return nil
}

func thereIsBooking(bookingAmount int) error {
	bookings = append(bookings, Spot{ID: "1", BookedDate: time.Date(2021, 9, 4, 0, 0, 0, 0, time.Now().Location()), BookingStatus: "NB", BookedOn: time.Now(), BookedBy: "123123"})
	return nil
}

func theRequestIsSentToTheEndpoint() error {
	rrBookings = httptest.NewRecorder()
	r := mux.NewRouter()

	// TODO: instead of rebuilding the routes, use the existing ones in main.go
	r.HandleFunc("/api/bookings", createBooking).Methods("POST")
	r.HandleFunc("/api/me/bookings", getBookings).Methods("GET")
	r.HandleFunc("/api/me/bookings", createBooking).Methods("POST")
	r.ServeHTTP(rrBookings, req)

	return nil
}

func theHTTPresponseCodeShouldBe(expectedMsg string) error {
	var statusMsg string
	switch rrBookings.Code {
	case 200:
		statusMsg = "success"
	case 201:
		statusMsg = "created"
	case 401:
		statusMsg = "unauthorized"
	case 503:
		statusMsg = "service unavailable"
	default:
		statusMsg = "failed"
	}

	if statusMsg != expectedMsg {
		return fmt.Errorf("Expected status code %s. Got %s", expectedMsg, statusMsg)
	}
	return nil
}

func theResponseShouldHaveAListOfObjects(expectedCount int) error {
	jsonBookings := []map[string]string{}
	if err := json.Unmarshal(rrBookings.Body.Bytes(), &jsonBookings); err != nil {
		return fmt.Errorf("Could not convert http response to map. Got: %s", rrBookings.Body.String())
	}

	if len(jsonBookings) != expectedCount {
		return fmt.Errorf("Expected the following object count: %d. Got %d", expectedCount, len(jsonBookings))
	}
	return nil
}
