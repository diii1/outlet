package products

import (
	"outlet/v1/bussiness/productTypes"
)

type serviceProducts struct {
	repository         Repository
	productTypeService productTypes.Service
}

func NewService(repositoryProducts Repository, serviceProductType productTypes.Service) Service {
	return &serviceProducts{
		repository:         repositoryProducts,
		productTypeService: serviceProductType,
	}
}

func (service *serviceProducts) AddProduct(products *Domain) (*Domain, error) {
	productType, err := service.productTypeService.FindByID(products.TypeID)
	if err != nil {
		return &Domain{}, err
	}
	products.TypeID = productType.ID
	result, err := service.repository.Insert(products)
	if err != nil {
		return &Domain{}, err
	}
	return result, nil
}

// func (service *serviceProducts) FindByID(id int) (*Domain, error) {
// 	product, err := service.repository.FindByID(id)
// 	if err != nil {
// 		return &Domain{}, err
// 	}
// 	return product, nil
// }

// func (s *serviceProducts) DeleteProduct(id int, product *Domain) (*Domain, error) {
// 	result, err := s.repository.Delete(id, product)
// 	if err != nil {
// 		return &Domain{}, err
// 	}
// 	return result, nil
// }
