package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"time"
)

var rrBookings *httptest.ResponseRecorder

func aRequestIsSentToTheEndpoint(method, endpoint string) error {
	req, _ := http.NewRequest(method, endpoint, nil)

	rrBookings = httptest.NewRecorder()
	handler := http.HandlerFunc(getBookings)
	handler.ServeHTTP(rrBookings, req)

	return nil
}

func theHTTPresponseCodeShouldBe(expectedCode int) error {
	if status := rrBookings.Code; status != expectedCode {
		return fmt.Errorf("Expected status code %d. Got %d", expectedCode, status)
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

func thereIsBooking(bookingAmount int) error {
	bookings = append(bookings, Spot{ID: "1", BookedDate: time.Date(2021, 9, 4, 0, 0, 0, 0, time.Now().Location()), BookingStatus: "NB", BookedOn: time.Now(), BookedBy: "123123"})
	return nil
}

func userIsVerifiedAs(expectedRole string) error {
	// TODO: implement token verification
	return nil
}
