package doctor

import "time"

type Domain struct {
	ID              int
	DoctorSessionID int
	Username        string
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

type Service interface {
	Register(domain *Domain) (Domain, error)
	Login(username, password string) (Domain, error)
	Update(docID int, domain *Domain) (Domain, error)
	DoctorByID(id int) (Domain, error)
	Delete(docID, id int) (string, error)
}

type Repository interface {
	Register(domain *Domain) (Domain, error)
	Login(username, password string) (Domain, error)
	Update(docID int, domain *Domain) (Domain, error)
	DoctorByID(id int) (Domain, error)
	Delete(docID, id int) (string, error)
}
