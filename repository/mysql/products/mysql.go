package products

import (
	"outlet/v1/bussiness/products"

	"gorm.io/gorm"
)

type repoProducts struct {
	DB *gorm.DB
}

func NewRepositoryMySQL(db *gorm.DB) products.Repository {
	return &repoProducts{
		DB: db,
	}
}

func (repo *repoProducts) Insert(product *products.Domain) (*products.Domain, error) {
	recordProduct := fromDomain(*product)
	if err := repo.DB.Create(&recordProduct).Error; err != nil {
		return &products.Domain{}, err
	}

	record, err := repo.FindByID(int(recordProduct.ID))
	if err != nil {
		return &products.Domain{}, err
	}
	return record, nil
}

func (repo *repoProducts) FindByID(id int) (*products.Domain, error) {
	var recordProduct Products

	if err := repo.DB.Where("products.id = ?", id).Joins("ProductTypes").Find(&recordProduct).Error; err != nil {
		return &products.Domain{}, err
	}
	result := toDomain(recordProduct)
	return &result, nil
}
