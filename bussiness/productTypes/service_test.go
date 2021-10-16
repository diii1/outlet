package productTypes_test

import (
	"errors"
	"os"
	"outlet/v1/bussiness/productTypes"
	_typesMock "outlet/v1/bussiness/productTypes/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	typesRepository _typesMock.Repository
	typesService    productTypes.Service
	typesDomain     productTypes.Domain
)

func TestMain(m *testing.M) {
	typesService = productTypes.NewService(&typesRepository)
	typesDomain = productTypes.Domain{
		ID:        1,
		Name:      "Test Product Type",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	os.Exit(m.Run())
}

func TestAddProductType(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T) {
		typesRepository.On("Insert", mock.AnythingOfType("*productTypes.Domain")).Return(&typesDomain, nil).Once()
		result, err := typesService.AddProductType(&typesDomain)
		assert.Nil(t, err)
		assert.Equal(t, &typesDomain, result)
	})

	t.Run("Invalid Test", func(t *testing.T) {
		typesRepository.On("Insert", mock.AnythingOfType("*productTypes.Domain")).Return(&typesDomain, errors.New("error")).Once()
		_, err := typesService.AddProductType(&typesDomain)
		assert.NotNil(t, err)
	})
}

func TestFindByID(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T) {
		typesRepository.On("FindByID", mock.Anything).Return(&typesDomain, nil).Once()
		result, err := typesService.FindByID(1)
		assert.Nil(t, err)
		assert.Equal(t, &typesDomain, result)
	})
	t.Run("Invalid Test", func(t *testing.T) {
		typesRepository.On("FindByID", mock.Anything).Return(&typesDomain, errors.New("error")).Once()
		_, err := typesService.FindByID(1)
		assert.NotNil(t, err)
	})
}

func TestDeleteProductType(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T) {
		typesRepository.On("Delete", mock.Anything, mock.Anything).Return(&typesDomain, nil).Once()
		input := productTypes.Domain{
			Name: "type test",
		}
		result, err := typesService.DeleteProductType(input.ID, &input)
		assert.Nil(t, err)
		assert.Equal(t, &typesDomain, result)
	})
	t.Run("Invalid Test", func(t *testing.T) {
		typesRepository.On("Delete", mock.Anything, mock.Anything).Return(&typesDomain, errors.New("error")).Once()
		input := productTypes.Domain{
			Name: "type test",
		}
		_, err := typesService.DeleteProductType(input.ID, &input)
		assert.NotNil(t, err)
	})
}

func TestGetAllProductType(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T) {
		typesRepository.On("GetAllProductType", mock.Anything, mock.Anything).Return(&[]productTypes.Domain{typesDomain}, nil).Once()
		result, err := typesService.GetAllProductType()
		assert.Nil(t, err)
		assert.Equal(t, &[]productTypes.Domain{typesDomain}, result)
	})

	t.Run("Invalid Test", func(t *testing.T) {
		typesRepository.On("GetAllProductType", mock.Anything, mock.Anything).Return(&[]productTypes.Domain{typesDomain}, errors.New("error")).Once()
		_, err := typesService.GetAllProductType()
		assert.NotNil(t, err)
	})
}
