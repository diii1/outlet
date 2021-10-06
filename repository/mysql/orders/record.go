package orders

import (
	"fmt"
	"outlet/v1/bussiness/orders"
	"outlet/v1/repository/mysql/customers"
	"outlet/v1/repository/mysql/paymentMethods"
	"outlet/v1/repository/mysql/products"

	"gorm.io/gorm"
)

type Orders struct {
	gorm.Model
	ID             uint `gorm:"primaryKey"`
	ProductID      int
	Product        string
	Products       products.Products `gorm:"foreignKey:product_id"`
	PaymentID      int
	Payment        string
	PaymentMethods paymentMethods.PaymentMethods `gorm:"foreignKey:payment_id"`
	CustomerID     int
	Customer       string
	Customers      customers.Customers `gorm:"foreignKey:customer_id"`
	TotalPrice     int
	Status         string
}

func toDomain(rec Orders) orders.Domain {
	fmt.Printf("%+v", rec)
	return orders.Domain{
		ID:         int(rec.ID),
		Product:    rec.Products.Name,
		Payment:    rec.PaymentMethods.Name,
		Customer:   rec.Customers.Name,
		TotalPrice: rec.Products.Price,
		Status:     rec.Status,
		CreatedAt:  rec.CreatedAt,
		UpdatedAt:  rec.UpdatedAt,
	}
}

func fromDomain(domain orders.Domain) Orders {
	return Orders{
		ID:         uint(domain.ID),
		ProductID:  domain.ProductID,
		Product:    domain.Product,
		PaymentID:  domain.PaymentID,
		Payment:    domain.Payment,
		CustomerID: domain.CustomerID,
		Customer:   domain.Customer,
		TotalPrice: domain.TotalPrice,
		Status:     domain.Status,
	}
}
