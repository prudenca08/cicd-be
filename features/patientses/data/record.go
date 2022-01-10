package data

import (
	"finalproject/features/patientses"

	"gorm.io/gorm"
)

type Patientses struct {
	gorm.Model
	ID int `gorm:"primary_key"`
	AdminID int
	DoctorID int `gorm:"primary_key"`
	PatientID int
	PatientScheduleID int `gorm:"primary_key"`
	Date string
	Status string
}
func toDomain(pss Patientses) patientses.Domain {
	return patientses.Domain{
		ID:                pss.ID,
		AdminID:           pss.AdminID,
		DoctorID:          pss.DoctorID,
		PatientID:         pss.PatientID,
		PatientScheduleID: pss.PatientScheduleID,
		Date:              pss.Date,
		Status:            pss.Status,
	}
}
func fromDomain(domain patientses.Domain) Patientses {
	return Patientses{
		ID:                domain.ID,
		AdminID:           domain.AdminID,
		DoctorID:          domain.DoctorID,
		PatientID:         domain.PatientID,
		PatientScheduleID: domain.PatientScheduleID,
		Date:              domain.Date,
		Status:            domain.Status,
	}
}
func toDomainUpdate(pss Patientses) patientses.Domain{
	return patientses.Domain{
		ID: pss.ID,
		AdminID: pss.AdminID,
		DoctorID: pss.DoctorID,
		PatientID: pss.PatientID,
		PatientScheduleID: pss.PatientScheduleID,
		Date: pss.Date,
		Status: pss.Status,
	}
}
func toDomainList(data []Patientses) []patientses.Domain {
	result := []patientses.Domain{}
	for _, pss := range data{
		result = append(result, toDomain(pss))
	}
	return result
}