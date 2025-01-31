// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	models "github.com/transcom/mymove/pkg/models"

	uuid "github.com/gofrs/uuid"
)

// CustomerFetcher is an autogenerated mock type for the CustomerFetcher type
type CustomerFetcher struct {
	mock.Mock
}

// FetchCustomer provides a mock function with given fields: customerID
func (_m *CustomerFetcher) FetchCustomer(customerID uuid.UUID) (*models.ServiceMember, error) {
	ret := _m.Called(customerID)

	var r0 *models.ServiceMember
	if rf, ok := ret.Get(0).(func(uuid.UUID) *models.ServiceMember); ok {
		r0 = rf(customerID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.ServiceMember)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uuid.UUID) error); ok {
		r1 = rf(customerID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
