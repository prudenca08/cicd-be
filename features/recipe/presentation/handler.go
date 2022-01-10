package presentation

import (
	"finalproject/features/recipe"
	"finalproject/features/recipe/presentation/request"
	"finalproject/features/recipe/presentation/response"
	"finalproject/middleware"
	"strconv"

	"net/http"

	"github.com/labstack/echo/v4"
)

type RecipeHandler struct {
	recipeService recipe.Service
}
func NewHandlerRecipe(serv recipe.Service) *RecipeHandler{
	return &RecipeHandler{
		recipeService: serv,
	}
}

func (ctrl *RecipeHandler) Create(c echo.Context) error {
	createReq := request.Recipe{}
	if err := c.Bind(&createReq); err != nil {
		return response.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	jwtGetID := middleware.GetUser(c)
	result, err := ctrl.recipeService.Create(jwtGetID.ID, createReq.ToDomain())

	if err != nil {
		return response.NewErrorResponse(c, http.StatusBadGateway, err)
	}
	return response.NewSuccessResponse(c, response.FromDomainCreate(result))
}

func (ctrl *RecipeHandler) Update(c echo.Context) error {
	updateReq := request.Recipe{}
	if err := c.Bind(&updateReq); err !=nil {
		return response.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	id,_ := strconv.Atoi(c. Param("id"))
	jwtGetID := middleware.GetUser(c)
	result, err := ctrl.recipeService.Update(jwtGetID.ID, id, updateReq.ToDomain())
	if err != nil {
		return response.NewErrorResponse(c, http.StatusBadGateway, err)
	}
	return response.NewSuccessResponse(c, response.FromDomainUpdateRecipe(result))
}

func (ctrl *RecipeHandler) Delete(c echo.Context) error {
	recID := middleware.GetUser(c)
	deletedId, _ := strconv.Atoi(c.Param("id"))
	result, err := ctrl.recipeService.Delete(recID.ID, deletedId)
	if err != nil {
		return response.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return response.NewSuccessResponse(c, result)
}
func (ctrl *RecipeHandler) AllRecipe(c echo.Context) error{
	result, err := ctrl.recipeService.AllRecipe()

	if err != nil {
		return response.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	return response.NewSuccessResponse(c, response.FromRecipeListDomain(result))
}

func (ctrl *RecipeHandler) RecipeByID(c echo.Context) error {
	recipeID, _ := strconv.Atoi(c.Param("id"))
	result, err := ctrl.recipeService.RecipeByID(recipeID)
	if err != nil {
		return response.NewErrorResponse(c, http.StatusBadGateway, err)
	}
	return response.NewSuccessResponse(c, response.FromDomainAllRecipe(result))
}