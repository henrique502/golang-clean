// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	team "github.com/henrique502/golang-clean/domain/team"
	mock "github.com/stretchr/testify/mock"
)

// Service is an autogenerated mock type for the Service type
type Service struct {
	mock.Mock
}

// GetTeamList provides a mock function with given fields: nextURL, query
func (_m *Service) GetTeamList(nextURL *string, query map[string]string) ([]team.Team, *string, error) {
	ret := _m.Called(nextURL, query)

	var r0 []team.Team
	if rf, ok := ret.Get(0).(func(*string, map[string]string) []team.Team); ok {
		r0 = rf(nextURL, query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]team.Team)
		}
	}

	var r1 *string
	if rf, ok := ret.Get(1).(func(*string, map[string]string) *string); ok {
		r1 = rf(nextURL, query)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*string)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(*string, map[string]string) error); ok {
		r2 = rf(nextURL, query)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}
