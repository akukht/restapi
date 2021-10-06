package model

import (
	"time"
)

//ResponseCode struct for response codes
type ResponseCode struct {
	StatusCode int    `json:"code"`
	Message    string `json:"message"`
	Token      string `json:"token"`
}

//Date struct for event date
type Date struct {
	Year    string `json:"year" validate:"required"`
	Month   string `json:"month" validate:"required"`
	Day     string `json:"day" validate:"required"`
	Hour    string `json:"hour" validate:"required"`
	Minutes string `json:"minutes" validate:"required"`
}

//Events struct for event
type Events struct {
	ID        int       `json:"id"`
	Name      string    `json:"name" validate:"required"`
	Time      Date      `json:"time" validate:"required"`
	Desc      string    `json:"desc" validate:"required"`
	TimeZone  string    `json:"timezone" validate:"required"`
	UserID    int       `json:"userid"`
	Reminding time.Time `json:"reminding"`
}
