package handler

import (
	"github.com/labstack/echo"
	"Bulldog0.1/src/model/reqmodel"
	"io/ioutil"
	"net/http"
	"encoding/json"
	"log"
	"Bulldog0.1/src/model/db/projectdb"
	"github.com/satori/go.uuid"
	"time"
	"Bulldog0.1/src/model/db/fileinfodb"
	"Bulldog0.1/src/model/db/userdb"
	"Bulldog0.1/src/model/resmodel"
)

func Start(c echo.Context)error{

	userID := resmodel.GetUserID(c)

	projectReq := reqmodel.ProjectStartReq{}
	body , err := ioutil.ReadAll(c.Request().Body)
	if err != nil{
		log.Printf("failed reading the body for add Start Info: %s",err)
		return c.JSON(http.StatusBadRequest,map[string]string{
			"status":"Error Reading Request Body",
		})
	}
	err = json.Unmarshal(body,&projectReq)
	if err != nil{
		log.Printf("failed unmarshaling: %s",err)
		return c.JSON(http.StatusBadRequest,map[string]string{
			"status":"Error Unmarshaling",
		})
	}
	projectDb := projectdb.ProjectDb{}
	projectDb.State = projectReq.State
	projectDb.City = projectReq.City
	projectDb.Status = "noncomplete"
	projectDb.UserId = userID//check user by id************************
	exist, err := userdb.SelectUserByID(projectDb.UserId)
	if err != nil{
		log.Println(err)
		return c.JSON(http.StatusInternalServerError,map[string]string{
			"status":"Error Databasae",
		})
	}
	if !exist{
		return c.JSON(http.StatusOK ,map[string]string{
			"status":"User Not Exists.",
		})
	}
	affected,err := projectDb.Create()
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
	defer c.Request().Body.Close()
	return c.JSON(http.StatusOK,map[string]uuid.UUID{
		"Id":projectDb.ID,
	})
}


func BasicInfo(c echo.Context)error{
	userID := resmodel.GetUserID(c)
	exist, err := userdb.SelectUserByID(userID)
	if err != nil{
		log.Println(err)
		return c.JSON(http.StatusInternalServerError,map[string]string{
			"status":"Error Databasae",
		})
	}
	if !exist{
		return c.JSON(http.StatusOK ,map[string]string{
			"status":"User Not Exists.",
		})
	}

	projectReq := reqmodel.ProjectBasicInfoReq{}
	body , err := ioutil.ReadAll(c.Request().Body)
	if err != nil{
		log.Printf("failed reading the body for add Basic Info: %s",err)
		return c.JSON(http.StatusBadRequest,map[string]string{
			"status":"Error Reading Request Body",
		})
	}
	err = json.Unmarshal(body,&projectReq)
	if err != nil{
		log.Printf("failed unmarshaling: %s",err)
		return c.JSON(http.StatusBadRequest,map[string]string{
			"status":"Error Unmarshaling",
		})
	}
	/************/
	projectDb := projectdb.ProjectDb{}
	projectDb.ID = projectReq.ID
	exist,err = projectdb.SelectPreProjectById(projectDb.ID)
	if err != nil{
		log.Println(err)
		return c.JSON(http.StatusInternalServerError,map[string]string{
			"status":"Error Databasae",
		})
	}
	if !exist{
		return c.JSON(http.StatusOK ,map[string]string{
			"status":"Project Not Exists.",
		})
	}
	projectDb.Title = projectReq.Title
	ptime,err := time.Parse(time.RFC3339,projectReq.Time )
	if err != nil{
		log.Printf("failed processing time: %s",err)
		return c.JSON(http.StatusBadRequest,map[string]string{
			"status":"Time Bad Format",
		})
	}
	projectDb.Time	= ptime.Format(time.RFC3339)
	pdate ,err := time.Parse(time.RFC3339,projectReq.EndDate)
	if err != nil{
		log.Printf("failed processing date: %s",err)
		return c.JSON(http.StatusBadRequest,map[string]string{
			"status":"Date Bad Format",
		})
	}
	projectDb.EndDate = pdate.Format("2006-01-02")
	projectDb.Description = projectReq.Description
	projectDb.Goal = projectReq.Goal
	info := fileinfodb.FileExists(projectReq.ImgId)
	if info.Err != nil{
		log.Println(err)
		return c.JSON(http.StatusInternalServerError,map[string]string{
			"status":"Error Databasae",
		})
	}
	if !info.Exist{
		return c.JSON(http.StatusBadRequest,map[string]string{
			"status":"File Not Exists",
		})
	}
	if info.FileType != "image"{
		return c.JSON(http.StatusBadRequest,map[string]string{
			"status":"Wrong type",
		})
	}
	projectDb.ImgId = projectReq.ImgId
	affected,err := projectDb.UpdateBasicInfo()
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
	/**************/
	defer c.Request().Body.Close()
	return c.JSON(http.StatusOK,map[string]string{
		"status" : "done babe",
	})
}


func Story(c echo.Context)error{
	userID := resmodel.GetUserID(c)
	exist, err := userdb.SelectUserByID(userID)
	if err != nil{
		log.Println(err)
		return c.JSON(http.StatusInternalServerError,map[string]string{
			"status":"Error Databasae",
		})
	}
	if !exist{
		return c.JSON(http.StatusOK ,map[string]string{
			"status":"User Not Exists.",
		})
	}


	projectReq := reqmodel.ProjectStoryReq{}
	body , err := ioutil.ReadAll(c.Request().Body)
	if err != nil{
		log.Printf("failed reading the body for add Story: %s",err)
		return c.JSON(http.StatusBadRequest,map[string]string{
			"status":"Error Reading Request Body",
		})
	}
	err = json.Unmarshal(body,&projectReq)
	if err != nil{
		log.Printf("failed unmarshaling: %s",err)
		return c.JSON(http.StatusBadRequest,map[string]string{
			"status":"Error Unmarshaling",
		})
	}
	/*********************************/
	projectDb := projectdb.ProjectDb{}
	projectDb.ID = projectReq.ID
	exist, err = projectdb.SelectPreProjectById(projectDb.ID)
	if err != nil{
		log.Println(err)
		return c.JSON(http.StatusInternalServerError,map[string]string{
			"status":"Error Databasae",
		})
	}
	if !exist{
		return c.JSON(http.StatusOK ,map[string]string{
			"status":"Project Not Exists.",
		})
	}
	projectDb.Story = projectReq.Story
	info := fileinfodb.FileExists(projectReq.VidId)
	if info.Err != nil{
		log.Println(err)
		return c.JSON(http.StatusInternalServerError,map[string]string{
			"status":"Error Databasae",
		})
	}
	if !info.Exist{
		return c.JSON(http.StatusBadRequest,map[string]string{
			"status":"File Not Exists",
		})
	}
	if info.FileType != "video"{
		return c.JSON(http.StatusBadRequest,map[string]string{
			"status":"Wrong type",
		})
	}
	projectDb.VidId = projectReq.VidId
	challenges := projectReq.Challenge
	b ,err := json.Marshal(challenges)
	if err != nil{
		log.Printf("failed marshaling challenges: %s",err)
		return c.JSON(http.StatusBadRequest,map[string]string{
			"status":"failed marshaling challenges",
		})
	}
	projectDb.Challenge = string(b)
	faqs := projectReq.Faq
	b , err = json.Marshal(faqs)
	if err != nil{
		log.Printf("failed marshaling faqs: %s",err)
		return c.JSON(http.StatusBadRequest,map[string]string{
			"status":"failed marshaling faqs",
		})
	}
	projectDb.Faq = string(b)

	/*********************************/

	affected,err := projectDb.UpdateStory()
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

	defer c.Request().Body.Close()
	return c.JSON(http.StatusOK,map[string]string{
		"status" : "done babe",
	})
}


func Program(c echo.Context)error{
	userID := resmodel.GetUserID(c)
	exist, err := userdb.SelectUserByID(userID)
	if err != nil{
		log.Println(err)
		return c.JSON(http.StatusInternalServerError,map[string]string{
			"status":"Error Databasae",
		})
	}
	if !exist{
		return c.JSON(http.StatusOK ,map[string]string{
			"status":"User Not Exists.",
		})
	}

	projectReq := reqmodel.ProjectProgramReq{}
	body , err := ioutil.ReadAll(c.Request().Body)
	if err != nil{
		log.Printf("failed reading the body for add Story: %s",err)
		return c.JSON(http.StatusBadRequest,map[string]string{
			"status":"Error Reading Request Body",
		})
	}
	err = json.Unmarshal(body,&projectReq)
	if err != nil{
		log.Printf("failed unmarshaling: %s",err)
		return c.JSON(http.StatusBadRequest,map[string]string{
			"status":"Error Unmarshaling",
		})
	}
	/*********************************/

	projectDb := projectdb.ProjectDb{}
	projectDb.ID = projectReq.ID//check project by id**********************
	exist, err = projectdb.SelectPreProjectById(projectDb.ID)
	if err != nil{
		log.Println(err)
		return c.JSON(http.StatusInternalServerError,map[string]string{
			"status":"Error Databasae",
		})
	}
	if !exist{
		return c.JSON(http.StatusOK ,map[string]string{
			"status":"Project Not Exists.",
		})
	}
	projectDb.OperatingProgram = projectReq.OperatingProgram
	projectDb.CostDesc = projectReq.CostDesc
	projectDb.Status = "awaiting"
	affected,err := projectDb.UpdateProgram()
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

	defer c.Request().Body.Close()
	return c.JSON(http.StatusOK,map[string]string{
		"status" : "done babe",
	})
}


func Confirm(c echo.Context)error{
	userID := resmodel.GetUserID(c)
	exist, err := userdb.SelectUserByID(userID)
	if err != nil{
		log.Println(err)
		return c.JSON(http.StatusInternalServerError,map[string]string{
			"status":"Error Databasae",
		})
	}
	if !exist{
		return c.JSON(http.StatusOK ,map[string]string{
			"status":"User Not Exists.",
		})
	}

	projectReq := reqmodel.ProjectIdReq{}
	body , err := ioutil.ReadAll(c.Request().Body)
	if err != nil{
		log.Printf("failed reading the body for add Story: %s",err)
		return c.JSON(http.StatusBadRequest,map[string]string{
			"status":"Error Reading Request Body",
		})
	}
	err = json.Unmarshal(body,&projectReq)
	if err != nil{
		log.Printf("failed unmarshaling: %s",err)
		return c.JSON(http.StatusBadRequest,map[string]string{
			"status":"Error Unmarshaling",
		})
	}
	/*********************************/
	exist, err = projectdb.SelectPreProjectById(projectReq.ID)
	if err != nil{
		log.Println(err)
		return c.JSON(http.StatusInternalServerError,map[string]string{
			"status":"Error Databasae",
		})
	}
	if !exist{
		return c.JSON(http.StatusOK ,map[string]string{
			"status":"Project Not Exists.",
		})
	}


	/***********************/
	affected,err :=projectdb.MigratePreproject(projectReq.ID)
	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusOK, map[string]string{
			"status":"error database",
		})
	}
	if affected == 0 {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"status":"error database aff",
		})
	}
	/*************************/

	defer c.Request().Body.Close()
	return c.JSON(http.StatusOK,map[string]string{
		"status" : "done babe",
	})
}
