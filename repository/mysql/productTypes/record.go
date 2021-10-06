package productTypes

import (
	"outlet/v1/bussiness/productTypes"

	"gorm.io/gorm"
)

type ProductTypes struct {
	gorm.Model
	ID   uint `gorm:"primaryKey"`
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
		ID:   uint(domain.ID),
		Name: domain.Name,
	}
}

func toDomainArray(record []ProductTypes) []productTypes.Domain {
	var res []productTypes.Domain
	for _, v := range record {
		res = append(res, toDomain(v))
	}
	return res
}
