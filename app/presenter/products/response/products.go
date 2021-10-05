package response

import (
	"outlet/v1/bussiness/products"
	"time"
)

type Products struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	TypeID      int       `json:"type_id"`
	Description string    `json:"description"`
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
		TypeID:      domain.TypeID,
		Description: domain.Description,
		Price:       domain.Price,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
	}
}
