package uploader

import (
	"github.com/labstack/echo"
	"Bulldog0.1/src/api/handler"
)

func SetUploaderRoute(g *echo.Group){
	g.POST("/upload",handler.Upload)
}
