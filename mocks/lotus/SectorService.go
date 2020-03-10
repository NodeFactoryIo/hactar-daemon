// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// SectorService is an autogenerated mock type for the SectorService type
type SectorService struct {
	mock.Mock
}

// GetNumberOfSectors provides a mock function with given fields: miner
func (_m *SectorService) GetNumberOfSectors(miner string) (int, error) {
	ret := _m.Called(miner)

	var r0 int
	if rf, ok := ret.Get(0).(func(string) int); ok {
		r0 = rf(miner)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(miner)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetSectorSize provides a mock function with given fields: miner
func (_m *SectorService) GetSectorSize(miner string) (string, error) {
	ret := _m.Called(miner)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(miner)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(miner)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}