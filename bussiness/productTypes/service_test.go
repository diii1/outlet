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
	t.Run("AddProductType | Valid", func(t *testing.T) {

		typesRepository.On("Insert", mock.AnythingOfType("*productTypes.Domain")).Return(&typesDomain, nil).Once()

		result, err := typesService.AddProductType(&typesDomain)

		assert.Nil(t, err)
		assert.Equal(t, &typesDomain, result)
	})

	t.Run("AddProductType | InValid", func(t *testing.T) {
		typesRepository.On("Insert", mock.AnythingOfType("*productTypes.Domain")).Return(&typesDomain, errors.New("error")).Once()
		_, err := typesService.AddProductType(&typesDomain)
		assert.NotNil(t, err)
	})
}
