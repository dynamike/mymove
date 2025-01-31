// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	adminmessages "github.com/transcom/mymove/pkg/gen/adminmessages"

	models "github.com/transcom/mymove/pkg/models"

	uuid "github.com/gofrs/uuid"

	validate "github.com/gobuffalo/validate/v3"
)

// OfficeUserUpdater is an autogenerated mock type for the OfficeUserUpdater type
type OfficeUserUpdater struct {
	mock.Mock
}

// UpdateOfficeUser provides a mock function with given fields: id, payload
func (_m *OfficeUserUpdater) UpdateOfficeUser(id uuid.UUID, payload *adminmessages.OfficeUserUpdatePayload) (*models.OfficeUser, *validate.Errors, error) {
	ret := _m.Called(id, payload)

	var r0 *models.OfficeUser
	if rf, ok := ret.Get(0).(func(uuid.UUID, *adminmessages.OfficeUserUpdatePayload) *models.OfficeUser); ok {
		r0 = rf(id, payload)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.OfficeUser)
		}
	}

	var r1 *validate.Errors
	if rf, ok := ret.Get(1).(func(uuid.UUID, *adminmessages.OfficeUserUpdatePayload) *validate.Errors); ok {
		r1 = rf(id, payload)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*validate.Errors)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(uuid.UUID, *adminmessages.OfficeUserUpdatePayload) error); ok {
		r2 = rf(id, payload)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}
