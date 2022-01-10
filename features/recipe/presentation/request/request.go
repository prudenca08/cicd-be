package request

import "finalproject/features/recipe"

type Recipe struct {
	Day          string `json:"day"`
	Time         string `json:"time"`
	Title        string `json:"title"`
	DetailRecipe string `json:"detailrecipe"`
}

func (req *Recipe) ToDomain() *recipe.Domain{
	return &recipe.Domain{
		Day: req.Day,
		Time: req.Time,
		Title: req.Title,
		DetailRecipe: req.DetailRecipe,
	}
}