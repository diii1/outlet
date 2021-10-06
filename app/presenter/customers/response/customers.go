package response

import (
	"outlet/v1/bussiness/customers"
	"time"
)

type Customer struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Alamat    string    `json:"alamat"`
	NoTelepon string    `json:"noTelepon"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CustomerLogin struct {
	ID    int    `json:"id"`
	Token string `json:"token"`
}

type Delete struct {
	Data string `json:"data"`
}

func FromDomain(domain customers.Domain) Customer {
	return Customer{
		ID:        domain.ID,
		Name:      domain.Name,
		Email:     domain.Email,
		Alamat:    domain.Alamat,
		NoTelepon: domain.NoTelepon,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func FromDomainArray(domain []customers.Domain) []Customer {
	var res []Customer
	for _, v := range domain {
		res = append(res, FromDomain(v))
	}
	return res
}
