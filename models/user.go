package models

import "time"

// TODO: consider doing ip search to get estimate location instead
type Location struct {
	Longitude float64 `json:"longitude"`
	Latitude  float64 `json:"latitude"`
}

type User struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	FirstName string    `json:"first_name" form:"first"`
	LastName  string    `json:"last_name" form:"last"`
	Email     string    `json:"email" form:"email"`
	// Location  Location  `json:"location"`
}
