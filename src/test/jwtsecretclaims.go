package handlers


import "github.com/dgrijalva/jwt-go"

type JWTSecretClaims struct {
	ID int64 `json:"name"`
	jwt.StandardClaims
}
