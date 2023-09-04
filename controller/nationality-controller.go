package controller

import (
	"encoding/json"
	"net/http"

	"github.com/Ihpaz/golang-restapi-userfamily/entity"
	"github.com/Ihpaz/golang-restapi-userfamily/errors"
	"github.com/Ihpaz/golang-restapi-userfamily/service"
)

type controllernationality struct{}

var (
	nationalityService service.NationalityService
)

type NationalityController interface {
	GetNationalities(response http.ResponseWriter, request *http.Request)
	AddNationality(response http.ResponseWriter, request *http.Request)
}

func NewNationalityController(service service.NationalityService) NationalityController {
	nationalityService = service
	return &controller{}
}

func (*controller) GetNationalities(response http.ResponseWriter, request *http.Request) {
	nationalities, err := nationalityService.FindAll()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error getting the nationalities"})
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(nationalities)
}

func (*controller) AddNationality(response http.ResponseWriter, request *http.Request) {
	var nationality entity.Nationality
	err := json.NewDecoder(request.Body).Decode(&nationality)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error decode data"})
		return
	}
	err1 := nationalityService.Validate(&nationality)
	if err1 != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: err1.Error()})
		return
	}
	result, err2 := nationalityService.Create(&nationality)
	if err2 != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error saving the nationality"})
		return
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(result)
}
