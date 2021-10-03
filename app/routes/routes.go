package routes

import (
	"outlet/v1/app/presenter/customers"
	"outlet/v1/app/presenter/productTypes"
	"outlet/v1/app/presenter/products"

	"github.com/labstack/echo/v4"
)

type HandlerList struct {
	HandlerProductType productTypes.Presenter
	HandlerCustomer    customers.Presenter
	HandlerProduct     products.Presenter
}

func (handler *HandlerList) RouteRegister(e *echo.Echo) {
	group := e.Group("/api")

	group.POST("/productType/add", handler.HandlerProductType.AddProductType)
	group.POST("/productType/update", handler.HandlerProductType.Update)
	group.GET("/productType/:id", handler.HandlerProductType.FindByID)

	group.POST("/customer/add", handler.HandlerCustomer.AddCustomer)
	group.POST("/customer/update", handler.HandlerCustomer.Update)
	group.GET("/customer/:id", handler.HandlerCustomer.FindByID)

	group.POST("/product/add", handler.HandlerProduct.AddProduct)
	group.POST("/product/update", handler.HandlerProduct.Update)
	group.GET("/product/:id", handler.HandlerProduct.FindByID)
}
