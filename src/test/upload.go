package handlers

import (
	"crypto/md5"
	"io"
	"net/http"
	"os"
	"path"
	"encoding/hex"
	"github.com/knq/dburl"
	"github.com/labstack/echo"
)

func Upload(cfg AppConfig) func(c echo.Context) error {
	return func(c echo.Context) error {
		// Open source
		source, err := c.FormFile("file")
		if err != nil {
			return err
		}
		src, err := source.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		// Hash source
		md5 := md5.New()
		if _, err = io.Copy(md5, src); err != nil {
			return err
		}
		hash := hex.EncodeToString(md5.Sum(nil))

		// Check if file with such hash exists
		ext := path.Ext(source.Filename)
		fileName := hash + ext
		filePath := cfg.UploadDir + fileName
		if _, err := os.Stat(filePath); os.IsNotExist(err) {

			// If such file doesn't exist, write it to the disk
			dst, err := os.Create(filePath)
			if err != nil {
				return err
			}
			defer dst.Close()
			if _, err = io.Copy(dst, src); err != nil {
				return err
			}
		}

		uploaderID := GetUserID(c)

		db, err := dburl.Open(cfg.DatabaseURL)
		if err != nil {
			return err
		}

		fi, err := models.FileInfoByUploaderIDName(db, uploaderID, fileName)
		if err != nil {
			return err
		}

		// Check if file with such name exist
		if !fi.Exists() {

			// If such file doesn't exist, add it to the database
			fi.UploaderID = uploaderID
			fi.Name = fileName
			fi.Extention = ext
			fi.Insert(db)
		}

		return c.JSON(http.StatusOK, fi)
	}
}
