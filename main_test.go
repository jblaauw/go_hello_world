package main

import (
	"context"

	"github.com/cucumber/godog"
)

func InitializeTestSuite(sc *godog.TestSuiteContext) {
	sc.BeforeSuite(func() {
		normalSpotsAvailable = 0
		emergencySpotsAvailable = 0
	})
}

// Godog Scenarios
func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		// TODO: clean the state before every scenario
		normalSpotsAvailable = 0
		emergencySpotsAvailable = 0

		return ctx, nil
	})

	// Get spots
	ctx.Step(`^there is one booking$`, thereIsOneBooking)
	ctx.Step(`^User is verified as employee or admin$`, userIsVerifiedAsEmployeeOrAdmin)
	ctx.Step(`^Employee visits bookings overview$`, employeeVisitsBookingsOverview)
	ctx.Step(`^Show a list of all available spots for each day for upcoming week$`, showAListOfAllAvailableSpotsForEachDayForUpcomingWeek)

	// Set normal spot
	ctx.Step(`^Add the spot object in the database and success message is shown$`, saveReservationAndShowSuccessMessage)
	ctx.Step(`^An error message is shown\. "([^"]*)"\.$`, anErrorMessageIsShown)
	ctx.Step(`^Employee tries to book the spot$`, bookSpot)
	ctx.Step(`^User input is not validated as employee$`, userIsNotValidatedAsEmployee)
	ctx.Step(`^User input is validated as an employee or admin and office spots are available$`, spotAvailable)
	ctx.Step(`^User input is validated as employee and office reservation spots are unavailable$`, spotUnavailable)
	ctx.Step(`^User tries to book the spot$`, bookSpot)
}

// =========== TODO: remove/replace code below ==============
/* var a App

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


*/
