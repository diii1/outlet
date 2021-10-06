package routes

import (
	"outlet/v1/app/presenter/customers"
	"outlet/v1/app/presenter/orders"
	"outlet/v1/app/presenter/paymentMethods"
	"outlet/v1/app/presenter/productTypes"
	"outlet/v1/app/presenter/products"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type HandlerList struct {
	HandlerCustomer      customers.Presenter
	HandlerProductType   productTypes.Presenter
	HandlerPaymentMethod paymentMethods.Presenter
	HandlerProduct       products.Presenter
	HandlerOrder         orders.Presenter
	JWTMiddleware        middleware.JWTConfig
}

func (handler *HandlerList) RouteRegister(e *echo.Echo) {
	group := e.Group("/api/v1")

	// Route for Customer Endpoint
	group.GET("/customers", handler.HandlerCustomer.GetAllCustomer)
	group.POST("/customers/add", handler.HandlerCustomer.AddCustomer)
	group.POST("/customers/login", handler.HandlerCustomer.LoginUser)
	group.PUT("/customers/update", handler.HandlerCustomer.Update, middleware.JWTWithConfig(handler.JWTMiddleware))
	group.GET("/customers/:id", handler.HandlerCustomer.FindByID)
	group.DELETE("/customers/:id", handler.HandlerCustomer.Delete)

	// Route for Product Type Endpoint
	group.GET("/productTypes", handler.HandlerProductType.GetAllProductType)
	group.POST("/productTypes/add", handler.HandlerProductType.AddProductType)
	group.GET("/productTypes/:id", handler.HandlerProductType.FindByID)
	group.DELETE("/productTypes/:id", handler.HandlerProductType.Delete)

	// Route for Product Endpoint
	group.GET("/products", handler.HandlerProduct.GetAllProduct)
	group.POST("/products/add", handler.HandlerProduct.AddProduct)
	group.PUT("/products/update/:id", handler.HandlerProduct.Update)
	group.GET("/products/:id", handler.HandlerProduct.FindByID)
	group.DELETE("/products/:id", handler.HandlerProduct.DeleteProduct)

	// Route for Payment Method Endpoint
	group.GET("/paymentMethods", handler.HandlerPaymentMethod.GetAllPaymentMethod)
	group.POST("/paymentMethods/add", handler.HandlerPaymentMethod.AddPaymentMethod)
	group.GET("/paymentMethods/:id", handler.HandlerPaymentMethod.FindByID)
	group.DELETE("/paymentMethods/:id", handler.HandlerPaymentMethod.DeletePaymentMethod)

	// Route for Order Endpoint
	group.GET("/orders", handler.HandlerOrder.GetAllOrder)
	group.POST("/orders/add", handler.HandlerOrder.AddOrder)
	group.GET("/orders/:id", handler.HandlerOrder.FindByID)
	group.DELETE("/orders/:id", handler.HandlerOrder.DeleteOrder)
}
