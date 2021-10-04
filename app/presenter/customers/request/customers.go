package request

import "outlet/v1/bussiness/customers"

type Customer struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Alamat    string `json:"alamat"`
	NoTelepon string `json:"noTelepon"`
}

type CustomerLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func ToDomain(request Customer) *customers.Domain {
	return &customers.Domain{
		Name:      request.Name,
		Email:     request.Email,
		Password:  request.Password,
		Alamat:    request.Alamat,
		NoTelepon: request.NoTelepon,
	}
}
