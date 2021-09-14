package main

import (
	"context"

	"github.com/cucumber/godog"
)

func InitializeTestSuite(sc *godog.TestSuiteContext) {
	// Runs before entire test Suite
	sc.BeforeSuite(func() {
		bookings = []Spot{}
	})
}

// Godog Scenarios
func InitializeScenario(sc *godog.ScenarioContext) {
	// Runs before every scenario
	sc.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		bookings = []Spot{}
		return ctx, nil
	})

	// Get spots
	sc.Step(`^there is (\d+) booking$`, thereIsBooking)
	sc.Step(`^user is verified as "([^"]*)"$`, userIsVerifiedAs)
	sc.Step(`^a "([^"]*)" request is sent to the endpoint "([^"]*)"$`, aRequestIsSentToTheEndpoint)
	sc.Step(`^the HTTP-response code should be "([^"]*)"$`, theHTTPresponseCodeShouldBe)
	sc.Step(`^the response should have a list of (\d+) objects$`, theResponseShouldHaveAListOfObjects)

	// Set normal spot
	sc.Step(`^a "([^"]*)" office spot is "([^"]*)"$`, aOfficeSpotIs)
	sc.Step(`^"([^"]*)" books a "([^"]*)" spot$`, booksASpot)
	sc.Step(`^save the booking$`, saveTheBooking)
	sc.Step(`^An "([^"]*)" message is shown\. "([^"]*)"$`, anMessageIsShown)

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
