package response

import (
	"outlet/v1/bussiness/products"
	"time"
)

type Products struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	ProductType string    `json:"product_type"`
	Price       int       `json:"price"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Delete struct {
	Data string `json:"data"`
}

func FromDomain(domain products.Domain) Products {
	return Products{
		ID:          domain.ID,
		Name:        domain.Name,
		Description: domain.Description,
		ProductType: domain.ProductType,
		Price:       domain.Price,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
	}
}
