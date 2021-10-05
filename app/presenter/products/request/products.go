package request

import (
	"outlet/v1/bussiness/products"
)

type Product struct {
	TypeID      int    `json:"type_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}

func ToDomain(request Product) *products.Domain {
	return &products.Domain{
		TypeID:      request.TypeID,
		Name:        request.Name,
		Description: request.Description,
		Price:       request.Price,
	}
}
