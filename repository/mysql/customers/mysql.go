package customers

import (
	"outlet/v1/bussiness/customers"

	"gorm.io/gorm"
)

type repositoryCustomer struct {
	DB *gorm.DB
}

func NewRepositoryMySQL(db *gorm.DB) customers.Repository {
	return &repositoryCustomer{
		DB: db,
	}
}

func (repository repositoryCustomer) Insert(customer *customers.Domain) (*customers.Domain, error) {
	recordCustomer := fromDomain(*customer)
	if err := repository.DB.Create(&recordCustomer).Error; err != nil {
		return &customers.Domain{}, err
	}
	result := toDomain(recordCustomer)
	return &result, nil
}

func (repository repositoryCustomer) Update(id int, customer *customers.Domain) (*customers.Domain, error) {
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

func (repository repositoryCustomer) FindByID(id int) (*customers.Domain, error) {
	recordCustomer := Customers{}
	if err := repository.DB.Where("id = ?", id).First(&recordCustomer).Error; err != nil {
		return &customers.Domain{}, err
	}
	result := toDomain(recordCustomer)
	return &result, nil
}

func (repository repositoryCustomer) FindByEmail(email string) (*customers.Domain, error) {
	recordCustomer := Customers{}
	if err := repository.DB.Where("email = ?", email).First(&recordCustomer).Error; err != nil {
		return &customers.Domain{}, err
	}
	result := toDomain(recordCustomer)
	return &result, nil
}

func (repository repositoryCustomer) Delete(id int, customer *customers.Domain) (*customers.Domain, error) {
	recordCustomer := fromDomain(*customer)
	if err := repository.DB.Where("id = ?", id).Delete(&recordCustomer).Error; err != nil {
		return &customers.Domain{}, err
	}
	result := toDomain(recordCustomer)
	return &result, nil
}

func (repository repositoryCustomer) GetAllCustomer() (*[]customers.Domain, error) {
	var recordCustomer []Customers
	if err := repository.DB.Find(&recordCustomer).Error; err != nil {
		return &[]customers.Domain{}, err
	}
	result := toDomainArray(recordCustomer)
	return &result, nil
}
