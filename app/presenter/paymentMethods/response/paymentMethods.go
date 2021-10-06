package response

import (
	"outlet/v1/bussiness/paymentMethods"
	"time"
)

type PaymentMethod struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Delete struct {
	Data string `json:"data"`
}

func FromDomain(domain paymentMethods.Domain) PaymentMethod {
	return PaymentMethod{
		ID:        domain.ID,
		Name:      domain.Name,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}

func FromDomainArray(domain []paymentMethods.Domain) []PaymentMethod {
	var res []PaymentMethod
	for _, v := range domain {
		res = append(res, FromDomain(v))
	}
	return res
}
