package projectdb

import (
	"Bulldog0.1/src/model"
	"log"
	"github.com/satori/go.uuid"
)


func SelectPreProjetByUserId(id uuid.UUID)(projects []ProjectDb,err error){
	rows, err := model.Db.Query("SELECT id,state,city,title,time,end_date,description,goal,story,challenge,faq,operating_program,cost_desc,vid_id,img_id FROM pre_project WHERE user_id = $1",id)
	if err != nil {
		log.Println(err)
		panic(err)
		return
	}
	for rows.Next(){
		project := ProjectDb{}
		err = rows.Scan(&project.ID,&project.State,
			&project.City,&project.Title,
				&project.Time,&project.EndDate,
					&project.Description,&project.Goal,
						&project.Story,&project.Challenge,
							&project.Faq,&project.OperatingProgram,
								&project.CostDesc, &project.VidId,&project.ImgId)
		if err != nil {
			log.Println(err)
			panic(err)
			return
		}
		projects = append(projects,project)
	}
	defer rows.Close()
	return
}

func SelectProjetByUserId(id uuid.UUID)(projects []ProjectDb,err error){
	rows, err := model.Db.Query("SELECT id,city,title,time,end_date,description,goal,story,challenge,faq,operating_program,cost_desc,vid_id,img_id FROM project WHERE user_id = $1",id)
	if err != nil {
		log.Println(err)
		return
	}
	for rows.Next(){
		project := ProjectDb{}
		err = rows.Scan(
			&project.ID,
			&project.City,&project.Title,
			&project.Time,&project.EndDate,
			&project.Description,&project.Goal,
			&project.Story,&project.Challenge,
			&project.Faq,&project.OperatingProgram,
			&project.CostDesc,
			&project.VidId,&project.ImgId)
		if err != nil {
			log.Println(err)
			return
		}
		projects = append(projects,project)
	}
	defer rows.Close()
	return
}