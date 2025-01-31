// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	models "github.com/transcom/mymove/pkg/models"
)

// ShipmentRouter is an autogenerated mock type for the ShipmentRouter type
type ShipmentRouter struct {
	mock.Mock
}

// Approve provides a mock function with given fields: shipment
func (_m *ShipmentRouter) Approve(shipment *models.MTOShipment) error {
	ret := _m.Called(shipment)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.MTOShipment) error); ok {
		r0 = rf(shipment)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ApproveDiversion provides a mock function with given fields: shipment
func (_m *ShipmentRouter) ApproveDiversion(shipment *models.MTOShipment) error {
	ret := _m.Called(shipment)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.MTOShipment) error); ok {
		r0 = rf(shipment)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Cancel provides a mock function with given fields: shipment
func (_m *ShipmentRouter) Cancel(shipment *models.MTOShipment) error {
	ret := _m.Called(shipment)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.MTOShipment) error); ok {
		r0 = rf(shipment)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Reject provides a mock function with given fields: shipment, rejectionReason
func (_m *ShipmentRouter) Reject(shipment *models.MTOShipment, rejectionReason *string) error {
	ret := _m.Called(shipment, rejectionReason)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.MTOShipment, *string) error); ok {
		r0 = rf(shipment, rejectionReason)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RequestCancellation provides a mock function with given fields: shipment
func (_m *ShipmentRouter) RequestCancellation(shipment *models.MTOShipment) error {
	ret := _m.Called(shipment)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.MTOShipment) error); ok {
		r0 = rf(shipment)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RequestDiversion provides a mock function with given fields: shipment
func (_m *ShipmentRouter) RequestDiversion(shipment *models.MTOShipment) error {
	ret := _m.Called(shipment)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.MTOShipment) error); ok {
		r0 = rf(shipment)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Submit provides a mock function with given fields: shipment
func (_m *ShipmentRouter) Submit(shipment *models.MTOShipment) error {
	ret := _m.Called(shipment)

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.MTOShipment) error); ok {
		r0 = rf(shipment)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
