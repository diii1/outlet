package routes

import (
	"github.com/diii1/outlet/app/presenter/customers"
	"github.com/diii1/outlet/app/presenter/productTypes"
	"github.com/labstack/echo/v4"
)

type HandlerList struct {
	HandlerProductType productTypes.Presenter
	HandlerCustomer    customers.Presenter
}

func (handler *HandlerList) RouteRegister(e *echo.Echo) {
	group := e.Group("/api")

	group.POST("/productType/add", handler.HandlerProductType.AddProductType)
	group.POST("/productType/update", handler.HandlerProductType.Update)
	group.GET("/productType/:id", handler.HandlerProductType.FindByID)

	group.POST("/customer/add", handler.HandlerCustomer.AddCustomer)
	group.POST("/customer/update", handler.HandlerCustomer.Update)
	group.GET("/customer/:id", handler.HandlerCustomer.FindByID)
}
