package data

import (
	patsesrecord "finalproject/features/patientses/data"
	"finalproject/features/recipe"
	"time"

	"gorm.io/gorm"
)

type Recipe struct {
	gorm.Model
	PatientSessionID int
	ID int
	Title string
	DetailRecipe string
	Patientses patsesrecord.Patientses `gorm:"foreignKey:ID;references:PatientSessionID"`
	CreatedAt time.Time
	UpdatedAt time.Time
	
}

func ToDomain(rec Recipe) recipe.Domain{
	return recipe.Domain{
		PatientSessionID: rec.PatientSessionID,
		ID: rec.ID,
		Title: rec.Title,
		DetailRecipe: rec.DetailRecipe,
		Patientses: patsesrecord.ToDomain(rec.Patientses),
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
}

func fromDomain(domain recipe.Domain) Recipe{
	return Recipe{
		PatientSessionID: domain.PatientSessionID,
		ID: domain.ID,
		Title: domain.Title,
		DetailRecipe: domain.DetailRecipe,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func toDomainUpdate(rec Recipe) recipe.Domain{
	return recipe.Domain{
		PatientSessionID: rec.PatientSessionID,
		ID: rec.ID,
		Title: rec.Title,
		DetailRecipe: rec.DetailRecipe,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
}

func toDomainList(data []Recipe) []recipe.Domain{
	result := []recipe.Domain{}
	for _, rec := range data {
		result = append(result, ToDomain(rec))
	}
	return result
}