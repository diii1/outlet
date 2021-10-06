package products

import (
	"net/http"
	_request "outlet/v1/app/presenter/products/request"
	_response "outlet/v1/app/presenter/products/response"
	"outlet/v1/bussiness/products"
	"outlet/v1/helpers"

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
		response := helpers.APIResponse("Failed Add Product", http.StatusBadRequest, "Error", nil)
		return echoContext.JSON(http.StatusBadRequest, response)
	}
	domain := _request.ToDomain(req)
	resp, err := handler.serviceProduct.AddProduct(domain)
	if err != nil {
		response := helpers.APIResponse("Failed Add Product", http.StatusInternalServerError, "Error", nil)
		return echoContext.JSON(http.StatusInternalServerError, response)
	}
	response := helpers.APIResponse("Success Add Product", http.StatusOK, "Success", _response.FromDomain(*resp))
	return echoContext.JSON(http.StatusOK, response)
}
