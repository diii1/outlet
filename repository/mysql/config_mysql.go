package mysql

import (
	"fmt"
	"log"

	"github.com/diii1/outlet/repository/mysql/customers"
	"github.com/diii1/outlet/repository/mysql/productTypes"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ConfigDB struct {
	DBUsername string
	DBPassword string
	DBHost     string
	DBPort     string
	DBDatabase string
}

func (config *ConfigDB) IntialDB() *gorm.DB {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		config.DBUsername,
		config.DBPassword,
		config.DBHost,
		config.DBPort,
		config.DBDatabase)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(&productTypes.ProductTypes{})
	db.AutoMigrate(&customers.Customers{})
}
