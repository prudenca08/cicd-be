package request

import "finalproject/features/patientses"

type Patientses struct {
	Status string `json:"status"`
	Date    string `json:"date"`
}

func (req *Patientses) ToDomain() *patientses.Domain {
	return &patientses.Domain{
		Date:              req.Date,
		Status:            req.Status,
	}
}
