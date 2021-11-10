package handler

import (
	"github.com/labstack/echo"
	"io/ioutil"
	"net/http"
	"encoding/json"
	"log"
	"Bulldog0.1/src/util"
	"Bulldog0.1/src/model/reqmodel"
	"Bulldog0.1/src/model/resmodel"
	"strings"
	"Bulldog0.1/src/model/db/userdb"
)


func SignIn(c echo.Context)error{
	userReq := reqmodel.UserReq{}
	body , err := ioutil.ReadAll(c.Request().Body)
	if err != nil{
		log.Printf("failed reading the body for add Story: %s",err)
		return c.JSON(http.StatusBadRequest,map[string]string{
			"status":"Error Reading Request Body",
		})
	}
	err = json.Unmarshal(body,&userReq)
	if err != nil{
		log.Printf("failed unmarshaling: %s",err)
		return c.JSON(http.StatusBadRequest,map[string]string{
			"status":"Error Unmarshaling",
		})
	}
	if err = c.Validate(userReq); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"status":"user pass wrong",
		})
	}
	pass := userReq.Password
	//email tooo lower case
	userReq.Email = strings.ToLower(userReq.Email)
	userDb := userdb.UserTo{}
	userDb.Email = userReq.Email
	userDb.Password = userReq.Password
	exist ,err := userDb.SelectByEmail()   /** "userTO selectbyemail" **/
	if err != nil {
		log.Print(err.Error())
		return c.JSON(http.StatusInternalServerError,map[string]string{
			"status":"Error Databasae",
		})
	}
	if !exist{
		return c.JSON(http.StatusBadRequest,map[string]string{
			"status":"User Not Exists",
		})
	}
	ok := util.ComparePassWithHash(pass, userDb.Password)
	if !ok {
		return c.JSON(http.StatusBadRequest,map[string]string{
			"status":"User Pass Wrong",
		})
	}
	userResp := &resmodel.UserResp{}
	userResp.Email = userDb.Email
	userResp.ID =  userDb.ID
	token, err := userResp.CreateJwtToken()
	if err != nil {
		log.Println("Error Creating JWT token ", err)
		return c.JSON(http.StatusInternalServerError,map[string]string{
			"status":"Error token",
		})
	}
	defer c.Request().Body.Close()
	return c.JSON(http.StatusOK, map[string]string{
		"token":token,
	})
}

func SignUp(c echo.Context) error {
	userReq := reqmodel.UserReq{}
	body , err := ioutil.ReadAll(c.Request().Body)
	if err != nil {
		log.Printf("failed reading the body for adding sign up: %s",err)
		return c.JSON(http.StatusBadRequest,map[string]string{
			"status":"Error Reading Request Body",
		})
	}
	err = json.Unmarshal(body,&userReq)
	if err != nil{
		log.Printf("failed unmarshaling: %s",err)
		return c.JSON(http.StatusBadRequest,map[string]string{
			"status":"Error Unmarshaling",
		})
	}
	if err = c.Validate(userReq); err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusBadRequest, map[string]string{
			"status":"user pass wrong",
		})
	}
	//email tooo lower case
	userReq.Email = strings.ToLower(userReq.Email)
	userDb := userdb.UserTo{}
	userDb.Email = userReq.Email
	userDb.Password = userReq.Password
	exist ,err := userDb.SelectByEmail()
	if exist {
		return c.JSON(http.StatusBadRequest,map[string]string{
			"status":"Email already exists!",
		})
	}
	if err != nil {
		log.Print(err.Error())
		return c.JSON(http.StatusInternalServerError,map[string]string{
			"status":"Error Databasae",
		})
	}
	userDb.Password , err = util.HashPassword(userDb.Password)
	if err != nil {
		log.Print(err.Error())
		return c.JSON(http.StatusBadRequest,map[string]string{
			"status":"Error hashingthe password",
		})
	}
	affected,err := userDb.Create()
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusOK, map[string]string{
			"status":"error database",
		})
	}
	if affected == 0 {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"status":"error database",
		})
	}
	userResp := &resmodel.UserResp{}
	userResp.ID = userDb.ID
	userResp.Email = userDb.Email
	token , err := userResp.CreateJwtToken()
	if err != nil {
		log.Println("Error creating jwt token ", err)
		return c.JSON(http.StatusInternalServerError,map[string]string{
			"status" : "Error creating jwt token!",
		})
	}
	defer c.Request().Body.Close()
	return c.JSON(http.StatusOK,map[string]string{
		"token":token,
	})
}

