package handlers

import (
	"github.com/labstack/echo"
	"database/sql"
	"github.com/dgrijalva/jwt-go"
)

func GetUserID(c echo.Context ) sql.NullInt64 {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*config.JWTSecretClaims)
	userId := sql.NullInt64{claims.ID, true}

	return userId
}