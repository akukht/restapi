package model

import (
	"errors"

	//	. "restapi/pkg/jwt"
	"strconv"
	"time"
)

//Authorization return token after successed authorization
func Authorization(login, password string) (string, error) {

	var token string

	if login == "" || password == "" {
		return "", errors.New("empty login or password")
	}

	token = AuthUserDB(login, password)

	if token == "" {
		return "", errors.New("wrong login or password")
	}
	return token, nil
}

//GetUserIDbyToken get user id by token (in struct)
func GetUserIDbyToken(token string) (string, error) {

	var userID string

	for i, v := range UsersList {
		if v.Token == token {
			userID = i
			return userID, nil
		}
	}
	return userID, errors.New("user id not found")
}

//GetUserEvents get user event in struct
func GetUserEvents(userID int, events map[string]Events) (map[string]Events, error) {

	var userEvent = map[string]Events{}
	if userID != 0 {
		for i, v := range events {
			if v.UserID == userID {
				userEvent[i] = events[i]
			}
		}
		return userEvent, nil
	}
	return userEvent, errors.New("empty user id")
}

//GetUserEvent get user event in struct
func GetUserEvent(userID int, eventID string, events map[string]Events) map[string]Events {

	var userEvent = map[string]Events{}

	for i, v := range events {
		if v.UserID == userID && i == eventID {
			userEvent[i] = events[i]
			break
		}
	}
	return userEvent
}

//TimeZoneConverter converter time zones
func TimeZoneConverter(userID int, userTimezone string, events map[string]Events) map[string]Events {
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
			ID:   v.ID,
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
