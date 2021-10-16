package orders_test

import (
	"errors"
	"os"
	"outlet/v1/bussiness/customers"
	_customerRepository "outlet/v1/bussiness/customers/mocks"
	"outlet/v1/bussiness/orders"
	_ordersRepository "outlet/v1/bussiness/orders/mocks"
	"outlet/v1/bussiness/paymentMethods"
	_paymentRepository "outlet/v1/bussiness/paymentMethods/mocks"
	"outlet/v1/bussiness/products"
	_productRepository "outlet/v1/bussiness/products/mocks"
	"outlet/v1/helpers"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	customerService _customerRepository.Service
	customerDomain  customers.Domain
	payService      _paymentRepository.Service
	payDomain       paymentMethods.Domain
	productService  _productRepository.Service
	productDomain   products.Domain
	orderRepository _ordersRepository.Repository
	orderService    orders.Service
	orderDomain     orders.Domain
)

func TestMain(m *testing.M) {
	orderService = orders.NewService(&orderRepository, &productService, &payService, &customerService)
	orderDomain = orders.Domain{
		ID:         1,
		ProductID:  1,
		Product:    "Test Product",
		PaymentID:  1,
		Payment:    "Test Payment",
		CustomerID: 1,
		Customer:   "Test Customer",
		TotalPrice: 15000,
		Status:     "Delivered",
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
	productDomain = products.Domain{
		ID:          1,
		TypeID:      1,
		Name:        "Test Product",
		ProductType: "Test Type",
		Description: "Test desc product",
		Price:       25000,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	payDomain = paymentMethods.Domain{
		ID:        1,
		Name:      "Test Payment Method",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	passHash, _ := helpers.PasswordHash("passTest")
	customerDomain = customers.Domain{
		ID:        2,
		Name:      "Test Customers",
		Email:     "test@example.com",
		Password:  passHash,
		Alamat:    "Test Alamat",
		NoTelepon: "0000000000",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	os.Exit(m.Run())
}

func TestAddOrder(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T) {
		productService.On("FindByID", mock.Anything, mock.Anything).Return(&productDomain, nil).Once()
		payService.On("FindByID", mock.Anything, mock.Anything).Return(&payDomain, nil).Once()
		customerService.On("FindByID", mock.Anything, mock.Anything).Return(&customerDomain, nil).Once()
		orderRepository.On("Insert", mock.Anything, mock.Anything).Return(&orderDomain, nil).Once()
		input := orders.Domain{
			ProductID:  1,
			Product:    "Test Product",
			PaymentID:  1,
			Payment:    "Test Payment",
			CustomerID: 1,
			Customer:   "Test Customer",
			TotalPrice: 15000,
			Status:     "Delivered",
		}
		result, err := orderService.AddOrder(&input)
		assert.Nil(t, err)
		assert.Equal(t, &orderDomain, result)
	})

	t.Run("Invalid Test", func(t *testing.T) {
		productService.On("FindByID", mock.Anything, mock.Anything).Return(&productDomain, errors.New("error")).Once()
		payService.On("FindByID", mock.Anything, mock.Anything).Return(&payDomain, nil).Once()
		customerService.On("FindByID", mock.Anything, mock.Anything).Return(&customerDomain, nil).Once()
		orderRepository.On("Insert", mock.Anything, mock.Anything).Return(&orderDomain, nil).Once()
		input := orders.Domain{
			ProductID:  1,
			Product:    "Test Product",
			PaymentID:  1,
			Payment:    "Test Payment",
			CustomerID: 1,
			Customer:   "Test Customer",
			TotalPrice: 15000,
			Status:     "Delivered",
		}
		_, err := orderService.AddOrder(&input)
		assert.NotNil(t, err)
	})

	t.Run("Invalid Test", func(t *testing.T) {
		productService.On("FindByID", mock.Anything, mock.Anything).Return(&productDomain, errors.New("error")).Once()
		payService.On("FindByID", mock.Anything, mock.Anything).Return(&payDomain, errors.New("error")).Once()
		customerService.On("FindByID", mock.Anything, mock.Anything).Return(&customerDomain, nil).Once()
		orderRepository.On("Insert", mock.Anything, mock.Anything).Return(&orderDomain, nil).Once()
		input := orders.Domain{
			ProductID:  1,
			Product:    "Test Product",
			PaymentID:  1,
			Payment:    "Test Payment",
			CustomerID: 1,
			Customer:   "Test Customer",
			TotalPrice: 15000,
			Status:     "Delivered",
		}
		_, err := orderService.AddOrder(&input)
		assert.NotNil(t, err)
	})

	t.Run("Invalid Test", func(t *testing.T) {
		productService.On("FindByID", mock.Anything, mock.Anything).Return(&productDomain, errors.New("error")).Once()
		payService.On("FindByID", mock.Anything, mock.Anything).Return(&payDomain, errors.New("error")).Once()
		customerService.On("FindByID", mock.Anything, mock.Anything).Return(&customerDomain, errors.New("error")).Once()
		orderRepository.On("Insert", mock.Anything, mock.Anything).Return(&orderDomain, nil).Once()
		input := orders.Domain{
			ProductID:  1,
			Product:    "Test Product",
			PaymentID:  1,
			Payment:    "Test Payment",
			CustomerID: 1,
			Customer:   "Test Customer",
			TotalPrice: 15000,
			Status:     "Delivered",
		}
		_, err := orderService.AddOrder(&input)
		assert.NotNil(t, err)
	})
}

func TestDeleteOrder(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T) {
		productService.On("FindByID", mock.Anything, mock.Anything).Return(&productDomain, nil).Once()
		payService.On("FindByID", mock.Anything, mock.Anything).Return(&payDomain, nil).Once()
		customerService.On("FindByID", mock.Anything, mock.Anything).Return(&customerDomain, nil).Once()
		orderRepository.On("Delete", mock.Anything, mock.Anything).Return(&orderDomain, nil).Once()
		input := orders.Domain{
			ProductID:  1,
			Product:    "Test Product",
			PaymentID:  1,
			Payment:    "Test Payment",
			CustomerID: 1,
			Customer:   "Test Customer",
			TotalPrice: 15000,
			Status:     "Delivered",
		}
		result, err := orderService.DeleteOrder(input.ID, &input)
		assert.Nil(t, err)
		assert.Equal(t, &orderDomain, result)
	})

	t.Run("Invalid Test", func(t *testing.T) {
		productService.On("FindByID", mock.Anything, mock.Anything).Return(&productDomain, nil).Once()
		payService.On("FindByID", mock.Anything, mock.Anything).Return(&payDomain, nil).Once()
		customerService.On("FindByID", mock.Anything, mock.Anything).Return(&customerDomain, nil).Once()
		orderRepository.On("Delete", mock.Anything, mock.Anything).Return(&orderDomain, errors.New("error")).Once()
		input := orders.Domain{
			ProductID:  1,
			Product:    "Test Product",
			PaymentID:  1,
			Payment:    "Test Payment",
			CustomerID: 1,
			Customer:   "Test Customer",
			TotalPrice: 15000,
			Status:     "Delivered",
		}
		_, err := orderService.DeleteOrder(input.ID, &input)
		assert.NotNil(t, err)
	})
}

func TestFindByID(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T) {
		orderRepository.On("FindByID", mock.Anything).Return(&orderDomain, nil).Once()
		result, err := orderService.FindByID(1)
		assert.Nil(t, err)
		assert.Equal(t, &orderDomain, result)
	})
	t.Run("Invalid Test", func(t *testing.T) {
		orderRepository.On("FindByID", mock.Anything).Return(&orderDomain, errors.New("error")).Once()
		_, err := orderService.FindByID(1)
		assert.NotNil(t, err)
	})
}

func TestGetAllOrder(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T) {
		orderRepository.On("GetAllOrder", mock.Anything, mock.Anything).Return(&[]orders.Domain{orderDomain}, nil).Once()
		result, err := orderService.GetAllOrder()
		assert.Nil(t, err)
		assert.Equal(t, &[]orders.Domain{orderDomain}, result)
	})

	t.Run("Invalid Test", func(t *testing.T) {
		orderRepository.On("GetAllOrder", mock.Anything, mock.Anything).Return(&[]orders.Domain{orderDomain}, errors.New("error")).Once()
		_, err := orderService.GetAllOrder()
		assert.NotNil(t, err)
	})
}
