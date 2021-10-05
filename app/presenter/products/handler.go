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
	serviceProducts products.Service
}

func NewHandler(productsService products.Service) *Presenter {
	return &Presenter{
		serviceProducts: productsService,
	}
}

func (handler *Presenter) AddProduct(echoContext echo.Context) error {
	var req _request.Product
	if err := echoContext.Bind(&req); err != nil {
		response := helpers.APIResponse("Failed Add Product", http.StatusBadRequest, "Error", err)
		return echoContext.JSON(http.StatusBadRequest, response)
	}
	domain := _request.ToDomain(req)
	resp, err := handler.serviceProducts.AddProduct(domain)
	if err != nil {
		response := helpers.APIResponse("Failed Add Product", http.StatusInternalServerError, "Error", nil)
		return echoContext.JSON(http.StatusInternalServerError, response)
	}
	response := helpers.APIResponse("Success Add Product", http.StatusOK, "Success", _response.FromDomain(*resp))
	return echoContext.JSON(http.StatusOK, response)
}

// func (handler *Presenter) FindByID(echoContext echo.Context) error {
// 	id, _ := strconv.Atoi(echoContext.Param("id"))
// 	resp, err := handler.serviceProducts.FindByID(id)
// 	if err != nil {
// 		response := helpers.APIResponse("Failed", http.StatusBadRequest, "Error", nil)
// 		return echoContext.JSON(http.StatusBadRequest, response)
// 	}
// 	response := helpers.APIResponse("Success", http.StatusOK, "Success", _response.FromDomain(*resp))
// 	return echoContext.JSON(http.StatusOK, response)
// }

// func (handler *Presenter) Delete(echoContext echo.Context) error {
// 	var req _request.Product
// 	if err := echoContext.Bind(&req); err != nil {
// 		response := helpers.APIResponse("Failed Delete Product", http.StatusBadRequest, "Error", err)
// 		return echoContext.JSON(http.StatusBadRequest, response)
// 	}
// 	id, err := strconv.Atoi(echoContext.Param("id"))
// 	if err != nil {
// 		response := helpers.APIResponse("Failed Delete Product", http.StatusBadRequest, "Error", err)
// 		return echoContext.JSON(http.StatusBadRequest, response)
// 	}
// 	domain := _request.ToDomain(req)
// 	_, err = handler.serviceProducts.DeleteProduct(id, domain)
// 	if err != nil {
// 		response := helpers.APIResponse("Failed Delete Product", http.StatusBadRequest, "Error", err)
// 		return echoContext.JSON(http.StatusBadRequest, response)
// 	}
// 	response := helpers.APIResponse("Success", http.StatusOK, "Success", _response.Delete{Data: "Success Delete Product"})
// 	return echoContext.JSON(http.StatusBadRequest, response)
// }
