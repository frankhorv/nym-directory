// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	models "github.com/nymtech/nym-directory/models"
	mock "github.com/stretchr/testify/mock"
)

// IService is an autogenerated mock type for the IService type
type IService struct {
	mock.Mock
}

// AddCocoNodePresence provides a mock function with given fields: info, ip
func (_m *IService) AddCocoNodePresence(info models.CocoHostInfo, ip string) {
	_m.Called(info, ip)
}

// AddGatewayPresence provides a mock function with given fields: info
func (_m *IService) AddGatewayPresence(info models.MixProviderHostInfo) {
	_m.Called(info)
}

// AddMixNodePresence provides a mock function with given fields: info
func (_m *IService) AddMixNodePresence(info models.MixHostInfo) {
	_m.Called(info)
}

// AddMixProviderPresence provides a mock function with given fields: info
func (_m *IService) AddMixProviderPresence(info models.MixProviderHostInfo) {
	_m.Called(info)
}

// Topology provides a mock function with given fields:
func (_m *IService) Topology() models.Topology {
	ret := _m.Called()

	var r0 models.Topology
	if rf, ok := ret.Get(0).(func() models.Topology); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(models.Topology)
	}

	return r0
}
