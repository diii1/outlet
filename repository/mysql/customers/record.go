package customers

import (
	"github.com/diii1/outlet/bussiness/customers"
	"gorm.io/gorm"
)

type Customers struct {
	gorm.Model
	Name      string
	Alamat    string
	NoTelepon string
	Email     string
}

func toDomain(record Customers) customers.Domain {
	return customers.Domain{
		ID:        int(record.ID),
		Name:      record.Name,
		Alamat:    record.Alamat,
		NoTelepon: record.NoTelepon,
		Email:     record.Email,
	}
}

func fromDomain(domain customers.Domain) Customers {
	return Customers{
		Name:      domain.Name,
		Alamat:    domain.Alamat,
		NoTelepon: domain.NoTelepon,
		Email:     domain.Email,
	}
}
