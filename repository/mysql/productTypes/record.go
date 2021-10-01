package productTypes

import (
	"github.com/diii1/outlet/bussiness/productTypes"
	"gorm.io/gorm"
)

type ProductTypes struct {
	gorm.Model
	Name string
}

func toDomain(record ProductTypes) productTypes.Domain {
	return productTypes.Domain{
		ID:        int(record.ID),
		Name:      record.Name,
		CreatedAt: record.CreatedAt,
		UpdatedAt: record.UpdatedAt,
	}
}

func fromDomain(domain productTypes.Domain) ProductTypes {
	return ProductTypes{
		Name: domain.Name,
	}
}
