package router

import (
	"github.com/labstack/echo"
	"Bulldog0.1/src/api/middleware"
	"Bulldog0.1/src/api/routes/enter"
	project2 "Bulldog0.1/src/api/routes/project"
	"Bulldog0.1/src/model/reqmodel"
	"github.com/go-playground/validator"
	"Bulldog0.1/src/api/routes/uploader"
)

func New () *echo.Echo{
	e := echo.New()


	/*
	*Set Validators
	*/
	v := reqmodel.UserValidator{}
	v.Validator = validator.New()
	e.Validator = &v

	/*
	* Create Group
	* someGroup := e.Group("/route");
	*/

	enterance := e.Group("/enter")
	api := e.Group("/api")
	project := api.Group("/project" )

	/*
	* Set all middleware
	* middleware.setSomeMiddleware(echo)
	*/
	middleware.SetMidWare(api)
	middleware.SetMidWare(project)

	/*
	* Set all endpoints (Routes) tooo groups
	* api.someGroup(ourGroup)
	*/

	enter.SetRout(enterance)
	project2.SetProjectRoute(project)
	uploader.SetUploaderRoute(api)



	return e
}
