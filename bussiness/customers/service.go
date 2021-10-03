package customers

type serviceCustomers struct {
	repository Repository
}

func NewService(repoCustomers Repository) Service {
	return &serviceCustomers{
		repository: repoCustomers,
	}
}

func (s *serviceCustomers) AddCustomer(customers *Domain) (*Domain, error) {
	result, err := s.repository.Insert(customers)
	if err != nil {
		return &Domain{}, err
	}
	return result, nil
}
func (s *serviceCustomers) Update(id int, customers *Domain) (*Domain, error) {
	result, err := s.repository.Update(id, customers)
	if err != nil {
		return &Domain{}, err
	}
	return result, nil
}
func (s *serviceCustomers) FindByID(id int) (*Domain, error) {
	customers, err := s.repository.FindByID(id)
	if err != nil {
		return &Domain{}, err
	}
	return customers, nil
}
