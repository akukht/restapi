package model

import (
	"restapi/pkg/database"

	"github.com/rs/zerolog/log"
)

//CreateNewEvent create new event
func CreateNewEvent(event Events, token string) (bool, error) {
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal().Err(err).Msg("Connection to DB")
	}
	// Get user id by token
	row := db.QueryRow("SELECT id FROM users WHERE token = $1", token)
	var userID int
	err = row.Scan(&userID)

	if err != nil {
		return false, err
	}

	// Insert new event
	lastInsertID := 0
	insertEvent := `
	INSERT INTO events (name, description, timezone)
	VALUES ($1, $2, $3) RETURNING id`
	err = db.QueryRow(insertEvent, event.Name, event.Desc, event.TimeZone).Scan(&lastInsertID)
	if err != nil {
		return false, err
	}

	//Attach event to user
	attachEvent := `
	INSERT INTO users_events (user_id, event_id)
	VALUES ($1, $2)`

	_, err = db.Exec(attachEvent, userID, lastInsertID)

	if err != nil {
		return false, err
	}

	// Insert date to event
	insertEventDate := `
	INSERT INTO date_lists (event_id, year, month, day, hour, minutes)
	VALUES ($1, $2, $3, $4, $5, $6)`
	_, err = db.Exec(insertEventDate, lastInsertID, event.Time.Year, event.Time.Month, event.Time.Day, event.Time.Hour, event.Time.Minutes)
	if err != nil {
		return false, err
	}

	return true, nil
}
