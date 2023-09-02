package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Ihpaz/golang-restapi-userfamily/entity"
	"github.com/Ihpaz/golang-restapi-userfamily/errors"
	"github.com/Ihpaz/golang-restapi-userfamily/service"
	"github.com/gorilla/mux"
)

type controller struct{}

var (
	customerService service.CustomerService
)

type CustomerController interface {
	GetCustomers(response http.ResponseWriter, request *http.Request)
	AddCustomer(response http.ResponseWriter, request *http.Request)
	GetCustomer(response http.ResponseWriter, request *http.Request)
	UpdateCustomer(response http.ResponseWriter, request *http.Request)
	DeleteCustomer(response http.ResponseWriter, request *http.Request)
}

func NewCustomerController(service service.CustomerService) CustomerController {
	customerService = service
	return &controller{}
}

func (*controller) GetCustomers(response http.ResponseWriter, request *http.Request) {
	// response.Header().Set("Content-Type", "application/json")
	posts, err := customerService.FindAll()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error getting the posts"})
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(posts)
}

func (*controller) AddCustomer(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	var customer entity.Customer

	// fmt.Printf(request)

	err := json.NewDecoder(request.Body).Decode(&customer)
	fmt.Printf("error euy:", err)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error unmarshalling data"})
		return
	}

	err1 := customerService.Validate(&customer)
	if err1 != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: err1.Error()})
		return
	}

	result, err2 := customerService.Create(&customer)
	if err2 != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error saving the customer"})
		return
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(result)
}

func (*controller) GetCustomer(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	var customer *entity.Customer
	vars := mux.Vars(request)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	dtcustomer, err := customerService.FindCustomerByCstId(customer, pid)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error getting the customer"})
	}

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(dtcustomer)
}

func (*controller) UpdateCustomer(response http.ResponseWriter, request *http.Request) {
	// response.Header().Set("Content-Type", "application/json")

	var customer *entity.Customer
	err := json.NewDecoder(request.Body).Decode(&customer)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error unmarshalling data"})
		return
	}

	err1 := customerService.Validate(customer)
	if err1 != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: err1.Error()})
		return
	}

	vars := mux.Vars(request)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)
	result, err := customerService.UpdateACustomer(customer, pid)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error update the customer"})
	}

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(result)
}

func (*controller) DeleteCustomer(response http.ResponseWriter, request *http.Request) {
	// response.Header().Set("Content-Type", "application/json")

	var customer *entity.Customer
	vars := mux.Vars(request)
	pid, err := strconv.ParseUint(vars["id"], 10, 64)

	fmt.Println("uid ini =", pid)
	result, err := customerService.DeleteACustomer(customer, pid)

	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(response).Encode(errors.ServiceError{Message: "Error delete the customer"})
	}

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(result)
}
