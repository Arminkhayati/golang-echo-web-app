package handler

import (
	"crypto/md5"
	"io"
	"log"
	"encoding/hex"
	"path"
	"net/http"
	"github.com/labstack/echo"
	"os"
	"Bulldog0.1/src/model/db/fileinfodb"
	"github.com/satori/go.uuid"
	"Bulldog0.1/src/model/resmodel"
	"Bulldog0.1/src/model/db/userdb"
)

func Upload (c echo.Context) error {
	/******************Open File******************************/
	file := fileinfodb.NewFileInfoDb()
	source, err := c.FormFile("file")
	if err != nil {
		log.Printf("failed reading File: %s",err)
		return c.JSON(http.StatusBadRequest,map[string]string{
			"status":"Error Reading File",
		})
	}
	src, err := source.Open()
	if err != nil {
		log.Printf("failed Open File: %s",err)
		return c.JSON(http.StatusBadRequest,map[string]string{
			"status":"Error Open File",
		})
	}
	defer src.Close()
	/**********************************************************/

	/*****************Check file Extention*******************/
	ext := path.Ext(source.Filename)
	file.FileType = fileinfodb.GetFileType(ext)
	if file.FileType == "bad"{
		log.Printf("Error file type\n")
		return c.JSON(http.StatusOK, map[string]string{
			"status":"Error file type",
		})
	}
	/**********************************************************/

	/******************Copy File to MD5 and Destination /var/www/vartow/xxx.mp3************************/
	md5 := md5.New()
	//file.Address += source.Filename
	source.Filename = file.Address + source.Filename
	dst, err := os.Create(source.Filename)
	if err != nil {
		log.Printf("Error creating file!\n")
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"status":"Error creating file",
		})
	}
	writer := io.MultiWriter(md5,dst)
	if _, err := io.Copy(writer, src); err != nil {
		log.Printf("Error copying file!\n")
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"status":"Error copying file",
		})
	}
	hash := hex.EncodeToString(md5.Sum(nil))
	dst.Close()
	/**********************************************************/

	/*******************Check If file exists delete the uploaded file**********************/
	hashedName := hash + ext
	file.Address += hashedName
	if _, err := os.Stat(file.Address);  err == nil{
		err = os.Remove(source.Filename)
		if err != nil{
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"status":"Error Removing file",
			})
		}
	}else{
		err = os.Rename(source.Filename,file.Address)
		if err != nil{
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"status":"Error Rename file",
			})
		}
	}
	/**********************************************************/

	/*************************Insert TO File_Info table If file not exists and User Exists************************/
	file.Id = uuid.NewV4()
	file.UploaderId = resmodel.GetUserID(c)//check user exists
	exist, err := userdb.SelectUserByID(file.UploaderId)
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

	file.Extention = ext
	exists, err := fileinfodb.FileExistsByUploaderId(file.UploaderId, file.Address)
	if err != nil {
		return err
	}
	// Check if file with such name exist
	if exists{
		// If such file doesn't exist, add it to the database
		return c.JSON(http.StatusOK, map[string]string{
			"status":"File Exists",
		})
	}
	affected,err := file.Create()
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
	return c.JSON(http.StatusOK, map[string]uuid.UUID{
		"id": file.Id,
	})

}