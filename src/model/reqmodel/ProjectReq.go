package reqmodel

import "github.com/satori/go.uuid"

type ProjectStartReq struct {
	State		string		`json:"state"`
	City		string		`json:"city"`
}

type ProjectBasicInfoReq struct {
	ID			uuid.UUID	`json:"id"`
	Title		string		`json:"title"`
	Time		string		`json:"time"`
	EndDate		string		`json:"enddate"`
	Description string		`json:"description"`
	Goal        string		`json:"goal"`
	ImgId		uuid.UUID	`json:"imgid"`
}

type ProjectStoryReq struct {
	ID			uuid.UUID			`json:"id"`
	Story		string				`json:"story"`
	Challenge 	[]string 			`json:"challenge"`
	VidId		uuid.UUID			`json:"vidid"`
	Faq			map[string]string	`json:"faq"`
}

type ProjectProgramReq struct {
	ID					uuid.UUID	`json:"id"`
	OperatingProgram 	string		`json:"operatingprogram"`
	CostDesc			string		`json:"costdesc"`
	//زمان لازم برای هر مرحله (برنامه عملیاتی)

}
type ProjectIdReq struct {
	ID					uuid.UUID	`json:"id"`
}

type PreProjectEdit struct {
	ID 			uuid.UUID 			`json:"id"`
	VidId		uuid.UUID			`json:"vidid"`
	ImgId		uuid.UUID			`json:"imgid"`
	Faq			map[string]string	`json:"faq"`
}
