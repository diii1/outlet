package products_test

import (
	"errors"
	"os"
	"outlet/v1/bussiness/productTypes"
	_typeRepository "outlet/v1/bussiness/productTypes/mocks"
	"outlet/v1/bussiness/products"
	_productRepository "outlet/v1/bussiness/products/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	productRepository _productRepository.Repository
	productService    products.Service
	typeService       _typeRepository.Service
	productDomain     products.Domain
	typeDomain        productTypes.Domain
)

func TestMain(m *testing.M) {
	productService = products.NewService(&productRepository, &typeService)
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
	typeDomain = productTypes.Domain{
		ID:        1,
		Name:      "Test Type",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	os.Exit(m.Run())
}

func TestAddProduct(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T) {
		typeService.On("FindByID", mock.Anything, mock.Anything).Return(&typeDomain, nil).Once()
		productRepository.On("Insert", mock.Anything, mock.Anything).Return(&productDomain, nil).Once()
		input := products.Domain{
			TypeID:      1,
			Name:        "Test Product",
			ProductType: "Test Type",
			Description: "Test desc product",
			Price:       25000,
		}
		result, err := productService.AddProduct(&input)
		assert.Nil(t, err)
		assert.Equal(t, &productDomain, result)
	})

	t.Run("Invalid Test", func(t *testing.T) {
		typeService.On("FindByID", mock.Anything, mock.Anything).Return(&typeDomain, errors.New("error")).Once()
		productRepository.On("Insert", mock.Anything, mock.Anything).Return(&productDomain, nil).Once()
		input := products.Domain{
			TypeID:      1,
			Name:        "Test Product",
			ProductType: "Test Type",
			Description: "Test desc product",
			Price:       25000,
		}
		_, err := productService.AddProduct(&input)
		assert.NotNil(t, err)
	})
}

func TestFindByID(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T) {
		productRepository.On("FindByID", mock.Anything).Return(&productDomain, nil).Once()
		result, err := productService.FindByID(1)
		assert.Nil(t, err)
		assert.Equal(t, &productDomain, result)
	})
	t.Run("Invalid Test", func(t *testing.T) {
		productRepository.On("FindByID", mock.Anything).Return(&productDomain, errors.New("error")).Once()
		_, err := productService.FindByID(1)
		assert.NotNil(t, err)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T) {
		typeService.On("FindByID", mock.Anything, mock.Anything).Return(&typeDomain, nil).Once()
		productRepository.On("Update", mock.Anything, mock.Anything).Return(&productDomain, nil).Once()
		input := products.Domain{
			TypeID:      1,
			Name:        "Test Product",
			ProductType: "Test Type",
			Description: "Test desc product",
			Price:       25000,
		}
		result, err := productService.Update(1, &input)
		assert.Nil(t, err)
		assert.Equal(t, &productDomain, result)
	})

	t.Run("Invalid Test", func(t *testing.T) {
		typeService.On("FindByID", mock.Anything, mock.Anything).Return(&typeDomain, errors.New("error")).Once()
		productRepository.On("Update", mock.Anything, mock.Anything).Return(&productDomain, nil).Once()
		input := products.Domain{
			TypeID:      1,
			Name:        "Test Product",
			ProductType: "Test Type",
			Description: "Test desc product",
			Price:       25000,
		}
		_, err := productService.Update(1, &input)
		assert.NotNil(t, err)
	})
}

func TestDeleteProduct(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T) {
		productRepository.On("Delete", mock.Anything, mock.Anything).Return(&productDomain, nil).Once()
		input := products.Domain{
			TypeID:      1,
			Name:        "Test Product",
			ProductType: "Test Type",
			Description: "Test desc product",
			Price:       25000,
		}
		result, err := productService.DeleteProduct(input.ID, &input)
		assert.Nil(t, err)
		assert.Equal(t, &productDomain, result)
	})
	t.Run("Invalid Test", func(t *testing.T) {
		productRepository.On("Delete", mock.Anything, mock.Anything).Return(&productDomain, errors.New("error")).Once()
		input := products.Domain{
			TypeID:      1,
			Name:        "Test Product",
			ProductType: "Test Type",
			Description: "Test desc product",
			Price:       25000,
		}
		_, err := productService.DeleteProduct(input.ID, &input)
		assert.NotNil(t, err)
	})
}

func TestGetAllProduct(t *testing.T) {
	t.Run("Valid Test", func(t *testing.T) {
		productRepository.On("GetAllProduct", mock.Anything, mock.Anything).Return(&[]products.Domain{productDomain}, nil).Once()
		result, err := productService.GetAllProduct()
		assert.Nil(t, err)
		assert.Equal(t, &[]products.Domain{productDomain}, result)
	})

	t.Run("Invalid Test", func(t *testing.T) {
		productRepository.On("GetAllProduct", mock.Anything, mock.Anything).Return(&[]products.Domain{productDomain}, errors.New("error")).Once()
		_, err := productService.GetAllProduct()
		assert.NotNil(t, err)
	})
}
