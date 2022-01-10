package bussiness

import (
	"finalproject/features/recipe"
)

type serviceRecipe struct {
	recipeRepository recipe.Repository
}
func NewServiceRecipe(repoRecipe recipe.Repository) recipe.Service{
	return &serviceRecipe{
		recipeRepository: repoRecipe,
	}
}

func (serv *serviceRecipe) AllRecipe() ([]recipe.Domain, error){
	result, err := serv.recipeRepository.AllRecipe()

	if err != nil {
		return []recipe.Domain{}, err
	}
	return result, nil
}

func (serv *serviceRecipe) Create(orgID int, domain *recipe.Domain)(recipe.Domain, error){
	result, err := serv.recipeRepository.Create(orgID, domain)
	if err != nil {
		return recipe.Domain{}, err
	}
	return result, nil
}

func (serv *serviceRecipe) Update(orgID int,prodID int, domain *recipe.Domain)(recipe.Domain, error){
	result, err := serv.recipeRepository.Update(orgID,prodID,domain)
	if err !=nil {
		return recipe.Domain{}, ErrDuplicateData
	}
	return result, nil
}

func (serv *serviceRecipe) Delete(orgID int, id int)(string, error){
	result, err := serv.recipeRepository.Delete(orgID, id)
	if err != nil {
		return "", ErrNotFound
	}
	return result, nil
}


func (serv *serviceRecipe) RecipeByID(id int) (recipe.Domain, error){
	result, err := serv.recipeRepository.RecipeByID(id)
	if err != nil {
		return recipe.Domain{}, err
	}
	return result, nil
}