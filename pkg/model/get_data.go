package model

import (
	"context"
	"fmt"
	"net/url"
	"restapi/pkg/database"
	"strconv"
	"strings"

	"github.com/rs/zerolog/log"
)

const getUserID = `SELECT id FROM users WHERE token = $1`

//GetUserIDbyToken get user id by token
func (q *Queries) GetUserIDbyToken(ctx context.Context, token string) (int, error) {
	row := q.db.QueryRowContext(ctx, getUserID, token)
	var userID int
	err := row.Scan(&userID)
	if err != nil {
		return 0, err
	}
	return userID, nil
}

//BetweenFilterEvents get events after filtering
func BetweenFilterEvents(token string, urlParams url.Values) (map[string]Events, error) {

	var events = map[string]Events{}
	// Connect tu database
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal().Err(err).Msg("Connection to DB")
	}

	// Get user id by token
	row := db.QueryRow("SELECT id FROM users WHERE token = $1", token)
	var userID int
	err = row.Scan(&userID)
	if err != nil {
		log.Fatal().Err(err).Msg("Get user by token")
		return events, err
	}

	startNum := strings.Join(urlParams["start"], "")
	endNum := strings.Join(urlParams["end"], "")
	month := strings.Join(urlParams["month"], "")
	year := strings.Join(urlParams["year"], "")

	rows, err := db.Query(`  SELECT events.id, events.name, events.description, events.timezone, date_lists.year, date_lists.month, date_lists.day, date_lists.hour, date_lists.minutes 
	FROM events
	RIGHT JOIN date_lists
	ON events.id = date_lists.event_id
	WHERE events.id IN (SELECT event_id FROM date_lists WHERE year=$1 AND month=$2 AND day BETWEEN $3 AND $4)`, year, month, startNum, endNum)

	if err != nil {
		log.Fatal().Err(err).Msg("Get events by filter")
		return events, err
	}
	defer rows.Close()

	// Generate event struct by db data
	for rows.Next() {
		eventsData := Events{}
		err := rows.Scan(&eventsData.ID, &eventsData.Name, &eventsData.Desc, &eventsData.TimeZone, &eventsData.Time.Year, &eventsData.Time.Month, &eventsData.Time.Day, &eventsData.Time.Hour, &eventsData.Time.Minutes)
		if err != nil {
			log.Fatal().Err(err).Msg("Scan events by filter")
			continue
		}

		events[strconv.Itoa(eventsData.ID)] = Events{
			ID:   eventsData.ID,
			Name: eventsData.Name,
			Time: Date{
				Year:    eventsData.Time.Year,
				Month:   eventsData.Time.Month,
				Day:     eventsData.Time.Day,
				Hour:    eventsData.Time.Hour,
				Minutes: eventsData.Time.Minutes},
			Desc:     eventsData.Desc,
			TimeZone: eventsData.TimeZone,
			UserID:   userID,
		}
	}

	return events, nil
}

//BasicFilter simple filter with few parameters
func BasicFilter(token string, urlParams url.Values) (map[string]Events, error) {

	// Connect tu database
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal().Err(err).Msg("Connection to DB")
	}

	// Create empty event struct
	UserEvents := map[string]Events{}

	// Get user id by token
	row := db.QueryRow("SELECT id FROM users WHERE token = $1", token)
	var userID int
	err = row.Scan(&userID)
	if err != nil {
		log.Fatal().Err(err).Msg("Get user by token")
		return UserEvents, err
	}

	day := strings.Join(urlParams["day"], "")
	month := strings.Join(urlParams["month"], "")
	year := strings.Join(urlParams["year"], "")

	rows, err := db.Query(`SELECT events.id, events.name, events.description, events.timezone, date_lists.year, date_lists.month, date_lists.day, date_lists.hour, date_lists.minutes 
	FROM events
	RIGHT JOIN date_lists
	ON events.id = date_lists.event_id
	WHERE events.id IN (SELECT event_id FROM date_lists WHERE year=$1 AND month=$2 AND day=$3)`, year, month, day)

	if err != nil {
		log.Fatal().Err(err).Msg("Get events by filter")
		return UserEvents, err
	}
	defer rows.Close()
	// Generate event struct by db data
	for rows.Next() {
		eventsData := Events{}
		err := rows.Scan(&eventsData.ID, &eventsData.Name, &eventsData.Desc, &eventsData.TimeZone, &eventsData.Time.Year, &eventsData.Time.Month, &eventsData.Time.Day, &eventsData.Time.Hour, &eventsData.Time.Minutes)
		if err != nil {
			log.Fatal().Err(err).Msg("Scan events by filter")
			continue
		}

		UserEvents[strconv.Itoa(eventsData.ID)] = Events{
			ID:   eventsData.ID,
			Name: eventsData.Name,
			Time: Date{
				Year:    eventsData.Time.Year,
				Month:   eventsData.Time.Month,
				Day:     eventsData.Time.Day,
				Hour:    eventsData.Time.Hour,
				Minutes: eventsData.Time.Minutes},
			Desc:     eventsData.Desc,
			TimeZone: eventsData.TimeZone,
			UserID:   userID,
		}
	}

	return UserEvents, nil
}

//GetUserEventDB get user event
func GetUserEventDB(token string, paramID int) (map[string]Events, error) {

	// Create empty events struct
	UserEvents := map[string]Events{}

	// Connect tu database
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal().Err(err).Msg("Connection to DB")
	}

	// Get user id by token
	row := db.QueryRow("SELECT id FROM users WHERE token = $1", token)
	var id int
	err = row.Scan(&id)
	if err != nil {
		return UserEvents, err
	}

	// Check if event have in current user
	UserEvent := db.QueryRow("SELECT event_id FROM users_events WHERE event_id = $1 AND user_id = $2", paramID, id)
	var checkEvent int
	err = UserEvent.Scan(&checkEvent)
	if err != nil {
		return UserEvents, err
	}

	// Get user event by query param
	row = db.QueryRow(`SELECT events.id, events.name, events.description, events.timezone, date_lists.year, date_lists.month, date_lists.day, date_lists.hour, date_lists.minutes FROM events
		RIGHT JOIN date_lists
		ON events.id = date_lists.event_id
		WHERE events.id = $1
		ORDER BY events.id`, paramID)

	eventsData := Events{}
	err = row.Scan(&eventsData.ID, &eventsData.Name, &eventsData.Desc, &eventsData.TimeZone, &eventsData.Time.Year, &eventsData.Time.Month, &eventsData.Time.Day, &eventsData.Time.Hour, &eventsData.Time.Minutes)

	if err != nil {
		return UserEvents, err
	}

	// Generate event struct by db data
	UserEvents[strconv.Itoa(eventsData.ID)] = Events{
		ID:   eventsData.ID,
		Name: eventsData.Name,
		Time: Date{
			Year:    eventsData.Time.Year,
			Month:   eventsData.Time.Month,
			Day:     eventsData.Time.Day,
			Hour:    eventsData.Time.Hour,
			Minutes: eventsData.Time.Minutes},
		Desc:     eventsData.Desc,
		TimeZone: eventsData.TimeZone,
		UserID:   id,
	}

	// Return event with time zone converter
	return TimeZoneConverter(id, "America/Chicago", UserEvents), nil
}

//GetUserEventsDB get all user eevnt
func GetUserEventsDB(token string) (map[string]Events, error) {

	// Create empty event struct
	UserEvents := map[string]Events{}

	// Connect to database
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal().Err(err).Msg("Connection to DB")
	}

	// Get user id by token
	row := db.QueryRow("SELECT id FROM users WHERE token = $1", token)

	var id int

	// Get all current user events
	err = row.Scan(&id)
	if err != nil {
		return UserEvents, err
	}
	rows, err := db.Query(`SELECT events.id, events.name, events.description, events.timezone, date_lists.year, date_lists.month, date_lists.day, date_lists.hour, date_lists.minutes  FROM events
		RIGHT JOIN date_lists
		ON events.id = date_lists.event_id
		WHERE events.id IN (SELECT event_id FROM users_events WHERE user_id = $1)
		ORDER BY events.id`, id)

	if err != nil {
		return UserEvents, err
	}

	defer rows.Close()

	// Generate event struct by db data
	for rows.Next() {
		eventsData := Events{}
		err := rows.Scan(&eventsData.ID, &eventsData.Name, &eventsData.Desc, &eventsData.TimeZone, &eventsData.Time.Year, &eventsData.Time.Month, &eventsData.Time.Day, &eventsData.Time.Hour, &eventsData.Time.Minutes)
		if err != nil {
			fmt.Println(err)
			continue
		}
		UserEvents[strconv.Itoa(eventsData.ID)] = Events{
			ID:   eventsData.ID,
			Name: eventsData.Name,
			Time: Date{
				Year:    eventsData.Time.Year,
				Month:   eventsData.Time.Month,
				Day:     eventsData.Time.Day,
				Hour:    eventsData.Time.Hour,
				Minutes: eventsData.Time.Minutes},
			Desc:     eventsData.Desc,
			TimeZone: eventsData.TimeZone,
			UserID:   id,
		}
	}

	// Return event with time zone converter
	return TimeZoneConverter(id, "America/Chicago", UserEvents), nil
}
