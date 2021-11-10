package userdb

import (
	"github.com/satori/go.uuid"
	"log"
	"Bulldog0.1/src/model"
)

func (user *UserTo) Create()(affected int64,err error){
	user.ID = uuid.NewV4()
	rows,err :=model.Db.Exec("INSERT INTO app_user (ID,email, email_confirmed,password)" +
		"values ($1,$2,$3,$4)",
		user.ID,user.Email, false,user.Password)
	if err != nil {
		return 0,err
	}
	affected,err = rows.RowsAffected()
	if err != nil{
		return 0,err
	}
	return
}

func (user *UserTo) Update()(err error){

	_, err = model.Db.Exec("UPDATE app_user SET email_confirmed = $1, password = $2 WHERE Id =$3",
		user.EmailConfirmed,user.Password,user.ID)
	if err != nil {
		return
	}
	return
}


func (usr *UserTo) SelectByEmail()(exist bool,err error){
	rows, err := model.Db.Query("SELECT ID,password FROM app_user WHERE email = $1",usr.Email)
	if err != nil {
		log.Print(err.Error())
		return
	}

	exist = rows.Next()
	if !exist {
		return
	}

	err = rows.Scan(&usr.ID,&usr.Password)
	if err != nil{
		log.Print(err.Error())
		return
	}
	defer rows.Close()
	return
}


func SelectUserByID(id uuid.UUID)(exist bool,err error){
	rows, err := model.Db.Query("SELECT ID FROM app_user WHERE ID = $1",id)
	if err != nil {
		log.Print(err.Error())
		return
	}

	exist = rows.Next()
	if !exist {
		return
	}
	defer rows.Close()
	return
}
