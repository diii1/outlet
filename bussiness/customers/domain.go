package customers

import "time"

type Domain struct {
	ID        int
	Name      string
	Email     string
	Password  string
	Alamat    string
	NoTelepon string
	Token     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Service interface {
	AddCustomer(customer *Domain) (*Domain, error)
	Update(id int, user *Domain) (*Domain, error)
	FindByID(id int) (*Domain, error)
	Login(email string, password string) (string, error)
	DeleteCustomer(id int, customer *Domain) (*Domain, error)
	GetAllCustomer() (*[]Domain, error)
}

type Repository interface {
	Insert(customer *Domain) (*Domain, error)
	Update(id int, user *Domain) (*Domain, error)
	FindByID(id int) (*Domain, error)
	FindByEmail(email string) (*Domain, error)
	Delete(id int, customer *Domain) (*Domain, error)
	GetAllCustomer() (*[]Domain, error)
}
