package database

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

//ConnectDB connection to postgres database
func ConnectDB() (*sql.DB, error) {

	//Load .env file
	err := godotenv.Load()
	if err != nil {
		return nil, errors.New("error loading .env file")
	}

	//Connect to database and check errors
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	db, _ := sql.Open("postgres", psqlInfo)

	if err != nil {
		return nil, errors.New("error connect to postgres")
	}
	//defer db.Close()

	err = db.Ping()
	if err != nil {
		return nil, errors.New("error DB ping")
	}

	return db, nil
}
