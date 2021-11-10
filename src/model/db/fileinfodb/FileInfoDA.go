package fileinfodb

import (
	"github.com/satori/go.uuid"
	"Bulldog0.1/src/model"
	"log"
)


func (file *fileInfoDb) Create()(affected int64,err error){
	rows,err :=model.Db.Exec("INSERT INTO file_info (id,address,extention,ftype,uploader_id) values ($1,$2,$3,$4,$5)",
		file.Id,file.Address,file.Extention,file.FileType,file.UploaderId)
	if err != nil {
		return 0,err
	}
	affected,err = rows.RowsAffected()
	if err != nil{
		return 0,err
	}
	return
}

func FileExists(id uuid.UUID) (info struct{ Id 		uuid.UUID
											Exist 		bool
											FileType 	string
											Err 		error}){
	rows, err := model.Db.Query("SELECT Id,ftype FROM file_info WHERE Id = $1",id)
	if err != nil {
		log.Print(err.Error())
		info.Err = err
		return
	}
	info.Exist = rows.Next()
	err = rows.Scan(&info.Id, &info.FileType)
	if err != nil {
		log.Print(err.Error())
		info.Err = err
		return
	}
	info.Err = err
	rows.Close()
	return
}

func FileExistsByUploaderId(uploaderId uuid.UUID, address string)(exists bool,err error){
	rows, err := model.Db.Query("SELECT Id FROM file_info WHERE uploader_id = $1 AND address = $2",uploaderId,address)
	if err != nil {
		log.Print(err.Error())
		return
	}
	exists = rows.Next()
	return
}