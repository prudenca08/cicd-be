package request

import "finalproject/features/patientses"

type Patientses struct {
	AdminID int `json:"adminid"`
	DoctorID int `json:"doctorid"`
	PatientID int `json:"patientid"`
	PatientScheduleID int `json:"patientscheduleid"`
	Title            string `json:"title"`
	DetailRecipe     string `json:"detailrecipe"`
	Status string `json:"status"`
	Date    string `json:"date"`
}

func (req *Patientses) ToDomain() *patientses.Domain {
	return &patientses.Domain{
		AdminID:            req.AdminID,
		DoctorID:           req.DoctorID,
		PatientID:          req.PatientID,
		Title:              req.Title,
		DetailRecipe:       req.DetailRecipe ,
		PatientScheduleID:  req.PatientScheduleID,
		Date:               req.Date,
		Status:             req.Status,
	}
}
