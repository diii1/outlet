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
	recordProductType := fromDomain(*productType)
	if err := repository.DB.Create(&recordProductType).Error; err != nil {
		return &productTypes.Domain{}, err
	}
	result := toDomain(recordProductType)
	return &result, nil
}
func (repository repositoryProductTypes) Update(id int, productType *productTypes.Domain) (*productTypes.Domain, error) {
	recordProductType := fromDomain(*productType)
	if err := repository.DB.Where("id = ?", id).Updates(&recordProductType).Error; err != nil {
		return &productTypes.Domain{}, err
	}
	if err := repository.DB.Where("id = ?", id).First(&recordProductType).Error; err != nil {
		return &productTypes.Domain{}, err
	}
	result := toDomain(recordProductType)
	return &result, nil
}
func (repository repositoryProductTypes) FindByID(id int) (*productTypes.Domain, error) {
	var recordProductType ProductTypes
	if err := repository.DB.Where("id = ?", id).First(&recordProductType).Error; err != nil {
		return &productTypes.Domain{}, err
	}
	result := toDomain(recordProductType)
	return &result, nil
}
