package response

import (
	"finalproject/features/patientses"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CreatePatientsesResponse struct {
	ID      int    `json:"id"`
	Status  string `json:"status"`
	Date    string `json:"date"`
	Message string `json:"message"`
}
type BaseResponse struct {
	Meta struct {
		Status   int      `json:"rc"`
		Message  string   `json:"message"`
		Messages []string `json:"messages,omitempty"`
	} `json:"meta"`
	Data interface{} `json:"data"`
}

func NewSuccessResponse(c echo.Context, param interface{}) error{
	response := BaseResponse{}
	response.Meta.Status = http.StatusOK
	response.Meta.Message = "Success"
	response.Data = param
	return c.JSON(http.StatusOK, response)
}

func NewErrorResponse(c echo.Context, status int, err error) error {
	response := BaseResponse{}
	response.Meta.Status = status
	response.Meta.Message = "Something Error"
	response.Meta.Messages = []string{err.Error()}
	return c.JSON(status, response)
}

func FromDomainCreate(domain patientses.Domain) CreatePatientsesResponse{
	return CreatePatientsesResponse{
		ID: domain.ID,
		Status: domain.Status,
		Date: domain.Date,
		Message: "Create Patient Session Success",
	}
}
type PatientsesResponse struct{
	ID int `json:"id"`
	Status string `json:"status"`
	Date string `json:"date"`

}
func FromDomainAllPatientses(domain patientses.Domain) PatientsesResponse{
	return PatientsesResponse{
		ID: domain.ID,
		Status: domain.Status,
		Date: domain.Date,
	}
}

func FromDomainUpdatePatientses(domain patientses.Domain) CreatePatientsesResponse{
	return CreatePatientsesResponse{
		Message : "Update Patient Session Success",
		ID: domain.ID,
		Status: domain.Status,
		Date: domain.Date,
	}
}
func FromPatientsesListDomain(domain []patientses.Domain) []PatientsesResponse{
	var response []PatientsesResponse
	for _, value := range domain {
		response = append(response, FromDomainAllPatientses(value))
	}
	return response
}