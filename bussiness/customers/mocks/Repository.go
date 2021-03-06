// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	customers "outlet/v1/bussiness/customers"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// Delete provides a mock function with given fields: id, customer
func (_m *Repository) Delete(id int, customer *customers.Domain) (*customers.Domain, error) {
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

// FindByEmail provides a mock function with given fields: email
func (_m *Repository) FindByEmail(email string) (*customers.Domain, error) {
	ret := _m.Called(email)

	var r0 *customers.Domain
	if rf, ok := ret.Get(0).(func(string) *customers.Domain); ok {
		r0 = rf(email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*customers.Domain)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByID provides a mock function with given fields: id
func (_m *Repository) FindByID(id int) (*customers.Domain, error) {
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
func (_m *Repository) GetAllCustomer() (*[]customers.Domain, error) {
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

// Insert provides a mock function with given fields: customer
func (_m *Repository) Insert(customer *customers.Domain) (*customers.Domain, error) {
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

// Update provides a mock function with given fields: id, user
func (_m *Repository) Update(id int, user *customers.Domain) (*customers.Domain, error) {
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
