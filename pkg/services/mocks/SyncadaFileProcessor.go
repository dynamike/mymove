// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"
	models "github.com/transcom/mymove/pkg/models"
)

// SyncadaFileProcessor is an autogenerated mock type for the SyncadaFileProcessor type
type SyncadaFileProcessor struct {
	mock.Mock
}

// EDIType provides a mock function with given fields:
func (_m *SyncadaFileProcessor) EDIType() models.EDIType {
	ret := _m.Called()

	var r0 models.EDIType
	if rf, ok := ret.Get(0).(func() models.EDIType); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(models.EDIType)
	}

	return r0
}

// ProcessFile provides a mock function with given fields: syncadaPath, text
func (_m *SyncadaFileProcessor) ProcessFile(syncadaPath string, text string) error {
	ret := _m.Called(syncadaPath, text)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(syncadaPath, text)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
