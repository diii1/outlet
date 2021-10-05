package main

import (
	"log"
	_middleware "outlet/v1/app/middleware"
	"outlet/v1/app/middleware/auth"
	_handlerCustomer "outlet/v1/app/presenter/customers"
	_handlerPaymentMethod "outlet/v1/app/presenter/paymentMethods"
	_handlerProductType "outlet/v1/app/presenter/productTypes"
	_handlerProduct "outlet/v1/app/presenter/products"
	"outlet/v1/app/routes"
	_serviceCustomer "outlet/v1/bussiness/customers"
	_servicePaymentMethod "outlet/v1/bussiness/paymentMethods"
	_serviceProductType "outlet/v1/bussiness/productTypes"
	_serviceProduct "outlet/v1/bussiness/products"
	mysqlRepo "outlet/v1/repository/mysql"
	_repositoryCustomer "outlet/v1/repository/mysql/customers"
	_repositoryPaymentMethod "outlet/v1/repository/mysql/paymentMethods"
	_repositoryProductType "outlet/v1/repository/mysql/productTypes"
	_repositoryProduct "outlet/v1/repository/mysql/products"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("test-config")
	viper.AddConfigPath("./config/")
	viper.AutomaticEnv()
	viper.SetConfigType("json")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func main() {
	configDB := mysqlRepo.ConfigDB{
		DBUsername: viper.GetString(`database.user`),
		DBPassword: viper.GetString(`database.pass`),
		DBHost:     viper.GetString(`database.host`),
		DBPort:     viper.GetString(`database.port`),
		DBDatabase: viper.GetString(`database.name`),
	}

	configJWT := auth.ConfigJWT{
		SecretJWT:       viper.GetString(`jwt.secret`),
		ExpiresDuration: viper.GetInt(`jwt.expired`),
	}

	db := configDB.IntialDB()
	mysqlRepo.MigrateDB(db)

	e := echo.New()

	customerRepository := _repositoryCustomer.NewRepositoryMySQL(db)
	customerService := _serviceCustomer.NewService(customerRepository, &configJWT)
	customerHandler := _handlerCustomer.NewHandler(customerService, &configJWT)

	productTypeRepository := _repositoryProductType.NewRepositoryMySQL(db)
	productTypeService := _serviceProductType.NewService(productTypeRepository)
	productTypeHandler := _handlerProductType.NewHandler(productTypeService)

	productRepository := _repositoryProduct.NewRepositoryMySQL(db)
	productService := _serviceProduct.NewService(productRepository, productTypeService)
	productHandler := _handlerProduct.NewHandler(productService)

	paymentMethodRepository := _repositoryPaymentMethod.NewRepositoryMySQL(db)
	paymentMethodService := _servicePaymentMethod.NewService(paymentMethodRepository)
	paymentMethodHandler := _handlerPaymentMethod.NewHandler(paymentMethodService)

	routesInit := routes.HandlerList{
		JWTMiddleware:        configJWT.Init(),
		HandlerCustomer:      *customerHandler,
		HandlerProductType:   *productTypeHandler,
		HandlerProduct:       *productHandler,
		HandlerPaymentMethod: *paymentMethodHandler,
	}

	routesInit.RouteRegister(e)
	_middleware.LogMiddleware(e)
	log.Fatal(e.Start(viper.GetString("server.address")))
}
