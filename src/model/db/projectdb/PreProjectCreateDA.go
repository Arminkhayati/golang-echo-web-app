package projectdb

import (
	"github.com/satori/go.uuid"
	"log"
	"Bulldog0.1/src/model"
)



func (preProject *ProjectDb) Create()(affected int64,err error){
	preProject.ID = uuid.NewV4()
	rows,err := model.Db.Exec("INSERT INTO pre_project (id,city,state,status,user_id) values ($1,$2,$3,$4,$5)",
		preProject.ID,preProject.City,preProject.State,preProject.Status,preProject.UserId)
	if err != nil {
		return 0,err
	}
	affected,err = rows.RowsAffected()
	if err != nil{
		return 0,err
	}
	return
}

func (preProject *ProjectDb) UpdateBasicInfo() (affected int64,err error){
	rows, err := model.Db.Exec("UPDATE pre_project SET title = $1 , time = $2," +
		" end_date = $3 , description = $4 , goal = $5 , img_id = $6 WHERE id = $7",
			preProject.Title , preProject.Time ,
				preProject.EndDate , preProject.Description ,
					preProject.Goal , preProject.ImgId , preProject.ID)
	if err != nil {
		return 0,err
	}
	affected , err = rows.RowsAffected()
	if err != nil{
		return 0,err
	}

	return
}


func (preProject *ProjectDb) UpdateStory() (affected int64,err error){
	rows, err := model.Db.Exec("UPDATE pre_project SET story = $1 , challenge = $2," +
		" vid_id = $3 , faq = $4 WHERE id = $5",
		preProject.Story , preProject.Challenge ,
		preProject.VidId , preProject.Faq , preProject.ID)
	if err != nil {
		return 0,err
	}
	affected , err = rows.RowsAffected()
	if err != nil{
		return 0,err
	}

	return
}

func (preProject *ProjectDb) UpdateProgram() (affected int64,err error){
	rows, err := model.Db.Exec("UPDATE pre_project SET operating_program = $1 , cost_desc = $2, status = $3 WHERE id = $4",
		preProject.OperatingProgram , preProject.CostDesc ,preProject.Status,
		preProject.ID )
	if err != nil {
		return 0,err
	}
	affected , err = rows.RowsAffected()
	if err != nil{
		return 0,err
	}

	return
}


func SelectPreProjectById(id uuid.UUID)(exist bool,err error){
	rows, err := model.Db.Query("SELECT ID FROM pre_project WHERE ID = $1",id)
	if err != nil {
		log.Print(err.Error())
		return
	}
	exist = rows.Next()
	defer rows.Close()
	return
}

func SelectProjectById(id uuid.UUID)(exist bool,err error){
	rows, err := model.Db.Query("SELECT ID FROM project WHERE ID = $1",id)
	if err != nil {
		log.Print(err.Error())
		return
	}
	exist = rows.Next()
	defer rows.Close()
	return
}

func MigratePreproject(id uuid.UUID)(affected int64,err error){
	tx , err := model.Db.Begin()
	if  err != nil{
		return 0,err
	}
	rows, err := tx.Exec("INSERT INTO project" +
		"(id,city,title,time,end_date,description,goal,story,challenge,faq,operating_program,cost_desc,user_id,vid_id,img_id)" +
		" Select " +
		"id,city,title,time,end_date,description,goal,story,challenge,faq,operating_program,cost_desc,user_id,vid_id,img_id" +
		" from pre_project WHERE id = $1",id)
	if err != nil {
		tx.Rollback()
		return 0,err
	}
	a , err := rows.RowsAffected()
	if err != nil{
		tx.Rollback()
		return 0,err
	}
	affected = a
	rows, err = tx.Exec("DELETE FROM pre_project where id = $1",id)
	if err != nil {
		tx.Rollback()
		return 0,err
	}
	a , err = rows.RowsAffected()
	if err != nil{
		tx.Rollback()
		return 0,err
	}
	affected += a
	tx.Commit()
	return

}
