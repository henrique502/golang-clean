// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	integration "github.com/henrique502/golang-clean/domain/integration"
	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// IntegrationUpSert provides a mock function with given fields: _a0
func (_m *Repository) IntegrationUpSert(_a0 integration.Integration) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(integration.Integration) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
