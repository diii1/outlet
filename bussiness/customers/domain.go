package customers

import "time"

type Domain struct {
	ID        int
	Name      string
	Alamat    string
	NoTelepon string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Service interface {
	AddCustomer(productType *Domain) (*Domain, error)
	Update(id int, productType *Domain) (*Domain, error)
	FindByID(id int) (*Domain, error)
}

type Repository interface {
	Insert(productType *Domain) (*Domain, error)
	Update(id int, productType *Domain) (*Domain, error)
	FindByID(id int) (*Domain, error)
}
