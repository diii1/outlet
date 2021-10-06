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

func (service *serviceProducts) AddProduct(product *Domain) (*Domain, error) {
	respo, err := service.productTypeDomain.FindByID(product.TypeID)
	if err != nil {
		return &Domain{}, err
	}
	product.ProductType = respo.Name
	result, err := service.repository.Insert(product)
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

func (service *serviceProducts) Update(id int, product *Domain) (*Domain, error) {
	respo, err := service.productTypeDomain.FindByID(product.TypeID)
	if err != nil {
		return &Domain{}, err
	}
	product.ProductType = respo.Name
	result, err := service.repository.Update(id, product)
	if err != nil {
		return &Domain{}, err
	}
	return result, nil
}

func (service *serviceProducts) DeleteProduct(id int, product *Domain) (*Domain, error) {
	result, err := service.repository.Delete(id, product)
	if err != nil {
		return &Domain{}, err
	}
	return result, nil
}
