package controller

import (
	"net/http"
	"net/http/httptest"
	. "restapi/pkg/jwt"
	"restapi/pkg/model"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestGetEvents(t *testing.T) {

	GetUserEvents := []struct {
		Name       string
		StatusCode int
		UserID     string
	}{
		{
			Name:       "Get events from user with zero events",
			StatusCode: 404,
			UserID:     "2",
		},
		{
			Name:       "Get user events by valid data",
			StatusCode: 200,
			UserID:     "1",
		},
	}
	for _, tt := range GetUserEvents {
		t.Run(tt.Name, func(t *testing.T) {

			request, _ := http.NewRequest(http.MethodGet, "/api/events", nil)
			token, _ := GenerateJWT()
			model.UsersList[tt.UserID] = model.Users{
				Login:    model.UsersList[tt.UserID].Login,
				Password: model.UsersList[tt.UserID].Password,
				Token:    token,
				TimeZone: model.UsersList[tt.UserID].TimeZone,
			}
			request.Header.Add("Token", token)

			response := httptest.NewRecorder()
			GetEvents(response, request)

			assert.Equal(t, tt.StatusCode, response.Result().StatusCode)
		})
	}

	t.Run("Get events without token", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/api/events", nil)
		response := httptest.NewRecorder()
		GetEvents(response, request)
		assert.Equal(t, 401, response.Result().StatusCode)
	})

}

func TestGetEvent(t *testing.T) {
	GetUserEvent := []struct {
		Name       string
		StatusCode int
		UserID     string
		EventID    string
	}{
		{
			Name:       "Get events from user with zero events",
			StatusCode: 404,
			UserID:     "2",
			EventID:    "100",
		},
		{
			Name:       "Get user events by valid data",
			StatusCode: 200,
			UserID:     "1",
			EventID:    "1",
		},
	}

	for _, tt := range GetUserEvent {
		t.Run(tt.Name, func(t *testing.T) {
			request, _ := http.NewRequest(http.MethodDelete, "/api/events/"+tt.EventID, nil)

			token, _ := GenerateJWT()
			request.Header.Add("Token", token)
			response := httptest.NewRecorder()
			vars := map[string]string{
				"id": tt.EventID,
			}
			request = mux.SetURLVars(request, vars)
			GetEvent(response, request)

			assert.Equal(t, tt.StatusCode, response.Result().StatusCode)
		})
	}
}

func TestDeleteEvent(t *testing.T) {
	RemoveEvents := []struct {
		Name       string
		StatusCode int
		EventID    string
	}{
		{
			Name:       "Event not found msg",
			StatusCode: 404,
			EventID:    "100",
		},
		{
			Name:       "Event was deleted msg",
			StatusCode: 200,
			EventID:    "1",
		},
	}

	for _, tt := range RemoveEvents {
		t.Run(tt.Name, func(t *testing.T) {
			request, _ := http.NewRequest(http.MethodDelete, "/api/events/"+tt.EventID, nil)

			token, _ := GenerateJWT()
			request.Header.Add("Token", token)
			response := httptest.NewRecorder()
			vars := map[string]string{
				"id": tt.EventID,
			}
			request = mux.SetURLVars(request, vars)
			DeleteEvent(response, request)
			assert.Equal(t, tt.StatusCode, response.Result().StatusCode)
		})
	}
}
