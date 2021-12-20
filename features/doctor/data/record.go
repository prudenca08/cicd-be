package data

import (
	"finalproject/features/doctor"
	"time"

	"gorm.io/gorm"
)

type Doctor struct {
	gorm.Model
	ID              int `gorm:"primary_key"`
	DoctorSessionID int
	Username        string `gorm:"unique"`
	Password        string
	Name            string
	NIP             string
	Experience      string
	Specialist      string
	Room            string
	Phone_Number    string
	Status          string
	Token           string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

func toDomain(doc Doctor) doctor.Domain {
	return doctor.Domain{
		ID:              doc.ID,
		DoctorSessionID: doc.DoctorSessionID,
		Username:        doc.Username,
		Password:        doc.Password,
		Name:            doc.Name,
		NIP:             doc.NIP,
		Experience:      doc.Experience,
		Specialist:      doc.Specialist,
		Room:            doc.Room,
		Phone_Number:    doc.Phone_Number,
		Status:          doc.Status,
		Token:           doc.Token,
		CreatedAt:       doc.CreatedAt,
		UpdatedAt:       doc.UpdatedAt,
	}
}

func fromDomain(domain doctor.Domain) Doctor {
	return Doctor{
		ID:              domain.ID,
		DoctorSessionID: domain.DoctorSessionID,
		Username:        domain.Username,
		Password:        domain.Password,
		Name:            domain.Name,
		NIP:             domain.NIP,
		Experience:      domain.Experience,
		Specialist:      domain.Specialist,
		Room:            domain.Room,
		Phone_Number:    domain.Phone_Number,
		Status:          domain.Status,
		Token:           domain.Token,
		CreatedAt:       domain.CreatedAt,
		UpdatedAt:       domain.UpdatedAt,
	}
}
func toDomainUpdate(doc Doctor) doctor.Domain {
	return doctor.Domain{
		ID:              doc.ID,
		DoctorSessionID: doc.DoctorSessionID,
		Username:        doc.Username,
		Password:        doc.Password,
		Name:            doc.Name,
		NIP:             doc.NIP,
		Experience:      doc.Experience,
		Specialist:      doc.Specialist,
		Room:            doc.Room,
		Phone_Number:    doc.Phone_Number,
		Status:          doc.Status,
		Token:           doc.Token,
		CreatedAt:       doc.CreatedAt,
		UpdatedAt:       doc.UpdatedAt,
	}
}
