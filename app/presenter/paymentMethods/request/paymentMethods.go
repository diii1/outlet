package request

import (
	"outlet/v1/bussiness/paymentMethods"
)

type PaymentMethod struct {
	Name string `json:"name"`
}

func ToDomain(request PaymentMethod) *paymentMethods.Domain {
	return &paymentMethods.Domain{
		Name: request.Name,
	}
}
