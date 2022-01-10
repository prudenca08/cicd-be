package patientses

type Domain struct {
	ID                int
	AdminID           int
	DoctorID          int
	PatientID         int
	PatientScheduleID int
	Date              string
	Status            string
	Message           string
}

type Service interface {
	AllPatientses() ([]Domain, error)
	Create(pssID int, domain *Domain) (Domain, error)
	Update(admID int, pssID int, domain *Domain) (Domain, error)
	Delete(pssID, id int) (string, error)
	PatientsesByID(id int) (Domain, error)
}
type Repository interface {
	AllPatientses() ([]Domain, error)
	Create(pssID int, domain *Domain) (Domain, error)
	Update(admID int, pssID int, domain *Domain) (Domain, error)
	Delete(pssID, id int) (string, error)
	PatientsesByID(id int) (Domain, error)
}