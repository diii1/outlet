package orders

import (
	"outlet/v1/bussiness/orders"

	"gorm.io/gorm"
)

type repositoryOrders struct {
	DB *gorm.DB
}

func NewRepositoryMySQL(db *gorm.DB) orders.Repository {
	return &repositoryOrders{
		DB: db,
	}
}

func (repository *repositoryOrders) Insert(order *orders.Domain) (*orders.Domain, error) {
	recordOrder := fromDomain(*order)
	if err := repository.DB.Create(&recordOrder).Error; err != nil {
		return &orders.Domain{}, err
	}
	record, err := repository.FindByID(int(recordOrder.ID))
	if err != nil {
		return &orders.Domain{}, err
	}
	return record, nil
}

func (repository *repositoryOrders) FindByID(id int) (*orders.Domain, error) {
	var recordOrder Orders
	if err := repository.DB.Where("orders.id = ?", id).Joins("Products").Joins("PaymentMethods").Joins("Customers").Find(&recordOrder).Error; err != nil {
		return &orders.Domain{}, err
	}
	result := toDomain(recordOrder)
	return &result, nil
}
