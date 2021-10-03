package request

import "outlet/v1/bussiness/productTypes"

type ProductType struct {
	Name string `json:"name"`
}

func ToDomain(request ProductType) *productTypes.Domain {
	return &productTypes.Domain{
		Name: request.Name,
	}
}
