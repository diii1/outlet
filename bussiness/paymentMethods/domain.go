package paymentMethods

import (
	"time"
)

type Domain struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Service interface {
	AddPaymentMethod(paymentMethod *Domain) (*Domain, error)
	FindByID(id int) (*Domain, error)
	DeletePaymentMethod(id int, paymentMethod *Domain) (*Domain, error)
}

type Repository interface {
	Insert(productType *Domain) (*Domain, error)
	FindByID(id int) (*Domain, error)
	Delete(id int, paymentMethod *Domain) (*Domain, error)
}
