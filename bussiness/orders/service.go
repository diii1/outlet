package orders

import (
	"outlet/v1/bussiness/customers"
	"outlet/v1/bussiness/paymentMethods"
	"outlet/v1/bussiness/products"
)

type serviceOrder struct {
	repository     Repository
	productDomain  products.Service
	paymentDomain  paymentMethods.Service
	customerDomain customers.Service
}

func NewService(repositoryOrder Repository, productServ products.Service, paymentServ paymentMethods.Service, customerServ customers.Service) Service {
	return &serviceOrder{
		repository:     repositoryOrder,
		productDomain:  productServ,
		paymentDomain:  paymentServ,
		customerDomain: customerServ,
	}
}

func (service *serviceOrder) AddOrder(order *Domain) (*Domain, error) {
	product, err := service.productDomain.FindByID(order.ProductID)
	if err != nil {
		return &Domain{}, err
	}
	order.Product = product.Name
	order.TotalPrice = product.Price
	payment, err := service.paymentDomain.FindByID(order.PaymentID)
	if err != nil {
		return &Domain{}, err
	}
	order.Payment = payment.Name
	customer, err := service.customerDomain.FindByID(order.CustomerID)
	if err != nil {
		return &Domain{}, err
	}
	order.Customer = customer.Name
	order.Status = "Devlivered"
	result, err := service.repository.Insert(order)
	if err != nil {
		return &Domain{}, err
	}
	return result, nil
}

func (service *serviceOrder) FindByID(id int) (*Domain, error) {
	order, err := service.repository.FindByID(id)
	if err != nil {
		return &Domain{}, err
	}
	return order, nil
}
func (service *serviceOrder) DeleteOrder(id int, order *Domain) (*Domain, error) {
	result, err := service.repository.Delete(id, order)
	if err != nil {
		return &Domain{}, err
	}
	return result, nil
}

func (service *serviceOrder) GetAllOrder() (*[]Domain, error) {
	result, err := service.repository.GetAllOrder()
	if err != nil {
		return &[]Domain{}, err
	}
	return result, nil
}
