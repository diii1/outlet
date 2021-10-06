package paymentMethods

type servicePaymentMethod struct {
	repository Repository
}

func NewService(repositoryPaymentMethod Repository) Service {
	return &servicePaymentMethod{
		repository: repositoryPaymentMethod,
	}
}

func (service *servicePaymentMethod) AddPaymentMethod(paymentMethod *Domain) (*Domain, error) {
	result, err := service.repository.Insert(paymentMethod)
	if err != nil {
		return &Domain{}, err
	}
	return result, nil
}

func (service *servicePaymentMethod) FindByID(id int) (*Domain, error) {
	paymentMethod, err := service.repository.FindByID(id)
	if err != nil {
		return &Domain{}, err
	}
	return paymentMethod, nil
}

func (service *servicePaymentMethod) DeletePaymentMethod(id int, paymentMethod *Domain) (*Domain, error) {
	result, err := service.repository.Delete(id, paymentMethod)
	if err != nil {
		return &Domain{}, err
	}
	return result, nil
}

func (service *servicePaymentMethod) GetAllPaymentMethod() (*[]Domain, error) {
	result, err := service.repository.GetAllPaymentMethod()
	if err != nil {
		return &[]Domain{}, err
	}
	return result, nil
}
