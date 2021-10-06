package request

import "outlet/v1/bussiness/orders"

type Orders struct {
	ProductID  int `json:"product_id"`
	PaymentID  int `json:"payment_id"`
	CustomerID int `json:"customer_id"`
}

func ToDomain(request Orders) *orders.Domain {
	return &orders.Domain{
		ProductID:  request.ProductID,
		PaymentID:  request.PaymentID,
		CustomerID: request.CustomerID,
	}
}
