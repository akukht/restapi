package jwt

import (
	"errors"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var mySigningKey = []byte("secret")

//GenerateJWT generate JWT Token
func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["exp"] = time.Now().Add(time.Minute * 60 * 24).Unix()

	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

//GetGWTToken get JWT Token
func GetGWTToken(tokenString string) (*jwt.Token, error) {

	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("token was expired")
		}
		return mySigningKey, nil

	})

}
