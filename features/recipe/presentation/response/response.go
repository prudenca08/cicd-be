package response

import (
	"finalproject/features/recipe"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

type CreateRecipeResponse struct {
	Message      string `json:"message"`
	PatientSessionID int `json:"patientsessionID"`
	ID           int    `json:"id"`
	Title string `json:"title"`
	DetailRecipe string `json:"detailrecipe"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"update_at"`
}

type RecipeResponse struct {
	Message string `json:"message"`
	ID int `json:"id"`
	Title string `json:"title"`
	PatientSessionID int `json:"patientsessionid"`
	DetailRecipe string `json:"detailrecipe"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
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
		PatientSessionID: domain.PatientSessionID,
		Title: domain.Title,
		DetailRecipe: domain.DetailRecipe,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func FromDomainUpdateRecipe(domain recipe.Domain) CreateRecipeResponse{
	return CreateRecipeResponse{
		Message :"Create Recipe Success",
		ID: domain.ID,
		Title: domain.Title,
		DetailRecipe: domain.DetailRecipe,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func FromDomainAllRecipe(domain recipe.Domain) RecipeResponse{
	return RecipeResponse{
		ID: domain.ID,
		Title: domain.Title,
		PatientSessionID: domain.PatientSessionID,
		DetailRecipe: domain.DetailRecipe,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func FromRecipeListDomain(domain []recipe.Domain)[]RecipeResponse{
	var response []RecipeResponse
	for _, value := range domain {
		response = append(response, FromDomainAllRecipe(value))
	}
	return response
}