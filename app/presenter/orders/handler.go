package orders

import (
	"net/http"
	_request "outlet/v1/app/presenter/orders/request"
	_response "outlet/v1/app/presenter/orders/response"
	"outlet/v1/bussiness/orders"
	"outlet/v1/helpers"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Presenter struct {
	serviceOrder orders.Service
}

func NewHandler(orderService orders.Service) *Presenter {
	return &Presenter{
		serviceOrder: orderService,
	}
}

func (handler *Presenter) AddOrder(echoContext echo.Context) error {
	var req _request.Orders
	if err := echoContext.Bind(&req); err != nil {
		response := helpers.APIResponse("Failed Add Order", http.StatusBadRequest, "Error", nil)
		return echoContext.JSON(http.StatusBadRequest, response)
	}
	domain := _request.ToDomain(req)
	resp, err := handler.serviceOrder.AddOrder(domain)
	if err != nil {
		response := helpers.APIResponse("Failed Add Order", http.StatusInternalServerError, "Error", nil)
		return echoContext.JSON(http.StatusInternalServerError, response)
	}
	response := helpers.APIResponse("Success Add Order", http.StatusOK, "Success", _response.FromDomain(*resp))
	return echoContext.JSON(http.StatusOK, response)
}

func (handler *Presenter) FindByID(echoContext echo.Context) error {
	id, _ := strconv.Atoi(echoContext.Param("id"))
	resp, err := handler.serviceOrder.FindByID(id)
	if err != nil {
		response := helpers.APIResponse("Failed", http.StatusBadRequest, "Error", nil)
		return echoContext.JSON(http.StatusBadRequest, response)
	}
	response := helpers.APIResponse("Success", http.StatusOK, "Success", _response.FromDomain(*resp))
	return echoContext.JSON(http.StatusOK, response)
}
