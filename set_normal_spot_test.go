package main

import "github.com/cucumber/godog"

// Given: spot available
func spotAvailable() error {
	return godog.ErrPending
}

// Given: spot unavailable
func spotUnavailable() error {
	return godog.ErrPending
}

// Given: user not validated as employee
func userIsNotValidatedAsEmployee() error {
	return godog.ErrPending
}

// When: employee books a spot
func bookSpot() error {
	return godog.ErrPending
}

// Then: add spot to db and show successmessage
func saveReservationAndShowSuccessMessage() error {
	return godog.ErrPending
}

// Then: show an error message
func anErrorMessageIsShown(arg1 string) error {
	return godog.ErrPending
}
