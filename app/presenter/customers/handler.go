package customers

import (
	"net/http"
	"outlet/v1/app/middleware/auth"
	_request "outlet/v1/app/presenter/customers/request"
	_response "outlet/v1/app/presenter/customers/response"
	"outlet/v1/bussiness/customers"
	"outlet/v1/helpers"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Presenter struct {
	serviceCustomer customers.Service
	jwtAuth         *auth.ConfigJWT
}

func NewHandler(customerService customers.Service, jwtauth *auth.ConfigJWT) *Presenter {
	return &Presenter{
		serviceCustomer: customerService,
		jwtAuth:         jwtauth,
	}
}

func (handler *Presenter) AddCustomer(echoContext echo.Context) error {
	var req _request.Customer
	if err := echoContext.Bind(&req); err != nil {
		response := helpers.APIResponse("Failed Add Customer", http.StatusBadRequest, "Error", nil)
		return echoContext.JSON(http.StatusBadRequest, response)
	}
	domain := _request.ToDomain(req)
	resp, err := handler.serviceCustomer.AddCustomer(domain)
	if err != nil {
		response := helpers.APIResponse("Failed Add Customer", http.StatusInternalServerError, "Error", nil)
		return echoContext.JSON(http.StatusInternalServerError, response)
	}
	response := helpers.APIResponse("Success Add Customer", http.StatusOK, "Success", _response.FromDomain(*resp))
	return echoContext.JSON(http.StatusOK, response)
}

func (handler *Presenter) Update(echoContext echo.Context) error {
	var req _request.Customer
	if err := echoContext.Bind(&req); err != nil {
		response := helpers.APIResponse("Failed FindByEmail", http.StatusBadRequest, "Error", nil)
		return echoContext.JSON(http.StatusBadRequest, response)
	}
	domain := _request.ToDomain(req)
	customer := auth.GetCustomer(echoContext) // ID Get From JWT
	customerID := customer.ID
	resp, err := handler.serviceCustomer.Update(customerID, domain)
	if err != nil {
		response := helpers.APIResponse("Failed", http.StatusBadRequest, "Error", nil)
		return echoContext.JSON(http.StatusBadRequest, response)
	}
	response := helpers.APIResponse("Success", http.StatusOK, "Success", _response.FromDomain(*resp))
	return echoContext.JSON(http.StatusOK, response)
}

func (handler *Presenter) LoginUser(echoContext echo.Context) error {
	var req _request.CustomerLogin
	if err := echoContext.Bind(&req); err != nil {
		response := helpers.APIResponse("Failed Login", http.StatusBadRequest, "Error", nil)
		return echoContext.JSON(http.StatusBadRequest, response)
	}
	resp, err := handler.serviceCustomer.Login(req.Email, req.Password)
	if err != nil {
		response := helpers.APIResponse("Failed Login", http.StatusBadRequest, "Error", nil)
		return echoContext.JSON(http.StatusBadRequest, response)
	}
	if err != nil {
		response := helpers.APIResponse("Generate Token Failed", http.StatusBadRequest, "Error", nil)
		return echoContext.JSON(http.StatusBadRequest, response)
	}
	response := helpers.APIResponse("Success Login", http.StatusOK, "Success", _response.CustomerLogin{Token: resp})
	return echoContext.JSON(http.StatusOK, response)
}

func (handler *Presenter) FindByID(echoContext echo.Context) error {
	id, err := strconv.Atoi(echoContext.Param("id"))
	if err != nil {
		response := helpers.APIResponse("Failed FindByEmail", http.StatusBadRequest, "Error", nil)
		return echoContext.JSON(http.StatusBadRequest, response)
	}
	resp, err := handler.serviceCustomer.FindByID(id)
	if err != nil {
		response := helpers.APIResponse("Failed", http.StatusBadRequest, "Error", nil)
		return echoContext.JSON(http.StatusBadRequest, response)
	}
	response := helpers.APIResponse("Success", http.StatusOK, "Success", _response.FromDomain(*resp))
	return echoContext.JSON(http.StatusOK, response)
}

func (handler *Presenter) Delete(echoContext echo.Context) error {
	var req _request.Customer
	if err := echoContext.Bind(&req); err != nil {
		response := helpers.APIResponse("Failed Delete Customer", http.StatusBadRequest, "Error", err)
		return echoContext.JSON(http.StatusBadRequest, response)
	}
	id, err := strconv.Atoi(echoContext.Param("id"))
	if err != nil {
		response := helpers.APIResponse("Failed Delete Customer", http.StatusBadRequest, "Error", err)
		return echoContext.JSON(http.StatusBadRequest, response)
	}
	domain := _request.ToDomain(req)
	_, err = handler.serviceCustomer.DeleteCustomer(id, domain)
	if err != nil {
		response := helpers.APIResponse("Failed Delete Customer", http.StatusBadRequest, "Error", err)
		return echoContext.JSON(http.StatusBadRequest, response)
	}
	response := helpers.APIResponse("Success", http.StatusOK, "Success", _response.Delete{Data: "Success Delete Customer"})
	return echoContext.JSON(http.StatusOK, response)
}

func (handler *Presenter) GetAllCustomer(echoContext echo.Context) error {
	resp, err := handler.serviceCustomer.GetAllCustomer()
	if err != nil {
		response := helpers.APIResponse("Failed Get All Customer", http.StatusBadRequest, "Error", err)
		return echoContext.JSON(http.StatusBadRequest, response)
	}
	response := helpers.APIResponse("Success", http.StatusOK, "Success", _response.FromDomainArray(*resp))
	return echoContext.JSON(http.StatusOK, response)
}
