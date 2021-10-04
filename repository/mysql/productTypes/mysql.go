package productTypes

import (
	"outlet/v1/bussiness/productTypes"

	"gorm.io/gorm"
)

type repositoryProductType struct {
	DB *gorm.DB
}

func NewRepositoryMySQL(db *gorm.DB) productTypes.Repository {
	return &repositoryProductType{
		DB: db,
	}
}

func (repository repositoryProductType) Insert(productType *productTypes.Domain) (*productTypes.Domain, error) {
	recordProductType := fromDomain(*productType)
	if err := repository.DB.Create(&recordProductType).Error; err != nil {
		return &productTypes.Domain{}, err
	}
	result := toDomain(recordProductType)
	return &result, nil
}

func (repository repositoryProductType) FindByID(id int) (*productTypes.Domain, error) {
	recordProductType := ProductTypes{}
	if err := repository.DB.Where("id = ?", id).First(&recordProductType).Error; err != nil {
		return &productTypes.Domain{}, err
	}
	result := toDomain(recordProductType)
	return &result, nil
}

func (repository repositoryProductType) Delete(id int, productType *productTypes.Domain) (*productTypes.Domain, error) {
	recordProductType := fromDomain(*productType)
	if err := repository.DB.Where("id = ?", id).Delete(&recordProductType).Error; err != nil {
		return &productTypes.Domain{}, err
	}
	result := toDomain(recordProductType)
	return &result, nil
}
