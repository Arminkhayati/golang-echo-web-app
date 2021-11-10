package handler

import (
	"github.com/labstack/echo"
	"net/http"
	"Bulldog0.1/src/model/resmodel"
	"Bulldog0.1/src/model/db/userdb"
	"log"
	"Bulldog0.1/src/model/db/projectdb"
)

func UserPreProject(c echo.Context)error{
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
	/**************************************************************************************/
	projects ,err:= projectdb.SelectPreProjetByUserId(userID)
	if err != nil{
		log.Println(err)
		return c.JSON(http.StatusInternalServerError,map[string]string{
			"status":"Error Databasae",
		})
	}
	projectResp := resmodel.PreProjectResp{}
	for i := 0 ; i < len(projects);i++{
		project := resmodel.PreProject{}
		project.ID = projects[i].ID
		project.State = projects[i].State
		project.City = projects[i].City
		project.Title = projects[i].Title
		project.Time = projects[i].Time
		project.EndDate = projects[i].EndDate
		project.Description = projects[i].Description
		project.Goal = projects[i].Goal
		project.Story = projects[i].Story
		project.Challenge = projects[i].Challenge
		project.Faq = projects[i].Faq
		project.OperatingProgram = projects[i].OperatingProgram
		project.CostDesc = projects[i].CostDesc
		project.VidId = projects[i].VidId
		project.ImgId = projects[i].ImgId
		projectResp.Preprojects = append(projectResp.Preprojects,project)
	}

	return c.JSON(http.StatusOK,projectResp)
}


func UserProject(c echo.Context)error{
	userID := resmodel.GetUserID(c)
	log.Println(userID)
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
	/**************************************************************************************/
	projects ,err:= projectdb.SelectProjetByUserId(userID)
	if err != nil{
		log.Println(err)
		return c.JSON(http.StatusInternalServerError,map[string]string{
			"status":"Error Databasae",
		})
	}
	projectResp := resmodel.ProjectResp{}
	for i := 0 ; i < len(projects);i++{
		project := resmodel.Project{}
		project.ID = projects[i].ID
		project.City = projects[i].City
		project.Title = projects[i].Title
		project.Time = projects[i].Time
		project.EndDate = projects[i].EndDate
		project.Description = projects[i].Description
		project.Goal = projects[i].Goal
		project.Story = projects[i].Story
		project.Challenge = projects[i].Challenge
		project.Faq = projects[i].Faq
		project.OperatingProgram = projects[i].OperatingProgram
		project.CostDesc = projects[i].CostDesc
		project.VidId = projects[i].VidId
		project.ImgId = projects[i].ImgId
		projectResp.Projects = append(projectResp.Projects,project)
	}
	return c.JSON(http.StatusOK,projectResp)
}
