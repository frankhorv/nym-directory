// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	models "github.com/nymtech/nym-directory/models"
	mock "github.com/stretchr/testify/mock"
)

// MixHostSanitizer is an autogenerated mock type for the MixHostSanitizer type
type MixHostSanitizer struct {
	mock.Mock
}

// Sanitize provides a mock function with given fields: _a0
func (_m *MixHostSanitizer) Sanitize(_a0 models.MixHostInfo) models.MixHostInfo {
	ret := _m.Called(_a0)

	var r0 models.MixHostInfo
	if rf, ok := ret.Get(0).(func(models.MixHostInfo) models.MixHostInfo); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(models.MixHostInfo)
	}

	return r0
}
