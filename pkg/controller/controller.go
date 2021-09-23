package controller

import (
	"encoding/json"
	"math/rand"
	"net/http"
	. "restapi/pkg/jwt"
	"restapi/pkg/model"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

var mySigningKey = []byte("secret")

const (
	EventMessage404      = "Event not found"
	EventMessage401      = "You need to authorize"
	EventDeleteMessag200 = "Event was deleted"
	AuthMessage          = "You need to authorize"
	NeenAuth401          = "You need to authorize"
	OldToken             = "Your token has expired"
	LoggedOut            = "You are logged out"
)

func BetweenFilterEvents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	urlParams := r.URL.Query()
	event := model.BetweenFilterEvents(urlParams)
	json.NewEncoder(w).Encode(event)
}

func WeekFilterEvents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	urlParams := r.URL.Query()
	event := model.WeekFilterEvents(urlParams)
	json.NewEncoder(w).Encode(event)
}

func FilterEvents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	urlParams := r.URL.Query()
	event := model.FilterEvents(urlParams)
	json.NewEncoder(w).Encode(event)
}

func DeleteEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)

	if _, ok := model.EventsData[id["id"]]; ok {
		delete(model.EventsData, id["id"])
		json.NewEncoder(w).Encode(model.ResponseCode{StatusCode: 200, Message: EventDeleteMessag200})
	} else {
		json.NewEncoder(w).Encode(model.ResponseCode{StatusCode: 404, Message: EventMessage404})
	}

}

func UpdateEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)
	var event model.Events
	_ = json.NewDecoder(r.Body).Decode(&event)

	if _, ok := model.EventsData[id["id"]]; ok {
		model.EventsData[id["id"]] = event
	} else {
		err := model.ResponseCode{StatusCode: 404, Message: EventMessage404}
		json.NewEncoder(w).Encode(err)
	}
	json.NewEncoder(w).Encode(model.EventsData)
}

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var createEvent model.Events
	_ = json.NewDecoder(r.Body).Decode(&createEvent)

	model.EventsData[strconv.Itoa(rand.Intn(1000000))] = createEvent
	json.NewEncoder(w).Encode(model.EventsData)
}

func GetEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)

	token, err := GetGWTToken(r.Header["Token"][0])
	if err != nil {
		log.Warn().Err(err).Msg("Get events, GetGWTToken action in GetEvent() function")
	}
	userID, err := model.GetUserIDbyToken(token.Raw)
	if err != nil {
		log.Warn().Err(err).Msg("Get events, GetUserIDbyToken action in GetEvent() function")
	}

	if _, ok := model.EventsData[id["id"]]; ok {
		userEvent := model.GetUserEvent(userID, id["id"], model.EventsData)
		json.NewEncoder(w).Encode(model.TimeZoneConverter(userID, model.UsersList[userID].TimeZone, userEvent))
	} else {
		json.NewEncoder(w).Encode(model.ResponseCode{StatusCode: 404, Message: EventMessage404})
	}
}

func GetEvents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Header["Token"] != nil {
		token, err := GetGWTToken(r.Header["Token"][0])

		if err != nil {
			log.Warn().Err(err).Msg("Get events request, GetGWTToken action in GetEvents() function")
		}

		userID, err := model.GetUserIDbyToken(token.Raw)

		if err != nil {
			log.Warn().Err(err).Msg("Get events request, GetUserIDbyToken action in GetEvents() function")
		}

		userEvent, err := model.GetUserEvents(userID, model.EventsData)

		if err != nil {
			log.Warn().Err(err).Msg("Get events request, GetUserEvents action in GetEvents() function")
		}

		if len(userEvent) != 0 {
			w.WriteHeader(200)
			json.NewEncoder(w).Encode(model.TimeZoneConverter(userID, model.UsersList[userID].TimeZone, userEvent))
		} else {
			w.WriteHeader(404)
			json.NewEncoder(w).Encode(model.ResponseCode{StatusCode: 404, Message: EventMessage404})
		}
	} else {
		w.WriteHeader(401)
		json.NewEncoder(w).Encode(model.ResponseCode{StatusCode: 404, Message: EventMessage401})
	}
}

func IsAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Header["Token"] != nil {
			token, err := GetGWTToken(r.Header["Token"][0])
			if err != nil {
				w.WriteHeader(401)
				json.NewEncoder(w).Encode(model.ResponseCode{StatusCode: 401, Message: OldToken})
			}
			if token.Valid {
				endpoint(w, r)
			}
		} else {
			w.WriteHeader(401)
			json.NewEncoder(w).Encode(model.ResponseCode{StatusCode: 401, Message: NeenAuth401})
		}
	})
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	token, err := GetGWTToken(r.Header["Token"][0])

	if err != nil {
		log.Warn().Err(err).Msg("Update user request, GetGWTToken action in UpdateUser() function")
	}

	id, err := model.GetUserIDbyToken(token.Raw)

	if err != nil {
		log.Warn().Err(err).Msg("Update user request, GetUserIDbyToken action in UpdateUser() function")
	}
	var user model.Users
	_ = json.NewDecoder(r.Body).Decode(&user)

	model.UsersList[id] = model.Users{
		Login:    model.UsersList[id].Login,
		Password: model.UsersList[id].Password,
		TimeZone: user.TimeZone,
		Token:    model.UsersList[id].Token,
	}

	json.NewEncoder(w).Encode(model.UsersList[id])
}

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user model.Users
	_ = json.NewDecoder(r.Body).Decode(&user)

	token, err := model.Authorization(user.Login, user.Password)

	if err != nil {
		log.Error().Err(err).Msg("After authorization in login() function")
		json.NewEncoder(w).Encode(model.ResponseCode{StatusCode: 401, Message: err.Error()})
		return
	}

	json.NewEncoder(w).Encode(model.ResponseCode{StatusCode: 200, Message: token})
}

func Logout(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("authorized", "false")

	token, err := GetGWTToken(r.Header["Token"][0])

	if err != nil {
		log.Warn().Err(err).Msg("Loguot action after GetGWTToken() function")
	}

	userID, err := model.GetUserIDbyToken(token.Raw)
	if err != nil {
		log.Warn().Err(err).Msg("Loguot action after GetUserIDbyToken() function")
	}

	model.UsersList[userID] = model.Users{
		Login:    model.UsersList[userID].Login,
		Password: model.UsersList[userID].Password,
		TimeZone: model.UsersList[userID].TimeZone,
	}

	response := model.ResponseCode{StatusCode: 200, Message: LoggedOut}
	json.NewEncoder(w).Encode(response)
}
