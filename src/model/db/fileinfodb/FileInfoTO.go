package fileinfodb

import "github.com/satori/go.uuid"

type fileInfoDb struct {
	Id         uuid.UUID
	Address    string
	Extention  string
	FileType   string
	UploaderId uuid.UUID
}


func NewFileInfoDb() fileInfoDb{
	file := fileInfoDb{}
	file.Address = "/var/www/vartow/uploads/"
	return file
}


func GetFileType(ext string) (string){
	switch ext {
		case ".mp4"  	 : return "video"
		case ".mp3"		 : return "audio"
		case ".jpg",".png" : return "image"
	}
	return "bad"
}