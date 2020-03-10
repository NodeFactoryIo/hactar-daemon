// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import hactar "github.com/NodeFactoryIo/hactar-daemon/internal/hactar"
import http "net/http"
import mock "github.com/stretchr/testify/mock"

// BlocksService is an autogenerated mock type for the BlocksService type
type BlocksService struct {
	mock.Mock
}

// AddMiningReward provides a mock function with given fields: blocks
func (_m *BlocksService) AddMiningReward(blocks []hactar.Block) (*http.Response, error) {
	ret := _m.Called(blocks)

	var r0 *http.Response
	if rf, ok := ret.Get(0).(func([]hactar.Block) *http.Response); ok {
		r0 = rf(blocks)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]hactar.Block) error); ok {
		r1 = rf(blocks)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
