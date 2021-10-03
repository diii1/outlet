package productTypes

import (
	producttypes "outlet/v1/bussiness/productTypes"

	"gorm.io/gorm"
)

type repositoryProductTypes struct {
	DB *gorm.DB
}

func NewRepositoryMySQL(db *gorm.DB) producttypes.Repository {
	return &repositoryProductTypes{
		DB: db,
	}
}

func (repository repositoryProductTypes) Insert(productType *producttypes.Domain) (*producttypes.Domain, error) {
	recordProductType := fromDomain(*productType)
	if err := repository.DB.Create(&recordProductType).Error; err != nil {
		return &producttypes.Domain{}, err
	}
	result := toDomain(recordProductType)
	return &result, nil
}
func (repository repositoryProductTypes) Update(id int, productType *producttypes.Domain) (*producttypes.Domain, error) {
	recordProductType := fromDomain(*productType)
	if err := repository.DB.Where("id = ?", id).Updates(&recordProductType).Error; err != nil {
		return &producttypes.Domain{}, err
	}
	if err := repository.DB.Where("id = ?", id).First(&recordProductType).Error; err != nil {
		return &producttypes.Domain{}, err
	}
	result := toDomain(recordProductType)
	return &result, nil
}
func (repository repositoryProductTypes) FindByID(id int) (*producttypes.Domain, error) {
	var recordProductType ProductTypes
	if err := repository.DB.Where("id = ?", id).First(&recordProductType).Error; err != nil {
		return &producttypes.Domain{}, err
	}
	result := toDomain(recordProductType)
	return &result, nil
}
