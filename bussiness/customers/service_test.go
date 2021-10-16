package customers_test

import (
	"errors"
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
	customersRepository _customersMock.Repository
	customersService    customers.Service
	customersDomain     customers.Domain
)

func TestMain(m *testing.M) {
	customersService = customers.NewService(&customersRepository, &auth.ConfigJWT{})
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
	t.Run("Valid Test", func(t *testing.T) {
		customersRepository.On("FindByEmail", mock.Anything, mock.Anything).Return(&customersDomain, nil).Once()
		customersRepository.On("Insert", mock.Anything, mock.Anything).Return(&customersDomain, nil).Once()
		input := customers.Domain{
			Name:      "Test Customers",
			Email:     "test@example.com",
			Password:  "passTest",
			Alamat:    "Test Alamat",
			NoTelepon: "0000000000",
		}
		result, err := customersService.AddCustomer(&input)
		assert.Nil(t, err)
		assert.Equal(t, &customersDomain, result)
	})

	t.Run("Invalid Test", func(t *testing.T) {
		customersRepository.On("FindByEmail", mock.Anything, mock.Anything).Return(&customersDomain, errors.New("error")).Once()
		customersRepository.On("Insert", mock.Anything, mock.Anything).Return(&customersDomain, nil).Once()
		input := customers.Domain{
			Name:      "Test Customers",
			Email:     "test@example.com",
			Password:  "passTest",
			Alamat:    "Test Alamat",
			NoTelepon: "0000000000",
		}
		_, err := customersService.AddCustomer(&input)
		assert.Nil(t, err)
	})

	t.Run("Invalid Test", func(t *testing.T) {
		customersRepository.On("FindByEmail", mock.Anything, mock.Anything).Return(&customersDomain, nil).Once()
		customersRepository.On("Insert", mock.Anything, mock.Anything).Return(&customersDomain, errors.New("error")).Once()
		input := customers.Domain{
			Name:      "Test Customers",
			Email:     "test@example.com",
			Password:  "passTest",
			Alamat:    "Test Alamat",
			NoTelepon: "0000000000",
		}
		_, err := customersService.AddCustomer(&input)
		assert.NotNil(t, err)
	})

	t.Run("Invalid Test", func(t *testing.T) {
		customersRepository.On("FindByEmail", mock.Anything, mock.Anything).Return(&customersDomain, errors.New("error")).Once()
		customersRepository.On("Insert", mock.Anything, mock.Anything).Return(&customersDomain, errors.New("error")).Once()
		input := customers.Domain{
			Name:      "Test Customers",
			Email:     "test@example.com",
			Password:  "passTest",
			Alamat:    "Test Alamat",
			NoTelepon: "0000000000",
		}
		_, err := customersService.AddCustomer(&input)
		assert.NotNil(t, err)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T) {
		customersRepository.On("Update", mock.Anything, mock.Anything).Return(&customersDomain, nil).Once()
		input := customers.Domain{
			Name:      "Test Customers",
			Email:     "test@example.com",
			Password:  "passTest",
			Alamat:    "Test Alamat",
			NoTelepon: "0000000000",
		}
		result, err := customersService.Update(1, &input)
		assert.Nil(t, err)
		assert.Equal(t, &customersDomain, result)
	})
	t.Run("Invalid Test", func(t *testing.T) {
		customersRepository.On("Update", mock.Anything, mock.Anything).Return(&customersDomain, errors.New("error")).Once()
		input := customers.Domain{
			Name:      "Test Customers",
			Email:     "test@example.com",
			Password:  "passTest",
			Alamat:    "Test Alamat",
			NoTelepon: "0000000000",
		}
		_, err := customersService.Update(1, &input)
		assert.NotNil(t, err)
	})
}

func TestLogin(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T) {
		customersRepository.On("FindByEmail", mock.Anything, mock.Anything).Return(&customersDomain, nil).Once()
		input := customers.Domain{
			Email:    "test@example.com",
			Password: "passTest",
		}
		result, _ := customersService.Login(input.Email, input.Password)
		assert.NotEmpty(t, result)
	})
	t.Run("Invalid Test", func(t *testing.T) {
		customersRepository.On("FindByEmail", mock.Anything, mock.Anything).Return(&customers.Domain{}, errors.New("error")).Once()
		input := customers.Domain{
			Email:    "test@example.com",
			Password: "passTest",
		}
		result, err := customersService.Login(input.Email, input.Password)
		assert.NotNil(t, err)
		assert.NotEmpty(t, result)
	})
	t.Run("Invalid Test", func(t *testing.T) {
		customersRepository.On("FindByEmail", mock.Anything, mock.Anything).Return(&customersDomain, nil).Once()
		input := customers.Domain{
			Email:    "test@example.com",
			Password: "passTest",
		}
		customersDomain.ID = 0
		result, err := customersService.Login(input.Email, input.Password)
		assert.NotNil(t, err)
		assert.NotEmpty(t, result)
	})
}

func TestFindByID(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T) {
		customersRepository.On("FindByID", mock.Anything).Return(&customersDomain, nil).Once()
		result, err := customersService.FindByID(1)
		assert.Nil(t, err)
		assert.Equal(t, &customersDomain, result)
	})
	t.Run("Invalid Test", func(t *testing.T) {
		customersRepository.On("FindByID", mock.Anything).Return(&customersDomain, errors.New("error")).Once()
		_, err := customersService.FindByID(1)
		assert.NotNil(t, err)
	})
}

func TestDeleteCustomer(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T) {
		customersRepository.On("Delete", mock.Anything, mock.Anything).Return(&customersDomain, nil).Once()
		passHash, _ := helpers.PasswordHash("secretTest")
		input := customers.Domain{
			Name:      "Test Customers",
			Email:     "test@example.com",
			Password:  passHash,
			Alamat:    "Test Alamat",
			NoTelepon: "0000000000",
		}
		result, err := customersService.DeleteCustomer(input.ID, &input)
		assert.Nil(t, err)
		assert.Equal(t, &customersDomain, result)
	})
	t.Run("Invalid Test", func(t *testing.T) {
		customersRepository.On("Delete", mock.Anything, mock.Anything).Return(&customersDomain, errors.New("error")).Once()
		passHash, _ := helpers.PasswordHash("secretTest")
		input := customers.Domain{
			Name:      "Test Customers",
			Email:     "test@example.com",
			Password:  passHash,
			Alamat:    "Test Alamat",
			NoTelepon: "0000000000",
		}
		_, err := customersService.DeleteCustomer(input.ID, &input)
		assert.NotNil(t, err)
	})
}

func TestGetCustomer(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T) {
		customersRepository.On("GetAllCustomer", mock.Anything, mock.Anything).Return(&[]customers.Domain{customersDomain}, nil).Once()
		result, err := customersService.GetAllCustomer()
		assert.Nil(t, err)
		assert.Equal(t, &[]customers.Domain{customersDomain}, result)
	})

	t.Run("Invalid Test", func(t *testing.T) {
		customersRepository.On("GetAllCustomer", mock.Anything, mock.Anything).Return(&[]customers.Domain{customersDomain}, errors.New("error")).Once()
		_, err := customersService.GetAllCustomer()
		assert.NotNil(t, err)
	})
}
