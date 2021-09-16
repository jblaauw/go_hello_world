package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"time"
)

var spotAvailable bool
var requestBody *bytes.Buffer

func aOfficeSpotIs(spotType, availability string) error {
	// TODO: change implementation if there is a real way to check if there is a spot available
	spotAvailable = availability == "available"
	return nil
}

func theRequestHeaderIsSetTo(key, value string) error {
	req.Header.Set(key, value)
	return nil
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

func responseBodyShouldContainAsIts(attrExpectedValue, attrName string) error {
	var resBooking Spot
	err := json.Unmarshal(rrBookings.Body.Bytes(), &resBooking)
	if err != nil {
		return fmt.Errorf("Error while deserializing response body: %s", err.Error())
	}

	if attrValue := resBooking.getFieldString(attrName); attrValue != attrExpectedValue {
		return fmt.Errorf("Expected the following value: %s. Got %s", attrExpectedValue, attrValue)
	}

	return nil
}

func (spot *Spot) getFieldString(field string) string {
	r := reflect.ValueOf(spot)
	f := reflect.Indirect(r).FieldByName(field)
	return string(f.String())
}
