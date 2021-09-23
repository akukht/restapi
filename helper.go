package main

import (
	"errors"
	"strconv"
	"time"
)

func GetUserIDbyToken(token string) (string, error) {

	var userID string

	for i, v := range UsersList {
		if v.Token == token {
			userID = i
			return userID, nil
		}
	}
	return userID, errors.New("User id not found")
}

func GetUserEvents(userID string, events map[string]Events) (map[string]Events, error) {

	var userEvent = map[string]Events{}
	if userID != "" {
		for i, v := range events {
			if v.UserID == userID {
				userEvent[i] = events[i]
			}
		}
		return userEvent, nil
	}
	return userEvent, errors.New("Empty user id")
}

func GetUserEvent(userID string, eventID string, events map[string]Events) map[string]Events {

	var userEvent = map[string]Events{}

	for i, v := range events {
		if v.UserID == userID && i == eventID {
			userEvent[i] = events[i]
			break
		}
	}
	return userEvent
}

func TimeZoneConverter(userID string, userTimezone string, events map[string]Events) map[string]Events {
	var userEvent = map[string]Events{}

	for i, v := range events {
		year, _ := strconv.Atoi(v.Time.Year)
		month, _ := strconv.Atoi(v.Time.Month)
		day, _ := strconv.Atoi(v.Time.Day)
		hour, _ := strconv.Atoi(v.Time.Hour)
		minutes, _ := strconv.Atoi(v.Time.Minutes)

		eur, _ := time.LoadLocation(events[i].TimeZone)

		t := time.Date(year, time.Month(month), day, hour, minutes, 0, 0, eur)

		phx, _ := time.LoadLocation(userTimezone)
		converted := t.In(phx)

		userEvent[i] = Events{
			Name: v.Name,
			Time: Date{
				Year:    v.Time.Year,
				Month:   v.Time.Month,
				Day:     v.Time.Day,
				Hour:    v.Time.Hour,
				Minutes: v.Time.Minutes},
			Desc:      v.Desc,
			TimeZone:  v.TimeZone,
			UserID:    userID,
			Reminding: converted,
		}
	}
	return userEvent
}
