package request

import "outlet/v1/bussiness/products"

type Products struct {
	TypeID      int    `json:"typeID"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Price       int    `json:"price"`
}

func ToDomain(request Products) *products.Domain {
	return &products.Domain{
		TypeID:      request.TypeID,
		Name:        request.Name,
		Description: request.Description,
		Price:       request.Price,
	}
}
