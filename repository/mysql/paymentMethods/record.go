package paymentMethods

import (
	"outlet/v1/bussiness/paymentMethods"

	"gorm.io/gorm"
)

type PaymentMethods struct {
	gorm.Model
	ID   uint `gorm:"primaryKey"`
	Name string
}

func toDomain(record PaymentMethods) paymentMethods.Domain {
	return paymentMethods.Domain{
		ID:        int(record.ID),
		Name:      record.Name,
		CreatedAt: record.CreatedAt,
		UpdatedAt: record.UpdatedAt,
	}
}

func fromDomain(domain paymentMethods.Domain) PaymentMethods {
	return PaymentMethods{
		ID:   uint(domain.ID),
		Name: domain.Name,
	}
}
