package models

import "time"

// TODO: consider doing ip search to get estimate location instead
type Location struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

type User struct {
	ID          string    `json:"id"`
	CountryCode string    `json:"country_code"`
	PhoneNumber string    `json:"phone_number"`
	CreatedAt   time.Time `json:"created_at"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Email       string    `json:"email"`
	Location    Location  `json:"location"`
	DOB         time.Time `json:"dob"`
}
