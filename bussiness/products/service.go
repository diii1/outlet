package products

type serviceProduct struct {
	repository Repository
}

func NewService(repoProduct Repository) Service {
	return &serviceProduct{
		repository: repoProduct,
	}
}

func (s *serviceProduct) AddProduct(product *Domain) (*Domain, error) {
	result, err := s.repository.Insert(product)
	if err != nil {
		return &Domain{}, err
	}
	return result, nil
}
func (s *serviceProduct) Update(id int, product *Domain) (*Domain, error) {
	result, err := s.repository.Update(id, product)
	if err != nil {
		return &Domain{}, err
	}
	return result, nil
}
func (s *serviceProduct) FindByID(id int) (*Domain, error) {
	product, err := s.repository.FindByID(id)
	if err != nil {
		return &Domain{}, err
	}
	return product, nil
}
