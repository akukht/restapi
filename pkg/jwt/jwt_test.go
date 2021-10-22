package jwt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateJWT(t *testing.T) {

	t.Run("First test", func(t *testing.T) {
		got, err := GenerateJWT()
		want := 129
		assert.Nil(t, err)
		assert.Equal(t, want, len(got), "JWT Token mast have 129 symbols")
	})

	t.Run("Second test", func(t *testing.T) {
		got, generateErr := GenerateJWT()
		TokenValidate, validateErr := GetGWTToken(got)
		want := 129
		assert.Nil(t, generateErr)
		assert.Nil(t, validateErr)
		assert.Equal(t, want, len(TokenValidate.Raw), "JWT Token mast have 129 symbols")
	})

}

func TestGetGWTToken(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got, generateErr := GenerateJWT()
		TokenValidate, validateErr := GetGWTToken(got)
		want := 129
		assert.Nil(t, generateErr)
		assert.Nil(t, validateErr)
		assert.Equal(t, want, len(TokenValidate.Raw), "JWT Token mast have 129 symbols")
	})
}
