package request

import "outlet/v1/bussiness/products"

type Products struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	TypeID      int    `json:"type_id"`
	Price       int    `json:"price"`
}

type ProductUpdate struct {
	Name        string `json:"name"`
	ProductType string `json:"product_type"`
}

func ToDomain(request Products) *products.Domain {
	return &products.Domain{
		Name:        request.Name,
		Description: request.Description,
		TypeID:      request.TypeID,
		Price:       request.Price,
	}
}
