package main

import (
	"net/http"
	"net/http/httptest"

	. "restapi/pkg/jwt"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
func TestDeleteEvent(t *testing.T) {
	tests := []struct {
		Name       string
		StatusCode int
		Message    string
	}{
		{
			Name:       "Event not found msg",
			StatusCode: 404,
			Message:    EventMessage404,
		},
		{
			Name:       "Event was deleted msg",
			StatusCode: 200,
			Message:    EventDeleteMessag200,
		},
	}

	for _, tt := range tests {

		t.Run(tt.Name, func(t *testing.T) {
			request, _ := http.NewRequest(http.MethodDelete, "/api/events/1", nil)

			response := httptest.NewRecorder()

			DeleteEvent(response, request)

			fmt.Println(response.Body)

			var respCode ResponseCode
			_ = json.NewDecoder(response.Body).Decode(&respCode)

			got := respCode

			if got.StatusCode != tt.StatusCode && got.Message != tt.Message {
				t.Errorf("got %q, want %q", got.Message, tt.Message)
			}
		})

	}

}
*/

func TestGetEvents(t *testing.T) {

	t.Run("Get user events by valid data", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/api/events", nil)
		token, _ := GenerateJWT()
		UsersList["1"] = Users{
			Login:    UsersList["1"].Login,
			Password: UsersList["1"].Password,
			Token:    token,
			TimeZone: UsersList["1"].TimeZone,
		}
		request.Header.Add("Token", token)
		response := httptest.NewRecorder()
		GetEvents(response, request)
		assert.Equal(t, 200, response.Result().StatusCode)
	})

	t.Run("Get events from user without events", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/api/events", nil)
		token, _ := GenerateJWT()

		UsersList["100"] = Users{
			Login:    UsersList["100"].Login,
			Password: UsersList["100"].Password,
			Token:    token,
			TimeZone: UsersList["100"].TimeZone,
		}

		request.Header.Add("Token", token)
		response := httptest.NewRecorder()
		GetEvents(response, request)

		assert.Equal(t, 404, response.Result().StatusCode)
	})

	t.Run("Get events without token", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/api/events", nil)
		response := httptest.NewRecorder()
		GetEvents(response, request)
		assert.Equal(t, 401, response.Result().StatusCode)
	})

	// t.Run("Get events with old token", func(t *testing.T) {
	// 	//Create request
	// 	request, _ := http.NewRequest(http.MethodGet, "/api/events", nil)

	// 	//Set expired token
	// 	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2MzIzMTU5NDh9.8T0HLWHU4l9Ham1P7WkcwLO9Gk5GP_lqG3SICNHewTc"

	// 	//Set token for user with ID-1
	// 	UsersList["1"] = Users{
	// 		Login:    UsersList["1"].Login,
	// 		Password: UsersList["1"].Password,
	// 		Token:    token,
	// 		TimeZone: UsersList["1"].TimeZone,
	// 	}

	// 	//Set token to headers
	// 	request.Header.Add("Token", token)

	// 	response := httptest.NewRecorder()
	// 	GetEvents(response, request)

	// 	assert.Equal(t, 401, response.Result().StatusCode)
	// })
}
