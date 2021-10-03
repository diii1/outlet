package customers

import (
	"outlet/v1/bussiness/customers"

	"gorm.io/gorm"
)

type repositoryCustomers struct {
	DB *gorm.DB
}

func NewRepositoryMySQL(db *gorm.DB) customers.Repository {
	return &repositoryCustomers{
		DB: db,
	}
}

func (repository repositoryCustomers) Insert(customer *customers.Domain) (*customers.Domain, error) {
	recordCustomer := fromDomain(*customer)
	if err := repository.DB.Create(&recordCustomer).Error; err != nil {
		return &customers.Domain{}, err
	}
	result := toDomain(recordCustomer)
	return &result, nil
}
func (repository repositoryCustomers) Update(id int, customer *customers.Domain) (*customers.Domain, error) {
	recordCustomer := fromDomain(*customer)
	if err := repository.DB.Where("id = ?", id).Updates(&recordCustomer).Error; err != nil {
		return &customers.Domain{}, err
	}
	if err := repository.DB.Where("id = ?", id).First(&recordCustomer).Error; err != nil {
		return &customers.Domain{}, err
	}
	result := toDomain(recordCustomer)
	return &result, nil
}
func (repository repositoryCustomers) FindByID(id int) (*customers.Domain, error) {
	var recordCustomer Customers
	if err := repository.DB.Where("id = ?", id).First(&recordCustomer).Error; err != nil {
		return &customers.Domain{}, err
	}
	result := toDomain(recordCustomer)
	return &result, nil
}
