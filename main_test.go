package main

import (
	"context"

	"github.com/cucumber/godog"
)

func InitializeTestSuite(sc *godog.TestSuiteContext) {
	sc.BeforeSuite(func() {
		bookings = []Spot{}
	})
}

// Godog Scenarios
func InitializeScenario(sc *godog.ScenarioContext) {
	sc.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		bookings = []Spot{}
		return ctx, nil
	})

	// Get spots
	sc.Step(`^there is (\d+) booking$`, thereIsBooking)
	sc.Step(`^user is verified as "([^"]*)"$`, userIsVerifiedAs)
	sc.Step(`^a "([^"]*)" request is sent to the endpoint "([^"]*)"$`, aRequestIsSentToTheEndpoint)
	sc.Step(`^the HTTP-response code should be "(\d+)"$`, theHTTPresponseCodeShouldBe)
	sc.Step(`^the response should have a list of (\d+) objects$`, theResponseShouldHaveAListOfObjects)

	// Set normal spot
	/* ctx.Step(`^Add the spot object in the database and success message is shown$`, saveReservationAndShowSuccessMessage)
	ctx.Step(`^An error message is shown\. "([^"]*)"\.$`, anErrorMessageIsShown)
	ctx.Step(`^Employee tries to book the spot$`, bookSpot)
	ctx.Step(`^User input is not validated as employee$`, userIsNotValidatedAsEmployee)
	ctx.Step(`^User input is validated as an employee or admin and office spots are available$`, spotAvailable)
	ctx.Step(`^User input is validated as employee and office reservation spots are unavailable$`, spotUnavailable)
	ctx.Step(`^User tries to book the spot$`, bookSpot) */
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
