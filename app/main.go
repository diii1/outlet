package main

import (
	"log"

	_handlerCustomer "github.com/diii1/outlet/app/presenter/customers"
	_handlerProductType "github.com/diii1/outlet/app/presenter/productTypes"
	"github.com/diii1/outlet/app/routes"
	_serviceCustomer "github.com/diii1/outlet/bussiness/customers"
	_serviceProductType "github.com/diii1/outlet/bussiness/productTypes"
	repoMysql "github.com/diii1/outlet/repository/mysql"
	_repoCustomer "github.com/diii1/outlet/repository/mysql/customers"
	_repoProductType "github.com/diii1/outlet/repository/mysql/productTypes"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func Init() {
	viper.SetConfigName("test-config")
	viper.AddConfigPath("./app/config/")
	viper.AutomaticEnv()
	viper.SetConfigType("json")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUNNING on DEBUG mode")
	}
}

func main() {
	configDB := repoMysql.ConfigDB{
		DBUsername: viper.GetString(`database.user`),
		DBPassword: viper.GetString(`database.pass`),
		DBHost:     viper.GetString(`database.host`),
		DBPort:     viper.GetString(`database.port`),
		DBDatabase: viper.GetString(`database.name`),
	}

	db := configDB.IntialDB()
	repoMysql.MigrateDB(db)

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
