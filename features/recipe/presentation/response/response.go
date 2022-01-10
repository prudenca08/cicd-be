package response

import (
	"finalproject/features/recipe"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CreateRecipeResponse struct {
	Message      string `json:"message"`
	ID           int    `json:"id"`
	Day          string `json:"day"`
	Time         string `json:"time"`
	DetailRecipe string `json:"detailrecipe"`
}

type RecipeResponse struct {
	Message string `json:"message"`
	ID int `json:"id"`
	Day string `json:"day"`
	Time string `json:"time"`
	DetailRecipe string `json:"detailrecipe"`
}

type BaseResponse struct {
	Meta struct {
		Status   int      `json:"rc"`
		Message  string   `json:"message"`
		Messages []string `json:"messages,omitempty"`
	} `json:"meta"`
	Data interface{} `json:"data"`
}

func NewSuccessResponse(c echo.Context, param interface{}) error {
	response := BaseResponse{}
	response.Meta.Status = http.StatusOK
	response.Meta.Message = "Success"
	response.Data = param

	return c.JSON(http.StatusOK, response)
}
func NewErrorResponse(c echo.Context, status int, err error) error {
	response := BaseResponse{}
	response.Meta.Status = status
	response.Meta.Message = "Something not right"
	response.Meta.Messages = []string{err.Error()}

	return c.JSON(status, response)
}

func FromDomainCreate(domain recipe.Domain) CreateRecipeResponse{
	return CreateRecipeResponse{
		Message: "Create Recipe Success",
		ID: domain.ID,
		Day: domain.Day,
		Time: domain.Time,
		DetailRecipe: domain.DetailRecipe,
	}
}

func FromDomainUpdateRecipe(domain recipe.Domain) CreateRecipeResponse{
	return CreateRecipeResponse{
		Message :"Create Recipe Success",
		ID: domain.ID,
		Day: domain.Day,
		Time: domain.Time,
		DetailRecipe: domain.DetailRecipe,
	}
}

func FromDomainAllRecipe(domain recipe.Domain) RecipeResponse{
	return RecipeResponse{
		ID: domain.ID,
		Day: domain.Day,
		Time: domain.Time,
		DetailRecipe: domain.DetailRecipe,
	}
}

func FromRecipeListDomain(domain []recipe.Domain)[]RecipeResponse{
	var response []RecipeResponse
	for _, value := range domain {
		response = append(response, FromDomainAllRecipe(value))
	}
	return response
}