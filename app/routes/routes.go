package routes

import (
	"outlet/v1/app/presenter/customers"
	"outlet/v1/app/presenter/productTypes"
	"outlet/v1/app/presenter/products"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type HandlerList struct {
	HandlerCustomer    customers.Presenter
	HandlerProductType productTypes.Presenter
	HandlerProduct     products.Presenter
	JWTMiddleware      middleware.JWTConfig
}

func (handler *HandlerList) RouteRegister(e *echo.Echo) {
	group := e.Group("/api/v1")
	group.POST("/customers/add", handler.HandlerCustomer.AddCustomer)
	group.POST("/customers/login", handler.HandlerCustomer.LoginUser)
	group.PUT("/customers/update", handler.HandlerCustomer.Update, middleware.JWTWithConfig(handler.JWTMiddleware))
	group.GET("/customers/:id", handler.HandlerCustomer.FindByID)
	group.DELETE("/customers/:id", handler.HandlerCustomer.Delete)

	group.POST("/productTypes/add", handler.HandlerProductType.AddProductType)
	group.GET("/productTypes/:id", handler.HandlerProductType.FindByID)
	group.DELETE("/productTypes/:id", handler.HandlerProductType.Delete)

	group.POST("/products/add", handler.HandlerProduct.AddProduct)
	group.GET("/products/:id", handler.HandlerProduct.FindByID)
	group.DELETE("/products/:id", handler.HandlerProduct.Delete)
}
