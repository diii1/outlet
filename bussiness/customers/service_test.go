package customers_test

import (
	"os"
	"outlet/v1/app/middleware/auth"
	"outlet/v1/bussiness/customers"
	_customersMock "outlet/v1/bussiness/customers/mocks"
	"outlet/v1/helpers"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	customersMockRepository _customersMock.Repository
	customersMockService    _customersMock.Service
	customersService        customers.Service
	customersDomain         customers.Domain
)

func TestMain(m *testing.M) {
	customersService = customers.NewService(&customersMockRepository, &auth.ConfigJWT{})
	passHash, _ := helpers.PasswordHash("secretTest")
	customersDomain = customers.Domain{
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

func TestAddCustomer(t *testing.T) {
	t.Run("AddCustomer | Valid", func(t *testing.T) {
		customersMockRepository.On("FindByEmail", mock.Anything, mock.Anything).Return(&customersDomain, nil).Once()
		customersMockRepository.On("Insert", mock.AnythingOfType("*customers.Domain")).Return(&customersDomain, nil).Once()
		input := customers.Domain{
			ID:        2,
			Name:      "Test Customers",
			Email:     "test@example.com",
			Password:  "secretTest",
			Alamat:    "Test Alamat",
			NoTelepon: "0000000000",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		_, err := customersMockRepository.Insert(&input)
		assert.Nil(t, err)
	})
}
