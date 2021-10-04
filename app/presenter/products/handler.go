package products

import (
	"net/http"
	"outlet/v1/bussiness/products"
	"outlet/v1/helpers"
	"strconv"

	_request "outlet/v1/app/presenter/products/request"
	_response "outlet/v1/app/presenter/products/response"

	"github.com/labstack/echo/v4"
)

type Presenter struct {
	serviceProduct products.Service
}

func NewHandler(productService products.Service) *Presenter {
	return &Presenter{
		serviceProduct: productService,
	}
}

func (handler *Presenter) AddProduct(echoContext echo.Context) error {
	var req _request.Products
	if err := echoContext.Bind(&req); err != nil {
		response := helpers.APIResponse("Failed add Product", http.StatusBadRequest, "Error", nil)
		return echoContext.JSON(http.StatusBadRequest, response)
	}
	domain := _request.ToDomain(req)
	resp, err := handler.serviceProduct.AddProduct(domain)
	if err != nil {
		response := helpers.APIResponse("Failed add Product", http.StatusInternalServerError, "Error", nil)
		return echoContext.JSON(http.StatusInternalServerError, response)
	}
	response := helpers.APIResponse("Success add Product", http.StatusOK, "Success", _response.FromDomain(*resp))
	return echoContext.JSON(http.StatusOK, response)
}
func (handler *Presenter) Update(echoContext echo.Context) error {
	var req _request.Products
	domain := _request.ToDomain(req)
	id, _ := strconv.Atoi(echoContext.Param("id"))
	resp, err := handler.serviceProduct.Update(domain, id)
	if err != nil {
		response := helpers.APIResponse("Failed update Product", http.StatusBadRequest, "Error", nil)
		return echoContext.JSON(http.StatusBadRequest, response)
	}
	response := helpers.APIResponse("Success update Product", http.StatusOK, "Success", _response.FromDomain(*resp))
	return echoContext.JSON(http.StatusOK, response)
}
func (handler *Presenter) FindByID(echoContext echo.Context) error {
	id, _ := strconv.Atoi(echoContext.Param("id"))
	resp, err := handler.serviceProduct.FindByID(id)
	if err != nil {
		response := helpers.APIResponse("Failed find Product", http.StatusBadRequest, "Error", nil)
		return echoContext.JSON(http.StatusBadRequest, response)
	}
	response := helpers.APIResponse("Success", http.StatusOK, "Success", _response.FromDomain(*resp))
	return echoContext.JSON(http.StatusOK, response)
}
