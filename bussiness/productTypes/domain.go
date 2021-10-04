package productTypes

import "time"

type Domain struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Service interface {
	AddProductType(productType *Domain) (*Domain, error)
	FindByID(id int) (*Domain, error)
	DeleteProductType(id int, productType *Domain) (*Domain, error)
}

type Repository interface {
	Insert(productType *Domain) (*Domain, error)
	FindByID(id int) (*Domain, error)
	Delete(id int, productType *Domain) (*Domain, error)
}
