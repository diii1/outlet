package products

import "time"

type Domain struct {
	ID          int
	TypeID      int
	Name        string
	Description string
	Price       int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Service interface {
	AddProduct(product *Domain) (*Domain, error)
	Update(id int, product *Domain) (*Domain, error)
	FindByID(id int) (*Domain, error)
}

type Repository interface {
	Insert(product *Domain) (*Domain, error)
	Update(id int, product *Domain) (*Domain, error)
	FindByID(id int) (*Domain, error)
}
