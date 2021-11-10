package projectdb

import (
	"Bulldog0.1/src/model"

)


func (project *ProjectDb) EditProjectDa() (affected int64,err error){
	rows, err := model.Db.Exec("UPDATE project SET vid_id = $1 , img_id = $2 ,faq = $3 WHERE id = $4",project.VidId,
		project.ImgId, project.Faq, project.ID)
	if err != nil {
		return 0,err
	}
	affected , err = rows.RowsAffected()
	if err != nil{
		return 0,err
	}

	return
}

func (preProject *ProjectDb) DeletePreProjectDa() (affected int64,err error){
	rows, err := model.Db.Exec("DELETE FROM pre_project WHERE id = $1",preProject.ID)
	if err != nil {
		return 0,err
	}
	affected , err = rows.RowsAffected()
	if err != nil{
		return 0,err
	}

	return
}