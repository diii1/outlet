package customers

import (
	"errors"
	"outlet/v1/app/middleware/auth"
	"outlet/v1/helpers"
)

type serviceCustomers struct {
	repository Repository
	jwtAuth    *auth.ConfigJWT
}

func NewService(repositoryCustomer Repository, jwtauth *auth.ConfigJWT) Service {
	return &serviceCustomers{
		repository: repositoryCustomer,
		jwtAuth:    jwtauth,
	}
}

func (service *serviceCustomers) AddCustomer(customer *Domain) (*Domain, error) {
	passHash, err := helpers.PasswordHash(customer.Password)
	if err != nil {
		panic(err)
	}
	customer.Password = passHash
	result, err := service.repository.Insert(customer)
	if err != nil {
		return &Domain{}, err
	}
	return result, nil
}

func (service *serviceCustomers) Update(id int, customer *Domain) (*Domain, error) {
	passHash, err := helpers.PasswordHash(customer.Password)
	if err != nil {
		panic(err)
	}
	customer.Password = passHash
	result, err := service.repository.Update(id, customer)
	if err != nil {
		return &Domain{}, err
	}
	return result, nil
}

func (service *serviceCustomers) FindByID(id int) (*Domain, error) {
	customer, err := service.repository.FindByID(id)
	if err != nil {
		return &Domain{}, err
	}
	return customer, nil
}

func (service *serviceCustomers) Login(email string, password string) (string, error) {
	customer, err := service.repository.FindByEmail(email)
	if err != nil {
		return "ID Not Found", errors.New("Customer Not Found")
	}
	if customer.ID == 0 {
		return "ID Not Found", errors.New("Customer Not Found")
	}
	if !helpers.ValidateHash(password, customer.Password) {
		return "Error Validate Hash", errors.New("Error Validate Hash")
	}
	token := service.jwtAuth.GenerateToken(customer.ID)
	return token, nil
}

func (service *serviceCustomers) DeleteCustomer(id int, customer *Domain) (*Domain, error) {
	result, err := service.repository.Delete(id, customer)
	if err != nil {
		return &Domain{}, err
	}
	return result, nil
}

func (service *serviceCustomers) GetAllCustomer() (*[]Domain, error) {
	result, err := service.repository.GetAllCustomer()
	if err != nil {
		return &[]Domain{}, err
	}
	return result, nil
}
