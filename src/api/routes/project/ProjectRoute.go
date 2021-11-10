package project

import (
	"github.com/labstack/echo"
	"Bulldog0.1/src/api/handler"
)

func SetProjectRoute(g *echo.Group){
	g.POST("/start",handler.Start)

	g.PUT("/basicinfo",handler.BasicInfo)
	g.PUT("/story",handler.Story)
	g.PUT("/program",handler.Program)
	g.PUT("/confirm",handler.Confirm)
	g.PUT("/editproject",handler.EditPreProject)

	g.GET("/getpreprojectbyuser",handler.UserPreProject)
	g.GET("/getprojectbyuser",handler.UserProject)

	g.DELETE("/deletepreproject",handler.DeletePreProject)

}
