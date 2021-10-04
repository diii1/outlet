package productTypes

import (
	"net/http"
	_request "outlet/v1/app/presenter/productTypes/request"
	_response "outlet/v1/app/presenter/productTypes/response"
	"outlet/v1/bussiness/productTypes"
	"outlet/v1/helpers"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Presenter struct {
	serviceProductType productTypes.Service
}

func NewHandler(productTypeService productTypes.Service) *Presenter {
	return &Presenter{
		serviceProductType: productTypeService,
	}
}

func (handler *Presenter) AddProductType(echoContext echo.Context) error {
	var req _request.ProductType
	if err := echoContext.Bind(&req); err != nil {
		response := helpers.APIResponse("Failed Add Product Type", http.StatusBadRequest, "Error", nil)
		return echoContext.JSON(http.StatusBadRequest, response)
	}
	domain := _request.ToDomain(req)
	resp, err := handler.serviceProductType.AddProductType(domain)
	if err != nil {
		response := helpers.APIResponse("Failed Add Product Type", http.StatusInternalServerError, "Error", nil)
		return echoContext.JSON(http.StatusInternalServerError, response)
	}
	response := helpers.APIResponse("Success Add Product Type", http.StatusOK, "Success", _response.FromDomain(*resp))
	return echoContext.JSON(http.StatusOK, response)
}

func (handler *Presenter) FindByID(echoContext echo.Context) error {
	id, _ := strconv.Atoi(echoContext.Param("id"))
	resp, err := handler.serviceProductType.FindByID(id)
	if err != nil {
		response := helpers.APIResponse("Failed", http.StatusBadRequest, "Error", nil)
		return echoContext.JSON(http.StatusBadRequest, response)
	}
	response := helpers.APIResponse("Success", http.StatusOK, "Success", _response.FromDomain(*resp))
	return echoContext.JSON(http.StatusOK, response)
}

func (handler *Presenter) Delete(echoContext echo.Context) error {
	var req _request.ProductType
	if err := echoContext.Bind(&req); err != nil {
		response := helpers.APIResponse("Failed Delete Product Type", http.StatusBadRequest, "Error", err)
		return echoContext.JSON(http.StatusBadRequest, response)
	}
	id, err := strconv.Atoi(echoContext.Param("id"))
	if err != nil {
		response := helpers.APIResponse("Failed Delete Product Type", http.StatusBadRequest, "Error", err)
		return echoContext.JSON(http.StatusBadRequest, response)
	}
	domain := _request.ToDomain(req)
	_, err = handler.serviceProductType.DeleteProductType(id, domain)
	if err != nil {
		response := helpers.APIResponse("Failed Delete Product Type", http.StatusBadRequest, "Error", err)
		return echoContext.JSON(http.StatusBadRequest, response)
	}
	response := helpers.APIResponse("Success", http.StatusOK, "Success", _response.Delete{Data: "Success Delete Product Type"})
	return echoContext.JSON(http.StatusBadRequest, response)
}
