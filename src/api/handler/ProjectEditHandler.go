package handler

import (
	"github.com/labstack/echo"
	"net/http"
	"Bulldog0.1/src/model/resmodel"
	"Bulldog0.1/src/model/db/userdb"
	"log"
	"io/ioutil"
	"encoding/json"
	"Bulldog0.1/src/model/reqmodel"
	"Bulldog0.1/src/model/db/projectdb"
	"Bulldog0.1/src/model/db/fileinfodb"
)

func EditPreProject(c echo.Context) error {

	/**********Edit FAQ , Image, Video***********/
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
	projectReq := reqmodel.PreProjectEdit{}
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
	/**************************************************************************************/
	projectDb := projectdb.ProjectDb{}
	projectDb.ID = projectReq.ID
	exist, err = projectdb.SelectProjectById(projectDb.ID)
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
	/*****check image****/
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
	/**********check video**********/
	info = fileinfodb.FileExists(projectReq.VidId)
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
	faqs := projectReq.Faq
	b , err := json.Marshal(faqs)
	if err != nil{
		log.Printf("failed marshaling faqs: %s",err)
		return c.JSON(http.StatusBadRequest,map[string]string{
			"status":"failed marshaling faqs",
		})
	}
	projectDb.Faq = string(b)
	/****************************************************************************************/
	affected,err := projectDb.EditProjectDa()
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


func DeletePreProject(c echo.Context) error {

	/**********Delete By ID***********/
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
	/**************************************************************************************/
	projectDb := projectdb.ProjectDb{}
	projectDb.ID = projectReq.ID
	affected,err := projectDb.DeletePreProjectDa()
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
