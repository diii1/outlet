package customers

import (
	"outlet/v1/bussiness/customers"

	"gorm.io/gorm"
)

type Customers struct {
	gorm.Model
	ID        uint `gorm:"primaryKey"`
	Name      string
	Email     string
	Password  string
	Alamat    string
	NoTelepon string
}

func toDomain(record Customers) customers.Domain {
	return customers.Domain{
		ID:        int(record.ID),
		Name:      record.Name,
		Email:     record.Email,
		Password:  record.Password,
		Alamat:    record.Alamat,
		NoTelepon: record.NoTelepon,
		CreatedAt: record.CreatedAt,
		UpdatedAt: record.UpdatedAt,
	}
}

func fromDomain(domain customers.Domain) Customers {
	return Customers{
		ID:        uint(domain.ID),
		Name:      domain.Name,
		Email:     domain.Email,
		Password:  domain.Password,
		Alamat:    domain.Alamat,
		NoTelepon: domain.NoTelepon,
	}
}

func toDomainArray(record []Customers) []customers.Domain {
	var res []customers.Domain
	for _, v := range record {
		res = append(res, toDomain(v))
	}
	return res
}
