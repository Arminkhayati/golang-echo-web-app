package resmodel

import "github.com/satori/go.uuid"

type PreProjectResp struct {
	Preprojects	[]PreProject		`json:"preprojects"`
}

type PreProject struct {
	ID 					uuid.UUID	`json:"id"`
	State 				string		`json:"state"`
	City 				string 		`json:"city"`
	Title			    string		`json:"title"`
	Time				string		`json:"time"`
	EndDate 			string		`json:"endDate"`
	Description 		string		`json:"description"`
	Goal 				string		`json:"goal"`
	Story 				string		`json:"story"`
	Challenge 			string		`json:"challenge"`
	Faq 				string		`json:"faq"`
	OperatingProgram	string		`json:"operating_program"`
	CostDesc 			string		`json:"costDesc"`
	VidId 				uuid.UUID	`json:"vidId"`
	ImgId 				uuid.UUID	`json:"imgId"`
	Status				string		`json:"status"`
}


type ProjectResp struct {
	Projects	[]Project		`json:"projects"`
}
type Project struct {
	ID 					uuid.UUID	`json:"id"`
	City 				string 		`json:"city"`
	Title			    string		`json:"title"`
	Time				string		`json:"time"`
	EndDate 			string		`json:"endDate"`
	Description 		string		`json:"description"`
	Goal 				string		`json:"goal"`
	Story 				string		`json:"story"`
	Challenge 			string		`json:"challenge"`
	Faq 				string		`json:"faq"`
	OperatingProgram	string		`json:"operating_program"`
	CostDesc 			string		`json:"costDesc"`
	VidId 				uuid.UUID	`json:"vidId"`
	ImgId 				uuid.UUID	`json:"imgId"`
}