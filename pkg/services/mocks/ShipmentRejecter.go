// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	models "github.com/transcom/mymove/pkg/models"

	uuid "github.com/gofrs/uuid"
)

// ShipmentRejecter is an autogenerated mock type for the ShipmentRejecter type
type ShipmentRejecter struct {
	mock.Mock
}

// RejectShipment provides a mock function with given fields: shipmentID, eTag, reason
func (_m *ShipmentRejecter) RejectShipment(shipmentID uuid.UUID, eTag string, reason *string) (*models.MTOShipment, error) {
	ret := _m.Called(shipmentID, eTag, reason)

	var r0 *models.MTOShipment
	if rf, ok := ret.Get(0).(func(uuid.UUID, string, *string) *models.MTOShipment); ok {
		r0 = rf(shipmentID, eTag, reason)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.MTOShipment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uuid.UUID, string, *string) error); ok {
		r1 = rf(shipmentID, eTag, reason)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
