package controller

import (
	"encoding/json"
	"net/http"
	cjwt "restapi/pkg/jwt"
	"restapi/pkg/model"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
)

var validate *validator.Validate

//Messages for responses
const (
	EventMessage404      = "Event not found"
	EventMessage401      = "You need to authorize"
	EventDeleteMessag200 = "Event was deleted"
	AuthMessage          = "You need to authorize"
	NeenAuth401          = "You need to authorize"
	OldToken             = "Your token has expired"
	LoggedOut            = "You are logged out"
)

//BetweenFilterEvents filte for star and enddate
func BetweenFilterEvents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	urlParams := r.URL.Query()

	//Get token from Headers
	token, err := cjwt.GetGWTToken(r.Header["Token"][0])
	if err != nil {
		log.Warn().Err(err).Msg("Get events, GetGWTToken action in GetEvent() function")
	}

	event, err := model.BetweenFilterEvents(token.Raw, urlParams)

	if err != nil {
		log.Warn().Err(err).Msg("Get events by filter")
	}

	err = json.NewEncoder(w).Encode(event)

	if err != nil {
		log.Warn().Err(err).Msg("Return BetweenFilterEvents JSON")
	}
}

//FilterEvents simple filter
func FilterEvents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//Get token from Headers
	token, err := cjwt.GetGWTToken(r.Header["Token"][0])
	if err != nil {
		log.Warn().Err(err).Msg("Get events, GetGWTToken action in GetEvent() function")
	}

	urlParams := r.URL.Query()
	event, err := model.BasicFilter(token.Raw, urlParams)

	if err != nil {
		log.Warn().Err(err).Msg("Get event by filter")
		return
	}
	if len(event) != 0 {
		w.WriteHeader(200)
		err = json.NewEncoder(w).Encode(event)
		log.Warn().Err(err).Msg("Return JSON response (FilterEvents)")
	} else {
		w.WriteHeader(404)
		err = json.NewEncoder(w).Encode(model.ResponseCode{StatusCode: 404, Message: EventMessage404})
		log.Warn().Err(err).Msg("Return JSON response (FilterEvents)")
	}
}

//DeleteEvent delete event
func DeleteEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)

	//Get token from Headers
	token, err := cjwt.GetGWTToken(r.Header["Token"][0])
	if err != nil {
		log.Warn().Err(err).Msg("Get events, GetGWTToken action in GetEvent() function")
	}

	eventID, err := strconv.Atoi(id["id"])
	if err != nil {
		w.WriteHeader(400)
		log.Warn().Err(err).Msg("Get event, convert parameter to integer in GetEvent() function")
		err = json.NewEncoder(w).Encode(model.ResponseCode{StatusCode: 400, Message: "The passed parameter is not a number"})
		if err != nil {
			log.Warn().Err(err).Msg("Return JSON response (DeleteEvent) 'The passed parameter is not a number' ")
		}
	}

	err = model.DeleteEvent(token.Raw, eventID)

	if err != nil {
		log.Warn().Err(err).Msg("Delete event")
		w.WriteHeader(404)
		err = json.NewEncoder(w).Encode(model.ResponseCode{StatusCode: 404, Message: EventMessage404})
		if err != nil {
			log.Warn().Err(err).Msg("Return JSON response (DeleteEvent) 'Event not found'")
		}
		return
	}

	w.WriteHeader(200)
	err = json.NewEncoder(w).Encode(model.ResponseCode{StatusCode: 200, Message: EventDeleteMessag200})
	if err != nil {
		log.Warn().Err(err).Msg("Return JSON response (DeleteEvent) 'Event was deleted' ")
	}
}

//UpdateEvent for update event
func UpdateEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)

	validate = validator.New()
	var event model.Events
	_ = json.NewDecoder(r.Body).Decode(&event)

	//Get token from Headers
	token, err := cjwt.GetGWTToken(r.Header["Token"][0])
	if err != nil {
		log.Warn().Err(err).Msg("Get events, GetGWTToken action in GetEvent() function")
	}

	err = validate.Struct(event)
	if err != nil {
		log.Warn().Err(err).Msg("Update event, validate failed")
		w.WriteHeader(400)
		err = json.NewEncoder(w).Encode(model.ResponseCode{StatusCode: 400, Message: "Invalid data submitted"})
		if err != nil {
			log.Warn().Err(err).Msg("Update event, validate failed")
		}
		return
	}

	eventID, err := strconv.Atoi(id["id"])
	if err != nil {
		w.WriteHeader(400)
		log.Warn().Err(err).Msg("Get event, convert parameter to integer in GetEvent() function")
		err = json.NewEncoder(w).Encode(model.ResponseCode{StatusCode: 400, Message: "The passed parameter is not a number"})
		if err != nil {
			log.Warn().Err(err).Msg("The passed parameter is not a number UpdateEvent() function")
		}
	}

	_, err = model.UpdateEvent(token.Raw, event, eventID)

	if err != nil {
		w.WriteHeader(400)
		log.Warn().Err(err).Msg("Update event")
		err = json.NewEncoder(w).Encode(model.ResponseCode{StatusCode: 400, Message: "The passed parameter is not a number"})
		if err != nil {
			log.Warn().Err(err).Msg("The passed parameter is not a number UpdateEvent() function")
		}
	}

	w.WriteHeader(200)
	err = json.NewEncoder(w).Encode(model.ResponseCode{StatusCode: 200, Message: "Event was updated"})
	if err != nil {
		log.Warn().Err(err).Msg("Event was updated")
	}

}

//CreateEvent create event
func CreateEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	validate = validator.New()
	//Get token from Headers
	token, err := cjwt.GetGWTToken(r.Header["Token"][0])
	if err != nil {
		log.Warn().Err(err).Msg("Get events, GetGWTToken action in GetEvent() function")
	}

	var createEvent model.Events
	_ = json.NewDecoder(r.Body).Decode(&createEvent)

	err = validate.Struct(createEvent)

	if err != nil {
		log.Warn().Err(err).Msg("Create event, validate failed")
		w.WriteHeader(400)
		_ = json.NewEncoder(w).Encode(model.ResponseCode{StatusCode: 400, Message: "Invalid data submitted"})
		return
	}

	_, err = model.CreateNewEvent(createEvent, token.Raw)
	if err != nil {
		log.Warn().Err(err).Msg("Failed to create new event")
	}

	_ = json.NewEncoder(w).Encode(model.ResponseCode{StatusCode: 200, Message: "Event was created"})

}

//GetEvent get user event
func GetEvent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)

	eventID, err := strconv.Atoi(id["id"])
	if err != nil {
		log.Warn().Err(err).Msg("Get event, convert parameter to integer in GetEvent() function")
		_ = json.NewEncoder(w).Encode(model.ResponseCode{StatusCode: 404, Message: "The passed parameter is not a number"})
	}

	// Get token from Headers
	token, err := cjwt.GetGWTToken(r.Header["Token"][0])
	if err != nil {
		log.Warn().Err(err).Msg("Get events, GetGWTToken action in GetEvent() function")
	}

	// Get user event from database
	userEvent, err := model.GetUserEventDB(token.Raw, eventID)
	if err != nil {
		log.Warn().Err(err).Msg("Get events request, GetUserEvents action in GetEvents() function")
	}

	// Return JSON answer
	if len(userEvent) != 0 {
		w.WriteHeader(200)
		_ = json.NewEncoder(w).Encode(userEvent)
	} else {
		w.WriteHeader(404)
		_ = json.NewEncoder(w).Encode(model.ResponseCode{StatusCode: 404, Message: EventMessage404})
	}
}

//GetEvents get user events
func GetEvents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get token from Headers
	token, err := cjwt.GetGWTToken(r.Header["Token"][0])
	if err != nil {
		log.Warn().Err(err).Msg("Get events request, GetGWTToken action in GetEvents() function")
	}

	// Get user events from database
	userEvent, err := model.GetUserEventsDB(token.Raw)
	if err != nil {
		log.Warn().Err(err).Msg("Get events request, GetUserEvents action in GetEvents() function")
	}

	// Return JSON answer
	if len(userEvent) != 0 {
		w.WriteHeader(200)
		_ = json.NewEncoder(w).Encode(userEvent)
	} else {
		w.WriteHeader(404)
		_ = json.NewEncoder(w).Encode(model.ResponseCode{StatusCode: 404, Message: EventMessage404})
	}
}

//IsAuthorized check if user authorized
func IsAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Header["Token"] != nil {
			token, err := cjwt.GetGWTToken(r.Header["Token"][0])
			if err != nil {
				w.WriteHeader(401)
				_ = json.NewEncoder(w).Encode(model.ResponseCode{StatusCode: 401, Message: OldToken})
			}
			if token.Valid {
				endpoint(w, r)
			}
		} else {
			w.WriteHeader(401)
			_ = json.NewEncoder(w).Encode(model.ResponseCode{StatusCode: 401, Message: NeenAuth401})
		}
	})
}

//UpdateUser update user
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Get token from Headers
	token, err := cjwt.GetGWTToken(r.Header["Token"][0])
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

	_ = json.NewEncoder(w).Encode(model.UsersList[id])
}

//Login for login user
func Login(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var user model.Users
	_ = json.NewDecoder(r.Body).Decode(&user)

	token, err := model.Authorization(user.Login, user.Password)

	if err != nil {
		log.Error().Err(err).Msg("After authorization in login() function")
		_ = json.NewEncoder(w).Encode(model.ResponseCode{StatusCode: 401, Message: err.Error()})
		return
	}

	_ = json.NewEncoder(w).Encode(model.ResponseCode{StatusCode: 200, Message: "Success", Token: token})

}

//Logout for logout user
func Logout(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("authorized", "false")

	// Get token from Headers
	token, err := cjwt.GetGWTToken(r.Header["Token"][0])
	if err != nil {
		log.Warn().Err(err).Msg("Loguot action after GetGWTToken() function")
	}

	res, err := model.UserLogout(token.Raw)
	if err != nil {
		log.Warn().Err(err).Msg("Loguot action after UserLogout() function")
	}
	//Return JSON answer
	if res {
		response := model.ResponseCode{StatusCode: 200, Message: LoggedOut}
		_ = json.NewEncoder(w).Encode(response)
	} else {
		response := model.ResponseCode{StatusCode: 401, Message: "Something went wrong"}
		_ = json.NewEncoder(w).Encode(response)
	}

}
