module restapi

go 1.16

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/rs/zerolog v1.25.0 // indirect
	github.com/stretchr/testify v1.7.0 // indirect
	golang.org/x/crypto v0.0.0-20210817164053-32db794688a5 // indirect
	restapi/pkg/hashing v0.0.0-00010101000000-000000000000 // indirect
	restapi/pkg/jwt v0.0.0-00010101000000-000000000000 // indirect
)

replace restapi/pkg/hashing => ./pkg/hashing

replace restapi/pkg/jwt => ./pkg/jwt
