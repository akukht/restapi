package model

import (
	"time"
)

type ResponseCode struct {
	StatusCode int    `json:"code"`
	Message    string `json:"message"`
}

type Date struct {
	Year    string `json:"year"`
	Month   string `json:"month"`
	Day     string `json:"day"`
	Hour    string `json:"hour"`
	Minutes string `json:"minutes"`
}

type Events struct {
	Name      string    `json:"name"`
	Time      Date      `json:"time"`
	Desc      string    `json:"desc"`
	TimeZone  string    `json:"timezone"`
	UserID    string    `json:"userid"`
	Reminding time.Time `json:"reminding"`
}

var EventsData = map[string]Events{}

func init() {
	EventsData["1"] = Events{
		Name:     "Google Cloud",
		Time:     Date{Year: "2021", Month: "03", Day: "01", Hour: "10", Minutes: "30"},
		Desc:     "Description Google event",
		TimeZone: "America/Chicago",
		UserID:   "1",
	}

	EventsData["2"] = Events{
		Name:     "Amazon AWS",
		Time:     Date{Year: "2021", Month: "03", Day: "01", Hour: "10", Minutes: "30"},
		Desc:     "Description Amazon AWS event",
		TimeZone: "Europe/Rome",
		UserID:   "1",
	}

	EventsData["3"] = Events{
		Name:     "Microsoft Azure",
		Time:     Date{Year: "2021", Month: "03", Day: "01", Hour: "10", Minutes: "30"},
		Desc:     "Description Microsoft Azure event",
		TimeZone: "America/New_York",
		UserID:   "3",
	}

	EventsData["4"] = Events{
		Name:     "Yahoo event",
		Time:     Date{Year: "2021", Month: "03", Day: "01", Hour: "10", Minutes: "30"},
		Desc:     "Description Yahoo event event",
		TimeZone: "Europe/Rome",
		UserID:   "3",
	}

}
