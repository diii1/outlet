package response

import (
	"outlet/v1/bussiness/productTypes"
	"time"
)

type ProductType struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Delete struct {
	Data string `json:"data"`
}

func FromDomain(domain productTypes.Domain) ProductType {
	return ProductType{
		ID:        domain.ID,
		Name:      domain.Name,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
