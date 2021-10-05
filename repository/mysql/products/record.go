package products

import (
	"outlet/v1/bussiness/products"
	"outlet/v1/repository/mysql/productTypes"

	"gorm.io/gorm"
)

type Products struct {
	gorm.Model
	ID           uint `gorm:"primaryKey"`
	TypeID       int
	Name         string
	Description  string
	Price        int
	ProductTypes productTypes.ProductTypes `gorm:"foreignKey:type_id"`
}

func toDomain(record Products) products.Domain {
	return products.Domain{
		ID:          int(record.ID),
		TypeID:      int(record.ProductTypes.ID),
		Name:        record.Name,
		Description: record.Description,
		Price:       record.Price,
		CreatedAt:   record.CreatedAt,
		UpdatedAt:   record.UpdatedAt,
	}
}

func fromDomain(domain products.Domain) Products {
	return Products{
		ID:          uint(domain.ID),
		TypeID:      domain.TypeID,
		Name:        domain.Name,
		Description: domain.Description,
		Price:       domain.Price,
	}
}
