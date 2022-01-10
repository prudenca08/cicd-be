package recipe

type Domain struct {
	ID               int
	DoctorID         int
	PatientSessionID int
	Day              string
	Time             string
	Title            string
	DetailRecipe     string
}

type Repository interface {
	AllRecipe() ([]Domain, error)
	Create(recID int, domain *Domain) (Domain, error)
	Update(docID int, recID int, domain *Domain) (Domain, error)
	Delete(recID, id int) (string, error)
	RecipeByID(id int) (Domain, error)
}
type Service interface {
	AllRecipe() ([]Domain, error)
	Create(recID int, domain *Domain) (Domain, error)
	Update(docID int, recID int, domain *Domain) (Domain, error)
	Delete(recID, id int) (string, error)
	RecipeByID(id int) (Domain, error)
}