package paymentMethods

import (
	"net/http"
	_request "outlet/v1/app/presenter/paymentMethods/request"
	_response "outlet/v1/app/presenter/paymentMethods/response"
	"outlet/v1/bussiness/paymentMethods"
	"outlet/v1/helpers"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Presenter struct {
	servicePaymentMethod paymentMethods.Service
}

func NewHandler(paymentMethodService paymentMethods.Service) *Presenter {
	return &Presenter{
		servicePaymentMethod: paymentMethodService,
	}
}

func (handler *Presenter) AddPaymentMethod(echoContext echo.Context) error {
	var req _request.PaymentMethod
	if err := echoContext.Bind(&req); err != nil {
		response := helpers.APIResponse("Failed Add Payment Method", http.StatusBadRequest, "Error", nil)
		return echoContext.JSON(http.StatusBadRequest, response)
	}
	domain := _request.ToDomain(req)
	resp, err := handler.servicePaymentMethod.AddPaymentMethod(domain)
	if err != nil {
		response := helpers.APIResponse("Failed Add Payment Method", http.StatusInternalServerError, "Error", nil)
		return echoContext.JSON(http.StatusInternalServerError, response)
	}
	response := helpers.APIResponse("Success Add Payment Method", http.StatusOK, "Success", _response.FromDomain(*resp))
	return echoContext.JSON(http.StatusOK, response)
}

func (handler *Presenter) FindByID(echoContext echo.Context) error {
	id, _ := strconv.Atoi(echoContext.Param("id"))
	resp, err := handler.servicePaymentMethod.FindByID(id)
	if err != nil {
		response := helpers.APIResponse("Failed", http.StatusBadRequest, "Error", nil)
		return echoContext.JSON(http.StatusBadRequest, response)
	}
	response := helpers.APIResponse("Success", http.StatusOK, "Success", _response.FromDomain(*resp))
	return echoContext.JSON(http.StatusOK, response)
}

func (handler *Presenter) Delete(echoContext echo.Context) error {
	var req _request.PaymentMethod
	if err := echoContext.Bind(&req); err != nil {
		response := helpers.APIResponse("Failed Delete Payment Method", http.StatusBadRequest, "Error", err)
		return echoContext.JSON(http.StatusBadRequest, response)
	}
	id, err := strconv.Atoi(echoContext.Param("id"))
	if err != nil {
		response := helpers.APIResponse("Failed Delete Payment Method", http.StatusBadRequest, "Error", err)
		return echoContext.JSON(http.StatusBadRequest, response)
	}
	domain := _request.ToDomain(req)
	_, err = handler.servicePaymentMethod.DeletePaymentMethod(id, domain)
	if err != nil {
		response := helpers.APIResponse("Failed Delete Payment Method", http.StatusBadRequest, "Error", err)
		return echoContext.JSON(http.StatusBadRequest, response)
	}
	response := helpers.APIResponse("Success", http.StatusOK, "Success", _response.Delete{Data: "Success Delete Payment Method"})
	return echoContext.JSON(http.StatusOK, response)
}
