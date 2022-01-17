package data

import (
	"finalproject/features/recipe"
	"finalproject/features/recipe/bussiness"

	"gorm.io/gorm"
)

type MysqlRecipeRepository struct {
	Conn *gorm.DB
}

func NewMysqlRecipeRepository(conn *gorm.DB) recipe.Repository{
	return &MysqlRecipeRepository{
		Conn: conn,
	}
}

func (rep *MysqlRecipeRepository) Create(recID int, domain *recipe.Domain) (recipe.Domain,error){
	rec := fromDomain(*domain)

	rec.ID = recID
	result := rep.Conn.Create(&rec)
	if result.Error != nil {
		return recipe.Domain{}, result.Error
	}
	return ToDomain(rec), nil
}

func (rep *MysqlRecipeRepository) Update(docID int, recID int, domain *recipe.Domain)(recipe.Domain, error){
	recipeUpdate := fromDomain(*domain)

	recipeUpdate.ID = recID
	result := rep.Conn.Where("id = ?", recID).Updates(&recipeUpdate)

	if result.Error != nil {
		return recipe.Domain{}, bussiness.ErrNotFound
	}
	return toDomainUpdate(recipeUpdate), nil
}

func (rep *MysqlRecipeRepository) Delete(recID int, id int)(string , error){
	rec := Recipe{}
	find := rep.Conn.Where("id = ?", id).First(&rec)
	if find.Error != nil {
		return "", bussiness.ErrUnathorized
	}
	err := rep.Conn.Delete(&rec, "id = ?", id). Error
	if err != nil {
		return "", bussiness.ErrNotFound
	}
	return "Recipe has been delete", nil
}

func (rep *MysqlRecipeRepository) AllRecipe() ([]recipe.Domain, error){
	var doc []Recipe
	
	result := rep.Conn.Preload("Patientses").Find(&doc)

	// ss, _ := json.MarshalIndent(doc, "", " ")
	// fmt.Println(string(ss))

	if result.Error != nil {
		return []recipe.Domain{}, result.Error
	}
	return toDomainList(doc), nil
}
func (rep *MysqlRecipeRepository) RecipeByID(id int) (recipe.Domain, error){
	var doc Recipe
	result := rep.Conn.Where("id = ?", id).First(&doc)
	if result.Error !=nil {
		return recipe.Domain{}, result.Error
	}
	return ToDomain(doc),nil
}