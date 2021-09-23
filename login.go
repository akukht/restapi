package main

import (
	"encoding/json"
	"errors"
	"net/http"
	hash "restapi/pkg/hashing"
	. "restapi/pkg/jwt"

	"github.com/rs/zerolog/log"
)

func login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user Users
	_ = json.NewDecoder(r.Body).Decode(&user)

	token, err := authorization(user.Login, user.Password)

	if err != nil {
		log.Error().Err(err).Msg("After authorization in login() function")
		json.NewEncoder(w).Encode(ResponseCode{StatusCode: 401, Message: err.Error()})
		return
	}

	json.NewEncoder(w).Encode(ResponseCode{StatusCode: 200, Message: token})
}

func logout(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("authorized", "false")

	token, err := GetGWTToken(r.Header["Token"][0])

	if err != nil {
		log.Warn().Err(err).Msg("Loguot action after GetGWTToken() function")
	}

	userID, err := GetUserIDbyToken(token.Raw)
	if err != nil {
		log.Warn().Err(err).Msg("Loguot action after GetUserIDbyToken() function")
	}

	UsersList[userID] = Users{
		Login:    UsersList[userID].Login,
		Password: UsersList[userID].Password,
		TimeZone: UsersList[userID].TimeZone,
	}

	response := ResponseCode{StatusCode: 200, Message: "You are logged out"}
	json.NewEncoder(w).Encode(response)
}

func authorization(login, password string) (string, error) {

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
