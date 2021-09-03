package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cucumber/godog"
)

// Given
func thereIsOneBooking(normalSpots int, emergencySpots int) error {
	// TODO: change this
	normalSpotsAvailable = normalSpots
	emergencySpotsAvailable = emergencySpots
	return nil
}

func userIsVerifiedAsEmployeeOrAdmin() error {
	// TODO: implement
	return nil
}

// When
func employeeVisitsBookingsOverview(t *testing.T) error {
	// TODO: change the request to get a range from now till 1 week
	req, _ := http.NewRequest("GET", "/api/bookings", nil)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getBookings)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Expected response code %v. Got %v\n", http.StatusOK, status)
	}

	expected := ``
	if rr.Body.String() != expected {
		t.Errorf("Expected atleast 1 reserved spot. Got %v", rr.Body.String())
	}
	//json.Unmarshal(res.Body.Bytes(), &result)

	return nil
}

// Then
func showAListOfAllAvailableSpotsForEachDayForUpcomingWeek() error {
	// TODO: compare spots available
	return godog.ErrPending
}
