package main

import (
	"log"
	_handlerCustomer "outlet/v1/app/presenter/customers"
	_handlerProductType "outlet/v1/app/presenter/productTypes"
	"outlet/v1/app/presenter/routes"
	_serviceCustomer "outlet/v1/bussiness/customers"
	_serviceProductType "outlet/v1/bussiness/productTypes"
	mysqlRepo "outlet/v1/repository/mysql"
	_repoCustomer "outlet/v1/repository/mysql/customers"
	_repoProductType "outlet/v1/repository/mysql/productTypes"

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

	db := configDB.IntialDB()
	mysqlRepo.MigrateDB(db)

	e := echo.New()

	productTypeRepository := _repoProductType.NewRepositoryMySQL(db)
	productTypeService := _serviceProductType.NewService(productTypeRepository)
	productTypeHandler := _handlerProductType.NewHandler(productTypeService)

	customerRepository := _repoCustomer.NewRepositoryMySQL(db)
	customerService := _serviceCustomer.NewService(customerRepository)
	customerHandler := _handlerCustomer.NewHandler(customerService)

	routesInit := routes.HandlerList{
		HandlerProductType: *productTypeHandler,
		HandlerCustomer:    *customerHandler,
	}
	routesInit.RouteRegister(e)
	log.Fatal(e.Start(viper.GetString("server.address")))
}
