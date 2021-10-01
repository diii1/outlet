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
	Update(id int, productType *Domain) (*Domain, error)
	FindByID(id int) (*Domain, error)
}

type Repository interface {
	Insert(productType *Domain) (*Domain, error)
	Update(id int, productType *Domain) (*Domain, error)
	FindByID(id int) (*Domain, error)
}
