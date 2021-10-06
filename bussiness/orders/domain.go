package orders

import "time"

type Domain struct {
	ID         int
	ProductID  int
	Product    string
	PaymentID  int
	Payment    string
	CustomerID int
	Customer   string
	TotalPrice int
	Status     string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Service interface {
	AddOrder(order *Domain) (*Domain, error)
	FindByID(id int) (*Domain, error)
	DeleteOrder(id int, order *Domain) (*Domain, error)
	GetAllOrder() (*[]Domain, error)
}

type Repository interface {
	Insert(order *Domain) (*Domain, error)
	FindByID(id int) (*Domain, error)
	Delete(id int, order *Domain) (*Domain, error)
	GetAllOrder() (*[]Domain, error)
}
