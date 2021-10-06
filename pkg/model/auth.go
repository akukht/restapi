package model

import (
	"errors"
	"restapi/pkg/database"
	hash "restapi/pkg/hashing"
	cjwt "restapi/pkg/jwt"

	"github.com/rs/zerolog/log"
)

//AuthUserDB User authorization
func AuthUserDB(login, password string) string {

	//Connect to database
	db, err := database.ConnectDB()
	var token string
	if err != nil {
		log.Fatal().Err(err).Msg("Connection to DB")
	}

	// Get User ID and Password by login
	row := db.QueryRow("SELECT id, password FROM users WHERE login = $1", login)
	var pass string
	var id int

	err = row.Scan(&id, &pass)
	if err != nil {
		return ""
	}

	// Compare user passwords
	if hash.CheckPasswordHash(password, pass) {
		// Generate token for user
		token, _ = cjwt.GenerateJWT()
		// Write token for user in database
		UpdateUserToken(id, token)
	}

	return token
}

//UserLogout Logout user
func UserLogout(token string) (bool, error) {
	//Connect to database
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal().Err(err).Msg("Connection to DB")
	}

	// Get user id by token
	row := db.QueryRow("SELECT id FROM users WHERE token = $1", token)
	var id int

	err = row.Scan(&id)

	if err != nil {
		return false, errors.New("user did not found")
	}
	//Remove user token
	UpdateUserToken(id, "")

	return true, nil
}
