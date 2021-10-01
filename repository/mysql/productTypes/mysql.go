package productTypes

import (
	"github.com/diii1/outlet/bussiness/productTypes"
	"gorm.io/gorm"
)

type repositoryProductTypes struct {
	DB *gorm.DB
}

func NewRepositoryMySQL(db *gorm.DB) productTypes.Repository {
	return &repositoryProductTypes{
		DB: db,
	}
}

func (repository repositoryProductTypes) Insert(productType *productTypes.Domain) (*productTypes.Domain, error) {
	return nil, nil
}
func (repository repositoryProductTypes) Update(id int, productType *productTypes.Domain) (*productTypes.Domain, error) {
	return nil, nil
}
func (repository repositoryProductTypes) FindByID(id int) (*productTypes.Domain, error) {
	return nil, nil
}
