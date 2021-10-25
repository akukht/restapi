package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewEvent(t *testing.T) {
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2MzUwOTc3Njh9.4goZSs_CdgVvjiDcrRhVnoU_kU15el7yBlooTG4kMy0"
	createEvent := Events{
		Name: "Yahoo event",
		Time: Date{Day: "12", Month: "03", Year: "2021"},
		Desc: "Description Yahoo event event",
	}

	res, err := CreateNewEvent(createEvent, token)

	assert.Equal(t, false, res)
	assert.Nil(t, err)
}
