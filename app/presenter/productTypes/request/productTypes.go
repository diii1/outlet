package request

import "github.com/diii1/outlet/bussiness/productTypes"

type ProductType struct {
	Name string `json:"name"`
}

func ToDomain(request ProductType) *productTypes.Domain {
	return &productTypes.Domain{
		Name: request.Name,
	}
}
