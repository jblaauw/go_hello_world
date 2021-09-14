package main

import "github.com/cucumber/godog"

func aOfficeSpotIs(spotType, availability string) error {
	// TODO: how to check and make it so that there is a spot available / unavailable?
	return godog.ErrPending
}

func booksASpot(user, spotType string) error {
	return godog.ErrPending
}

func saveTheBooking() error {
	return godog.ErrPending
}

func anMessageIsShown(msgType, msg string) error {
	return godog.ErrPending
}
