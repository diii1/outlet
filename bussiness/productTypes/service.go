package productTypes

type serviceProductTypes struct {
	repository Repository
}

func NewService(repoProductTypes Repository) Service {
	return &serviceProductTypes{
		repository: repoProductTypes,
	}
}

func (s *serviceProductTypes) AddProductType(productType *Domain) (*Domain, error) {
	result, err := s.repository.Insert(productType)
	if err != nil {
		return &Domain{}, err
	}
	return result, nil
}
func (s *serviceProductTypes) Update(id int, productType *Domain) (*Domain, error) {
	result, err := s.repository.Update(id, productType)
	if err != nil {
		return &Domain{}, err
	}
	return result, nil
}
func (s *serviceProductTypes) FindByID(id int) (*Domain, error) {
	productType, err := s.repository.FindByID(id)
	if err != nil {
		return &Domain{}, err
	}
	return productType, nil
}
