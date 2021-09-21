package main

import "time"

type Employee struct {
	ID          string `json:"id"`
	Designation string `json:"designation"`
	Department  string `json:"department"`
	Role        string `json:"role"`
	FlagActive  bool   `json:"flag_active"`
}

type Spot struct {
	ID            string    `json:"id"`
	BookedDate    time.Time `json:"booked_date"`
	BookingStatus string    `json:"booking_status"`
	BookedOn      time.Time `json:"booked_on"`
	BookedBy      string    `json:"booked_by"`
}

type ResponseError struct {
	Error string `json:"error"`
}
