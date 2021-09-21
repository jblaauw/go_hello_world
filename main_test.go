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
		// Variable from main application
		bookings = []Spot{}

		// Variables from get_spots_test
		rrBookings = nil
		req = nil

		// Variables from set_spots_test
		requestBody = nil
		return ctx, nil
	})

	// Get spots
	sc.Step(`^a "([^"]*)" request is created for the endpoint "([^"]*)"$`, aRequestIsCreatedForTheEndpoint)
	sc.Step(`^the user is verified as "([^"]*)"$`, theUserIsVerifiedAs)
	sc.Step(`^there is (\d+) booking$`, thereIsBooking)
	sc.Step(`^the request is sent to the endpoint$`, theRequestIsSentToTheEndpoint)
	sc.Step(`^the HTTP-response code should be "([^"]*)"$`, theHTTPresponseCodeShouldBe)
	sc.Step(`^the response should have a list of (\d+) objects$`, theResponseShouldHaveAListOfObjects)

	// Set normal spot
	sc.Step(`^a "([^"]*)" office spot is "([^"]*)"$`, aOfficeSpotIs)
	sc.Step(`^the request header "([^"]*)" is set to "([^"]*)"$`, theRequestHeaderIsSetTo)
	sc.Step(`^the request body contains a new "([^"]*)" booking$`, theRequestBodyContainsANewBooking)
	sc.Step(`^response body should contain "([^"]*)" as it\'s "([^"]*)"$`, responseBodyShouldContainAsIts)
}
