package enter

import (
	"github.com/labstack/echo"
	"Bulldog0.1/src/api/handler"
)

func SetRout(g *echo.Group)  {
	g.POST("/signin", handler.SignIn)
	g.POST("/signup", handler.SignUp)

}
