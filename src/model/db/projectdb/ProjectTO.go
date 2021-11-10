package projectdb




import (
	"github.com/satori/go.uuid"
)

type ProjectDb struct {
	ID 					uuid.UUID
	State 				string
	City 				string
	Title			    string
	Time				string
	EndDate 			string
	Description 		string
	Goal 				string
	Story 				string
	Challenge 			string  // JSON
	Faq 				string	// JSON
	OperatingProgram	string
	CostDesc 			string
	UserId 				uuid.UUID
	VidId 				uuid.UUID
	ImgId 				uuid.UUID
	Status				string
}

/*
	ye json migiram time tabdil mikonam ba formati ke mikham
	1- time.RFC3339
	ya
	2- "2006-01-02"
	midam be db

	time.Parse(time.RFC3339,"string time" )  Json ke pars kardam be time tabdil mikonam
	time.Format("2006-01-02")  ye string be form Date dar DB mide
	time.Format(time.RFC3339)  ye string be form Timestamp dar DB mide

 */