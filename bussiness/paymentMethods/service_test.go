package paymentMethods_test

import (
	"errors"
	"os"
	"outlet/v1/bussiness/paymentMethods"
	_payRepository "outlet/v1/bussiness/paymentMethods/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	payRepository _payRepository.Repository
	payService    paymentMethods.Service
	payDomain     paymentMethods.Domain
)

func TestMain(m *testing.M) {
	payService = paymentMethods.NewService(&payRepository)
	payDomain = paymentMethods.Domain{
		ID:        1,
		Name:      "Test Payment Method",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	os.Exit(m.Run())
}

func TestAddPaymentMethod(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T) {
		payRepository.On("Insert", mock.Anything, mock.Anything).Return(&payDomain, nil).Once()
		input := paymentMethods.Domain{
			Name: "payment test",
		}
		result, err := payService.AddPaymentMethod(&input)
		assert.Nil(t, err)
		assert.Equal(t, &payDomain, result)
	})
	t.Run("Invalid Test", func(t *testing.T) {
		payRepository.On("Insert", mock.Anything, mock.Anything).Return(&payDomain, errors.New("error")).Once()
		input := paymentMethods.Domain{
			Name: "payment test",
		}
		_, err := payService.AddPaymentMethod(&input)
		assert.NotNil(t, err)
	})
}

func TestFindByID(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T) {
		payRepository.On("FindByID", mock.Anything).Return(&payDomain, nil).Once()
		result, err := payService.FindByID(1)
		assert.Nil(t, err)
		assert.Equal(t, &payDomain, result)
	})
	t.Run("Invalid Test", func(t *testing.T) {
		payRepository.On("FindByID", mock.Anything).Return(&payDomain, errors.New("error")).Once()
		_, err := payService.FindByID(1)
		assert.NotNil(t, err)
	})
}

func TestDeletePaymentMethod(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T) {
		payRepository.On("Delete", mock.Anything, mock.Anything).Return(&payDomain, nil).Once()
		input := paymentMethods.Domain{
			Name: "payment test",
		}
		result, err := payService.DeletePaymentMethod(input.ID, &input)
		assert.Nil(t, err)
		assert.Equal(t, &payDomain, result)
	})
	t.Run("Invalid Test", func(t *testing.T) {
		payRepository.On("Delete", mock.Anything, mock.Anything).Return(&payDomain, errors.New("error")).Once()
		input := paymentMethods.Domain{
			Name: "payment test",
		}
		_, err := payService.DeletePaymentMethod(input.ID, &input)
		assert.NotNil(t, err)
	})
}

func TestGetAllProductType(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T) {
		payRepository.On("GetAllPaymentMethod", mock.Anything, mock.Anything).Return(&[]paymentMethods.Domain{payDomain}, nil).Once()
		result, err := payService.GetAllPaymentMethod()
		assert.Nil(t, err)
		assert.Equal(t, &[]paymentMethods.Domain{payDomain}, result)
	})

	t.Run("Invalid Test", func(t *testing.T) {
		payRepository.On("GetAllPaymentMethod", mock.Anything, mock.Anything).Return(&[]paymentMethods.Domain{payDomain}, errors.New("error")).Once()
		_, err := payService.GetAllPaymentMethod()
		assert.NotNil(t, err)
	})
}
