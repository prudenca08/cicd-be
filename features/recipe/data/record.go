package data

import (
	"finalproject/features/recipe"

	"gorm.io/gorm"
)

type Recipe struct {
	gorm.Model
	PatientSessionID int
	ID int
	DoctorID int
	Day string
	Time string
	Title string
	DetailRecipe string
}

func toDomain(rec Recipe) recipe.Domain{
	return recipe.Domain{
		PatientSessionID: rec.PatientSessionID,
		ID: rec.ID,
		DoctorID: rec.DoctorID,
		Day : rec.Day,
		Time: rec.Time,
		Title: rec.Title,
		DetailRecipe: rec.DetailRecipe,
	}
}

func fromDomain(domain recipe.Domain) Recipe{
	return Recipe{
		PatientSessionID: domain.PatientSessionID,
		ID: domain.ID,
		DoctorID: domain.DoctorID,
		Day: domain.Day,
		Time: domain.Time,
		Title: domain.Title,
		DetailRecipe: domain.DetailRecipe,
	}
}

func toDomainUpdate(rec Recipe) recipe.Domain{
	return recipe.Domain{
		PatientSessionID: rec.PatientSessionID,
		ID: rec.ID,
		DoctorID: rec.DoctorID,
		Day: rec.Day,
		Time: rec.Time,
		Title: rec.Title,
		DetailRecipe: rec.DetailRecipe,
	}
}

func toDomainList(data []Recipe) []recipe.Domain{
	result := []recipe.Domain{}
	for _, rec := range data {
		result = append(result, toDomain(rec))
	}
	return result
}