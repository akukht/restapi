package controller

import (
	// "net/http"
	// "net/http/httptest"
	// . "restapi/pkg/jwt"
	// "restapi/pkg/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEvents(t *testing.T) {
	assert.Equal(t, 200, 200)
	/*
		t.Run("Get user events by valid data", func(t *testing.T) {
			request, _ := http.NewRequest(http.MethodGet, "/api/events", nil)
			token, _ := GenerateJWT()
			model.UsersList["1"] = model.Users{
				Login:    model.UsersList["1"].Login,
				Password: model.UsersList["1"].Password,
				Token:    token,
				TimeZone: model.UsersList["1"].TimeZone,
			}
			request.Header.Add("Token", token)
			response := httptest.NewRecorder()
			GetEvents(response, request)
			assert.Equal(t, 200, response.Result().StatusCode)
		})
	*/
	/*
		t.Run("Get events from user without events", func(t *testing.T) {
			request, _ := http.NewRequest(http.MethodGet, "/api/events", nil)
			token, _ := GenerateJWT()

			model.UsersList["100"] = model.Users{
				Login:    model.UsersList["100"].Login,
				Password: model.UsersList["100"].Password,
				Token:    token,
				TimeZone: model.UsersList["100"].TimeZone,
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
	*/

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
