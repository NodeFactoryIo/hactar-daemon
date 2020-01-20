// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import hactar "github.com/NodeFactoryIo/hactar-daemon/internal/hactar"
import http "net/http"
import mock "github.com/stretchr/testify/mock"

// NodesService is an autogenerated mock type for the NodesService type
type NodesService struct {
	mock.Mock
}

// Add provides a mock function with given fields: node
func (_m *NodesService) Add(node hactar.Node) (*hactar.Node, *http.Response, error) {
	ret := _m.Called(node)

	var r0 *hactar.Node
	if rf, ok := ret.Get(0).(func(hactar.Node) *hactar.Node); ok {
		r0 = rf(node)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*hactar.Node)
		}
	}

	var r1 *http.Response
	if rf, ok := ret.Get(1).(func(hactar.Node) *http.Response); ok {
		r1 = rf(node)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*http.Response)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(hactar.Node) error); ok {
		r2 = rf(node)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}