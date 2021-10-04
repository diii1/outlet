package routes

import (
	"outlet/v1/app/presenter/customers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type HandlerList struct {
	HandlerCustomer customers.Presenter
	JWTMiddleware   middleware.JWTConfig
}

func (handler *HandlerList) RouteRegister(e *echo.Echo) {
	group := e.Group("/api/v1")
	group.POST("/customers/add", handler.HandlerCustomer.AddCustomer)
	group.POST("/customers/login", handler.HandlerCustomer.LoginUser)
	group.PUT("/customers/update", handler.HandlerCustomer.Update, middleware.JWTWithConfig(handler.JWTMiddleware))
	group.GET("/customers/:id", handler.HandlerCustomer.FindByID)
	group.DELETE("/customers/:id", handler.HandlerCustomer.Delete)
}
