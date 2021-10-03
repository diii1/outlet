package response

import (
	"time"

	"github.com/diii1/outlet/bussiness/customers"
)

type Customer struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Alamat    string    `json:"alamat"`
	NoTelepon string    `json:"noTelepon"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromDomain(domain customers.Domain) Customer {
	return Customer{
		ID:        domain.ID,
		Name:      domain.Name,
		Alamat:    domain.Alamat,
		NoTelepon: domain.NoTelepon,
		Email:     domain.Email,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
