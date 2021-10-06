// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	customers "outlet/v1/bussiness/customers"

	mock "github.com/stretchr/testify/mock"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// AddCustomer provides a mock function with given fields: customer
func (_m *Service) AddCustomer(customer *customers.Domain) (*customers.Domain, error) {
	ret := _m.Called(customer)

	var r0 *customers.Domain
	if rf, ok := ret.Get(0).(func(*customers.Domain) *customers.Domain); ok {
		r0 = rf(customer)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*customers.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*customers.Domain) error); ok {
		r1 = rf(customer)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteCustomer provides a mock function with given fields: id, customer
func (_m *Service) DeleteCustomer(id int, customer *customers.Domain) (*customers.Domain, error) {
	ret := _m.Called(id, customer)

	var r0 *customers.Domain
	if rf, ok := ret.Get(0).(func(int, *customers.Domain) *customers.Domain); ok {
		r0 = rf(id, customer)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*customers.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, *customers.Domain) error); ok {
		r1 = rf(id, customer)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByID provides a mock function with given fields: id
func (_m *Service) FindByID(id int) (*customers.Domain, error) {
	ret := _m.Called(id)

	var r0 *customers.Domain
	if rf, ok := ret.Get(0).(func(int) *customers.Domain); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*customers.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAllCustomer provides a mock function with given fields:
func (_m *Service) GetAllCustomer() (*[]customers.Domain, error) {
	ret := _m.Called()

	var r0 *[]customers.Domain
	if rf, ok := ret.Get(0).(func() *[]customers.Domain); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]customers.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Login provides a mock function with given fields: email, password
func (_m *Service) Login(email string, password string) (string, error) {
	ret := _m.Called(email, password)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(email, password)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(email, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: id, user
func (_m *Service) Update(id int, user *customers.Domain) (*customers.Domain, error) {
	ret := _m.Called(id, user)

	var r0 *customers.Domain
	if rf, ok := ret.Get(0).(func(int, *customers.Domain) *customers.Domain); ok {
		r0 = rf(id, user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*customers.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, *customers.Domain) error); ok {
		r1 = rf(id, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
