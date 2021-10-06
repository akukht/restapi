package model

import (
	"restapi/pkg/database"

	"github.com/rs/zerolog/log"
)

//DeleteEvent delete event
func DeleteEvent(token string, id int) error {

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
		return err
	}

	// Check if event have in current user
	UserEvent := db.QueryRow("SELECT event_id FROM users_events WHERE event_id = $1 AND user_id = $2", id, userID)
	var checkEvent int
	err = UserEvent.Scan(&checkEvent)
	if err != nil {
		return err
	}

	sqlStatement := `
		DELETE FROM events
		WHERE id = $1;`
	_, err = db.Exec(sqlStatement, id)
	if err != nil {
		return err
	}

	return nil
}
