package main

import "github.com/cucumber/godog"

func addTheSpotObjectInTheDatabaseAndSuccessMessageIsShown() error {
	return godog.ErrPending
}

func anErrorMessageIsShown(arg1 string) error {
	return godog.ErrPending
}

func employeeTriesToBookTheSpot() error {
	return godog.ErrPending
}

func userInputIsNotValidatedAsEmployee() error {
	return godog.ErrPending
}

func userInputIsValidatedAsAnEmployeeOrAdminAndOfficeSpotsAreAvailable() error {
	return godog.ErrPending
}

func userInputIsValidatedAsEmployeeAndOfficeReservationSpotsAreUnavailable() error {
	return godog.ErrPending
}

func userTriesToBookTheSpot() error {
	return godog.ErrPending
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^Add the spot object in the database and success message is shown$`, addTheSpotObjectInTheDatabaseAndSuccessMessageIsShown)
	ctx.Step(`^An error message is shown\. "([^"]*)"\.$`, anErrorMessageIsShown)
	ctx.Step(`^Employee tries to book the spot$`, employeeTriesToBookTheSpot)
	ctx.Step(`^User input is not validated as employee$`, userInputIsNotValidatedAsEmployee)
	ctx.Step(`^User input is validated as an employee or admin and office spots are available$`, userInputIsValidatedAsAnEmployeeOrAdminAndOfficeSpotsAreAvailable)
	ctx.Step(`^User input is validated as employee and office reservation spots are unavailable$`, userInputIsValidatedAsEmployeeAndOfficeReservationSpotsAreUnavailable)
	ctx.Step(`^User tries to book the spot$`, userTriesToBookTheSpot)
}
