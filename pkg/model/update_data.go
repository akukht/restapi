package model

import (
	"database/sql"
	"restapi/pkg/database"

	"github.com/rs/zerolog/log"
)

//UpdateUserToken update user token in database
func UpdateUserToken(userID int, token string) sql.Result {

	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal().Err(err).Msg("Connection to DB")
	}

	row, err := db.Exec("UPDATE users SET token = $2 WHERE id = $1", userID, token)
	if err != nil {
		log.Fatal().Err(err).Msg("Update user token")
	}

	return row
}

//UpdateEvent update event in databaase
func UpdateEvent(token string, event Events, id int) (bool, error) {

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

	// Check if event have in current user
	UserEvent := db.QueryRow("SELECT event_id FROM users_events WHERE event_id = $1 AND user_id = $2", id, userID)
	var checkEvent int
	err = UserEvent.Scan(&checkEvent)
	if err != nil {
		log.Warn().Err(err).Msg("Check if event exist")
		return false, err
	}
	_, err = db.Exec("UPDATE events SET name = $2, description = $3, timezone = $4 WHERE id = $1", id, event.Name, event.Desc, event.TimeZone)
	if err != nil {
		return false, err
	}
	_, err = db.Exec("UPDATE date_lists SET year = $2, month = $3, day = $4, hour = $5, minutes = $6 WHERE event_id = $1", id, event.Time.Year, event.Time.Month, event.Time.Day, event.Time.Hour, event.Time.Minutes)
	if err != nil {
		return false, err
	}

	return true, nil
}
