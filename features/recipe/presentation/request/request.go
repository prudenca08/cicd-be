package request

import "finalproject/features/recipe"

type Recipe struct {
	PatientSessionID int `json:"patientsessionid"`
	Title        string `json:"title"`
	DetailRecipe string `json:"detailrecipe"`
}

func (req *Recipe) ToDomain() *recipe.Domain{
	return &recipe.Domain{
		PatientSessionID: req.PatientSessionID,
		Title: req.Title,
		DetailRecipe: req.DetailRecipe,
	}
}