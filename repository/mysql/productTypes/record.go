package productTypes

import (
	producttypes "outlet/v1/bussiness/productTypes"

	"gorm.io/gorm"
)

type ProductTypes struct {
	gorm.Model
	Name string
}

func toDomain(record ProductTypes) producttypes.Domain {
	return producttypes.Domain{
		ID:        int(record.ID),
		Name:      record.Name,
		CreatedAt: record.CreatedAt,
		UpdatedAt: record.UpdatedAt,
	}
}

func fromDomain(domain producttypes.Domain) ProductTypes {
	return ProductTypes{
		Name: domain.Name,
	}
}
