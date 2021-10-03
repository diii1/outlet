package response

import (
	"outlet/v1/bussiness/products"
	"time"
)

type Product struct {
	ID          int       `json:"id"`
	TypeID      int       `json:"typeID"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Price       int       `json:"price"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func FromDomain(domain products.Domain) Product {
	return Product{
		ID:          domain.ID,
		TypeID:      domain.TypeID,
		Name:        domain.Name,
		Description: domain.Description,
		Price:       domain.Price,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
	}
}
