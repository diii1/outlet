package customers

import (
	"net/http"
	"strconv"

	_request "github.com/diii1/outlet/app/presenter/customers/request"
	_response "github.com/diii1/outlet/app/presenter/customers/response"
	"github.com/diii1/outlet/bussiness/customers"
	"github.com/diii1/outlet/helpers"
	"github.com/labstack/echo/v4"
)

type Presenter struct {
	serviceCustomer customers.Service
}

func NewHandler(customerService customers.Service) *Presenter {
	return &Presenter{
		serviceCustomer: customerService,
	}
}

func (handler *Presenter) AddCustomer(echoContext echo.Context) error {
	var req _request.Customer
	if err := echoContext.Bind(&req); err != nil {
		response := helpers.APIResponse("Failed add Customer", http.StatusBadRequest, "Error", nil)
		return echoContext.JSON(http.StatusBadRequest, response)
	}
	domain := _request.ToDomain(req)
	resp, err := handler.serviceCustomer.AddCustomer(domain)
	if err != nil {
		response := helpers.APIResponse("Failed add Customer", http.StatusInternalServerError, "Error", nil)
		return echoContext.JSON(http.StatusInternalServerError, response)
	}
	response := helpers.APIResponse("Success add Customer", http.StatusOK, "Success", _response.FromDomain(*resp))
	return echoContext.JSON(http.StatusOK, response)
}
func (handler *Presenter) Update(echoContext echo.Context) error {
	var req _request.Customer
	domain := _request.ToDomain(req)
	id, _ := strconv.Atoi(echoContext.Param("id"))
	resp, err := handler.serviceCustomer.Update(id, domain)
	if err != nil {
		response := helpers.APIResponse("Failed update Customer", http.StatusBadRequest, "Error", nil)
		return echoContext.JSON(http.StatusBadRequest, response)
	}
	response := helpers.APIResponse("Success update Customer", http.StatusOK, "Success", _response.FromDomain(*resp))
	return echoContext.JSON(http.StatusOK, response)
}
func (handler *Presenter) FindByID(echoContext echo.Context) error {
	id, _ := strconv.Atoi(echoContext.Param("id"))
	resp, err := handler.serviceCustomer.FindByID(id)
	if err != nil {
		response := helpers.APIResponse("Failed find Customer", http.StatusBadRequest, "Error", nil)
		return echoContext.JSON(http.StatusBadRequest, response)
	}
	response := helpers.APIResponse("Success", http.StatusOK, "Success", _response.FromDomain(*resp))
	return echoContext.JSON(http.StatusOK, response)
}
