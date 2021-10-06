package paymentMethods

import (
	"outlet/v1/bussiness/paymentMethods"

	"gorm.io/gorm"
)

type repositoryPaymentMethod struct {
	DB *gorm.DB
}

func NewRepositoryMySQL(db *gorm.DB) paymentMethods.Repository {
	return &repositoryPaymentMethod{
		DB: db,
	}
}

func (repository repositoryPaymentMethod) Insert(paymentMethod *paymentMethods.Domain) (*paymentMethods.Domain, error) {
	recordPaymentMethod := fromDomain(*paymentMethod)
	if err := repository.DB.Create(&recordPaymentMethod).Error; err != nil {
		return &paymentMethods.Domain{}, err
	}
	result := toDomain(recordPaymentMethod)
	return &result, nil
}

func (repository repositoryPaymentMethod) FindByID(id int) (*paymentMethods.Domain, error) {
	recordPaymentMethod := PaymentMethods{}
	if err := repository.DB.Where("id = ?", id).First(&recordPaymentMethod).Error; err != nil {
		return &paymentMethods.Domain{}, err
	}
	result := toDomain(recordPaymentMethod)
	return &result, nil
}

func (repository repositoryPaymentMethod) Delete(id int, paymentMethod *paymentMethods.Domain) (*paymentMethods.Domain, error) {
	recordPaymentMethod := fromDomain(*paymentMethod)
	if err := repository.DB.Where("id = ?", id).Delete(&recordPaymentMethod).Error; err != nil {
		return &paymentMethods.Domain{}, err
	}
	result := toDomain(recordPaymentMethod)
	return &result, nil
}
