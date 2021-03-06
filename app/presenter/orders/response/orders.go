package response

import (
	"outlet/v1/bussiness/orders"
	"time"
)

type Orders struct {
	ID         int       `json:"id"`
	Product    string    `json:"product"`
	Payment    string    `json:"payment_method"`
	Customer   string    `json:"customer"`
	TotalPrice int       `json:"totalPrice"`
	Status     string    `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type Delete struct {
	Data string `json:"data"`
}

func FromDomain(domain orders.Domain) Orders {
	return Orders{
		ID:         domain.ID,
		Product:    domain.Product,
		Payment:    domain.Payment,
		Customer:   domain.Customer,
		TotalPrice: domain.TotalPrice,
		Status:     domain.Status,
		CreatedAt:  domain.CreatedAt,
		UpdatedAt:  domain.UpdatedAt,
	}
}

func FromDomainArray(domain []orders.Domain) []Orders {
	var res []Orders
	for _, v := range domain {
		res = append(res, FromDomain(v))
	}
	return res
}
