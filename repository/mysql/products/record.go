package products

import (
	"outlet/v1/bussiness/products"

	"gorm.io/gorm"
)

type Products struct {
	gorm.Model
	TypeID      int
	Name        string
	Description string
	Price       int
}

func toDomain(record Products) products.Domain {
	return products.Domain{
		ID:          int(record.ID),
		TypeID:      record.TypeID,
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
