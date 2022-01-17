package request

import "finalproject/features/patientses"

type Patientses struct {
	AdminID int `json:"adminid"`
	DoctorID int `json:"doctorid"`
	PatientID int `json:"patientid"`
	PatientScheduleID int `json:"patientscheduleid"`
	Status string `json:"status"`
	Date    string `json:"date"`
}

func (req *Patientses) ToDomain() *patientses.Domain {
	return &patientses.Domain{
		AdminID:           req.AdminID,
		DoctorID:          req.DoctorID,
		PatientID:         req.PatientID,
		PatientScheduleID: req.PatientScheduleID,
		Date:              req.Date,
		Status:            req.Status,
	}
}
