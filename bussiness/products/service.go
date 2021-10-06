package products

import "outlet/v1/bussiness/productTypes"

type serviceProducts struct {
	repository        Repository
	productTypeDomain productTypes.Service
}

func NewService(repoProduct Repository, productTypeServ productTypes.Service) Service {
	return &serviceProducts{
		repository:        repoProduct,
		productTypeDomain: productTypeServ,
	}
}

func (servProduct *serviceProducts) AddProduct(product *Domain) (*Domain, error) {
	respo, err := servProduct.productTypeDomain.FindByID(product.TypeID)
	if err != nil {
		return &Domain{}, err
	}
	product.ProductType = respo.Name
	result, err := servProduct.repository.Insert(product)
	if err != nil {
		return &Domain{}, err
	}
	return result, nil
}

func (service *serviceProducts) FindByID(id int) (*Domain, error) {
	product, err := service.repository.FindByID(id)
	if err != nil {
		return &Domain{}, err
	}
	return product, nil
}
