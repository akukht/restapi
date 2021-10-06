package model

//Users struct for user data
type Users struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	TimeZone string `json:"timezone"`
	Token    string `json:"token"`
}

//UsersList empty user struct
var UsersList = map[string]Users{}
