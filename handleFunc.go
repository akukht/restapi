package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"strings"

	"github.com/rs/zerolog/log"

	. "restapi/pkg/jwt"

	"github.com/gorilla/mux"
)

var mySigningKey = []byte("secret")

func betweenFilterEvents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	urlParams := r.URL.Query()
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
	json.NewEncoder(w).Encode(event)
}

func weekFilterEvents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	urlParams := r.URL.Query()
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
	json.NewEncoder(w).Encode(event)
}

func filterEvents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	urlParams := r.URL.Query()
	var event = map[string]Events{}
	for id, val := range EventsData {
		if val.Time.Day == strings.Join(urlParams["day"], "") && val.Time.Month == strings.Join(urlParams["month"], "") && val.Time.Year == strings.Join(urlParams["year"], "") {
			event[id] = EventsData[id]
		}
	}

	json.NewEncoder(w).Encode(event)
}

func DeleteEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)

	if _, ok := EventsData[id["id"]]; ok {
		delete(EventsData, id["id"])
		json.NewEncoder(w).Encode(ResponseCode{StatusCode: 200, Message: EventDeleteMessag200})
	} else {
		json.NewEncoder(w).Encode(ResponseCode{StatusCode: 404, Message: EventMessage404})
	}

}

func updateEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)
	var event Events
	_ = json.NewDecoder(r.Body).Decode(&event)

	if _, ok := EventsData[id["id"]]; ok {
		EventsData[id["id"]] = event
	} else {
		err := ResponseCode{StatusCode: 404, Message: EventMessage404}
		json.NewEncoder(w).Encode(err)
	}
	json.NewEncoder(w).Encode(EventsData)
}

func createEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var createEvent Events
	_ = json.NewDecoder(r.Body).Decode(&createEvent)

	EventsData[strconv.Itoa(rand.Intn(1000000))] = createEvent
	json.NewEncoder(w).Encode(EventsData)
}

func GetEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)

	token, err := GetGWTToken(r.Header["Token"][0])
	if err != nil {
		//fmt.Fprintf(w, err.Error())
		log.Warn().Err(err).Msg("Get events, GetGWTToken action in GetEvent() function")
	}
	userID, err := GetUserIDbyToken(token.Raw)
	if err != nil {
		log.Warn().Err(err).Msg("Get events, GetUserIDbyToken action in GetEvent() function")
	}

	if _, ok := EventsData[id["id"]]; ok {
		userEvent := GetUserEvent(userID, id["id"], EventsData)
		json.NewEncoder(w).Encode(TimeZoneConverter(userID, UsersList[userID].TimeZone, userEvent))
	} else {
		json.NewEncoder(w).Encode(ResponseCode{StatusCode: 404, Message: EventMessage404})
	}
}

func GetEvents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Header["Token"] != nil {
		token, err := GetGWTToken(r.Header["Token"][0])

		if err != nil {
			//log.Fatal(w, err.Error())
			log.Warn().Err(err).Msg("Get events request, GetGWTToken action in GetEvents() function")
		}

		userID, err := GetUserIDbyToken(token.Raw)

		if err != nil {
			//log.Fatal(w, err.Error())
			log.Warn().Err(err).Msg("Get events request, GetUserIDbyToken action in GetEvents() function")
		}

		userEvent, err := GetUserEvents(userID, EventsData)

		if err != nil {
			log.Warn().Err(err).Msg("Get events request, GetUserEvents action in GetEvents() function")
		}

		if len(userEvent) != 0 {
			w.WriteHeader(200)
			json.NewEncoder(w).Encode(TimeZoneConverter(userID, UsersList[userID].TimeZone, userEvent))
		} else {
			w.WriteHeader(404)
			json.NewEncoder(w).Encode(ResponseCode{StatusCode: 404, Message: EventMessage404})
		}
	} else {
		w.WriteHeader(401)
		json.NewEncoder(w).Encode(ResponseCode{StatusCode: 404, Message: EventMessage401})
	}
}

func isAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Header["Token"] != nil {
			token, err := GetGWTToken(r.Header["Token"][0])

			if err != nil {
				w.WriteHeader(401)
				json.NewEncoder(w).Encode(ResponseCode{StatusCode: 401, Message: OldToken})
			}
			if token.Valid {
				endpoint(w, r)
			}
		} else {
			w.WriteHeader(401)
			json.NewEncoder(w).Encode(ResponseCode{StatusCode: 401, Message: NeenAuth401})
		}
	})
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	token, err := GetGWTToken(r.Header["Token"][0])

	if err != nil {
		log.Warn().Err(err).Msg("Update user request, GetGWTToken action in UpdateUser() function")
	}

	id, err := GetUserIDbyToken(token.Raw)

	if err != nil {
		log.Warn().Err(err).Msg("Update user request, GetUserIDbyToken action in UpdateUser() function")
	}
	var user Users
	_ = json.NewDecoder(r.Body).Decode(&user)

	UsersList[id] = Users{
		Login:    UsersList[id].Login,
		Password: UsersList[id].Password,
		TimeZone: user.TimeZone,
		Token:    UsersList[id].Token,
	}

	json.NewEncoder(w).Encode(UsersList[id])
}
