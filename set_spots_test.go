package main

import (
	"bytes"
	"encoding/json"
	"time"

	"github.com/cucumber/godog"
)

var spotAvailable bool
var requestBody *bytes.Buffer

func aOfficeSpotIs(spotType, availability string) error {
	// TODO: change implementation if there is a real way to check if there is a spot available
	spotAvailable = availability == "available"
	return nil
}

func theRequestHeaderContentTypeIsSetTo(arg1 string) error {
	return godog.ErrPending
}

func theRequestBodyContainsANewBooking() error {
	newSpot := Spot{
		ID:            "0",
		BookedDate:    time.Date(2021, 9, 4, 0, 0, 0, 0, time.Now().Location()),
		BookingStatus: "NB",
		BookedOn:      time.Now(),
		BookedBy:      "123123",
	}

	newSpotJson, _ := json.Marshal(newSpot)
	requestBody = bytes.NewBuffer(newSpotJson)
	return nil
}

func responseBodyShouldContainAsIts(arg1, arg2 string) error {
	return godog.ErrPending
}
