package products

import (
	"outlet/v1/bussiness/products"

	"gorm.io/gorm"
)

type repositoryProducts struct {
	DB *gorm.DB
}

func NewRepositoryMySQL(db *gorm.DB) products.Repository {
	return &repositoryProducts{
		DB: db,
	}
}

func (repository repositoryProducts) Insert(product *products.Domain) (*products.Domain, error) {
	recordProduct := fromDomain(*product)
	if err := repository.DB.Create(&recordProduct).Error; err != nil {
		return &products.Domain{}, err
	}
	result := toDomain(recordProduct)
	return &result, nil
}

// func (repository repositoryProducts) GetProudctTyeID(productTypeID int) (*products.Domain, error) {
// 	recordProduct := Products{}
// 	if err := repository.DB.Where("type_id = ?", productTypeID).Joins("ProductTypes").Last(&recordProduct).Error; err != nil {
// 		return &products.Domain{}, err
// 	}
// 	result := toDomain(recordProduct)
// 	return &result, nil
// }

// func (repository repositoryProducts) FindByID(id int) (*products.Domain, error) {
// 	recordProduct := Products{}
// 	if err := repository.DB.Where("id = ?", id).Joins("ProductTypes").First(&recordProduct).Error; err != nil {
// 		return &products.Domain{}, err
// 	}
// 	result := toDomain(recordProduct)
// 	return &result, nil
// }

// func (repository repositoryProducts) Delete(id int, product *products.Domain) (*products.Domain, error) {
// 	recordProduct := fromDomain(*product)
// 	if err := repository.DB.Where("id = ?", id).Delete(&recordProduct).Error; err != nil {
// 		return &products.Domain{}, err
// 	}
// 	result := toDomain(recordProduct)
// 	return &result, nil
// }
