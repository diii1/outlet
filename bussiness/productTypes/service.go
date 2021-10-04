package productTypes

type serviceProductTypes struct {
	repository Repository
}

func NewService(repositoryProductType Repository) Service {
	return &serviceProductTypes{
		repository: repositoryProductType,
	}
}

func (service *serviceProductTypes) AddProductType(productType *Domain) (*Domain, error) {
	result, err := service.repository.Insert(productType)
	if err != nil {
		return &Domain{}, err
	}
	return result, nil
}

func (service *serviceProductTypes) FindByID(id int) (*Domain, error) {
	productType, err := service.repository.FindByID(id)
	if err != nil {
		return &Domain{}, err
	}
	return productType, nil
}

func (s *serviceProductTypes) DeleteProductType(id int, productType *Domain) (*Domain, error) {
	result, err := s.repository.Delete(id, productType)
	if err != nil {
		return &Domain{}, err
	}
	return result, nil
}
