package products

import (
	"fmt"
	"outlet/v1/bussiness/products"
	"outlet/v1/repository/mysql/productTypes"

	"gorm.io/gorm"
)

type Products struct {
	gorm.Model
	ID           uint `gorm:"primaryKey"`
	TypeID       int
	Name         string
	ProductType  string
	Description  string
	Price        int
	ProductTypes productTypes.ProductTypes `gorm:"foreignKey:type_id"`
}

func toDomain(rec Products) products.Domain {
	fmt.Printf("%+v", rec)
	return products.Domain{
		ID:          int(rec.ID),
		Name:        rec.Name,
		ProductType: rec.ProductTypes.Name,
		Description: rec.Description,
		Price:       rec.Price,
		CreatedAt:   rec.CreatedAt,
		UpdatedAt:   rec.UpdatedAt,
	}
}

func fromDomain(domain products.Domain) Products {
	return Products{
		ID:          uint(domain.ID),
		Name:        domain.Name,
		TypeID:      domain.TypeID,
		Description: domain.Description,
		Price:       domain.Price,
	}
}
