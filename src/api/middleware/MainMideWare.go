package middleware

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"Bulldog0.1/src/model/resmodel"
)

func SetMidWare(g *echo.Group)  {
	g.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		Claims:     &resmodel.UserResp{},
		SigningMethod:"HS512",
		SigningKey:[]byte("ArminRoozbehRezaBackendShireFrontChaiiGhahvehArianChaiiAvordBaCakeAbJooshNadashtim1234462194124628913746912846198246289713462389146298461238974621987461289746128946128746298146284564512874581273452378451238745123874521409650858585901786"),
	}))
}

