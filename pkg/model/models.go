package model

import (
	"errors"
	"net/url"
	hash "restapi/pkg/hashing"
	. "restapi/pkg/jwt"
	"strconv"
	"strings"
	"time"
)

func BetweenFilterEvents(urlParams url.Values) map[string]Events {
	var day int

	startNum, _ := strconv.Atoi(strings.Join(urlParams["start"], ""))
	endNum, _ := strconv.Atoi(strings.Join(urlParams["end"], ""))

	var event = map[string]Events{}
	for id, val := range EventsData {
		day, _ = strconv.Atoi(val.Time.Day)

		if day >= startNum && day <= endNum && val.Time.Month == strings.Join(urlParams["month"], "") && val.Time.Year == strings.Join(urlParams["year"], "") {
			event[id] = EventsData[id]
		}
	}

	return event
}

func WeekFilterEvents(urlParams url.Values) map[string]Events {
	weekNum, _ := strconv.Atoi(strings.Join(urlParams["week"], ""))

	weekNum = weekNum * 7
	var day int
	var event = map[string]Events{}
	for id, val := range EventsData {
		day, _ = strconv.Atoi(val.Time.Day)

		if day < weekNum && val.Time.Month == strings.Join(urlParams["month"], "") && val.Time.Year == strings.Join(urlParams["year"], "") {
			event[id] = EventsData[id]
		}
	}

	return event
}

func FilterEvents(urlParams url.Values) map[string]Events {
	var event = map[string]Events{}
	for id, val := range EventsData {
		if val.Time.Day == strings.Join(urlParams["day"], "") && val.Time.Month == strings.Join(urlParams["month"], "") && val.Time.Year == strings.Join(urlParams["year"], "") {
			event[id] = EventsData[id]
		}
	}

	return event
}

func Authorization(login, password string) (string, error) {

	var token string

	if login == "" || password == "" {
		return "", errors.New("empty login or password")
	}

	for i, v := range UsersList {
		if v.Login == login {

			if hash.CheckPasswordHash(password, v.Password) {
				token, _ = GenerateJWT()

				UsersList[i] = Users{
					Login:    UsersList[i].Login,
					Password: UsersList[i].Password,
					Token:    token,
					TimeZone: UsersList[i].TimeZone,
				}

				break
			}
		}
	}

	if token != "" {
		return token, nil
	} else {
		return "", errors.New("wrong login or password")
	}

}

//------------helper

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
