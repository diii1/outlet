package products

import (
	"outlet/v1/bussiness/products"
	"outlet/v1/repository/mysql/productTypes"

	"gorm.io/gorm"
)

type Products struct {
	gorm.Model
	TypeID      int
	IdType      productTypes.ProductTypes `gorm:"foreignKey:type_id"`
	Name        string
	Description string
	Price       int
}

func toDomain(record Products) products.Domain {
	return products.Domain{
		ID:          int(record.ID),
		TypeID:      int(record.IdType.ID),
		Name:        record.Name,
		Description: record.Description,
		Price:       record.Price,
	}
}

func fromDomain(domain products.Domain) Products {
	return Products{
		TypeID:      domain.TypeID,
		Name:        domain.Name,
		Description: domain.Description,
		Price:       domain.Price,
	}
}
