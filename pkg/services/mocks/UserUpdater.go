// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	models "github.com/transcom/mymove/pkg/models"

	uuid "github.com/gofrs/uuid"

	validate "github.com/gobuffalo/validate/v3"
)

// UserUpdater is an autogenerated mock type for the UserUpdater type
type UserUpdater struct {
	mock.Mock
}

// UpdateUser provides a mock function with given fields: id, user
func (_m *UserUpdater) UpdateUser(id uuid.UUID, user *models.User) (*models.User, *validate.Errors, error) {
	ret := _m.Called(id, user)

	var r0 *models.User
	if rf, ok := ret.Get(0).(func(uuid.UUID, *models.User) *models.User); ok {
		r0 = rf(id, user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.User)
		}
	}

	var r1 *validate.Errors
	if rf, ok := ret.Get(1).(func(uuid.UUID, *models.User) *validate.Errors); ok {
		r1 = rf(id, user)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*validate.Errors)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(uuid.UUID, *models.User) error); ok {
		r2 = rf(id, user)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}
