package request

import "github.com/diii1/outlet/bussiness/customers"

type Customer struct {
	Name      string `json:"name"`
	Alamat    string `json:"alamat"`
	NoTelepon string `json:"noTelepon"`
	Email     string `json:"email"`
}

func ToDomain(request Customer) *customers.Domain {
	return &customers.Domain{
		Name:      request.Name,
		Alamat:    request.Alamat,
		NoTelepon: request.NoTelepon,
		Email:     request.Email,
	}
}
