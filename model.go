package main

type Employee struct {
	ID          string `json:"id"`
	Designation string `json:"designation"`
	Department  string `json:"department"`
	Role        string `json:"role"`
	FlagActive  bool   `json:"flagactive"`
}

type Spot struct {
	// TODO: implement
}

// TODO: remove code below
// Book Struct
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Author Struct
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}
