package request

import "finalproject/features/patientses"

type Patientses struct {

	DoctorID          int    `json:"doctorid"`
	PatientID         int    `json:"patientid"`
	PatientScheduleID int    `json:"patientscheduleid"`
    Symptoms          string `json:"symptoms"`
	Title             string `json:"title"`
	DetailRecipe      string `json:"detailrecipe"`
	Status            string `json:"status"`
	Date              string `json:"date"`

}

func (req *Patientses) ToDomain() *patientses.Domain {
	return &patientses.Domain{

		DoctorID:          req.DoctorID,
		PatientID:         req.PatientID,
		Symptoms:          req.Symptoms,
		Title:             req.Title,
		DetailRecipe:      req.DetailRecipe,
		PatientScheduleID: req.PatientScheduleID,
		Date:              req.Date,
		Status:            req.Status,

	}
}
