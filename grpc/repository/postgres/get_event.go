package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"restapi/pkg/model"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

func GetEventById(eventID int32) (map[string]model.Events, error) {
	err := godotenv.Load()
	UserEvents := map[string]model.Events{}

	if err != nil {
		log.Fatal().Err(err).Msg("error loading .env file")
	}

	//Connect to database and check errors
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatal().Err(err).Msg("Conn")
	}
	err = db.Ping()
	if err != nil {
		return UserEvents, errors.New("ping connection error")
	}

	// Get user event by query param
	row := db.QueryRow(`SELECT events.id, events.name, events.description, events.timezone, date_lists.year, date_lists.month, date_lists.day, date_lists.hour, date_lists.minutes FROM events
	RIGHT JOIN date_lists
	ON events.id = date_lists.event_id
	WHERE events.id = $1
	ORDER BY events.id`, eventID)

	eventsData := model.Events{}
	err = row.Scan(&eventsData.ID, &eventsData.Name, &eventsData.Desc, &eventsData.TimeZone, &eventsData.Time.Year, &eventsData.Time.Month, &eventsData.Time.Day, &eventsData.Time.Hour, &eventsData.Time.Minutes)

	if err != nil {
		return UserEvents, errors.New("row.Scan Error")
	}

	// Generate event struct by db data
	UserEvents[strconv.Itoa(eventsData.ID)] = model.Events{
		ID:   eventsData.ID,
		Name: eventsData.Name,
		Time: model.Date{
			Year:    eventsData.Time.Year,
			Month:   eventsData.Time.Month,
			Day:     eventsData.Time.Day,
			Hour:    eventsData.Time.Hour,
			Minutes: eventsData.Time.Minutes},
		Desc:     eventsData.Desc,
		TimeZone: eventsData.TimeZone,
	}

	return UserEvents, nil
}
