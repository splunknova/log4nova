// Code generated by MockGen. DO NOT EDIT.
// Source: vendor/github.com/Polarishq/logface-sdk-go/client/events/events_iface.go

// Package mock_events is a generated GoMock package.
package mock_events

import (
	. "github.com/Polarishq/logface-sdk-go/client/events"
	runtime "github.com/go-openapi/runtime"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockClientInterface is a mock of ClientInterface interface
type MockClientInterface struct {
	ctrl     *gomock.Controller
	recorder *MockClientInterfaceMockRecorder
}

// MockClientInterfaceMockRecorder is the mock recorder for MockClientInterface
type MockClientInterfaceMockRecorder struct {
	mock *MockClientInterface
}

// NewMockClientInterface creates a new mock instance
func NewMockClientInterface(ctrl *gomock.Controller) *MockClientInterface {
	mock := &MockClientInterface{ctrl: ctrl}
	mock.recorder = &MockClientInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClientInterface) EXPECT() *MockClientInterfaceMockRecorder {
	return m.recorder
}

// GetEvents mocks base method
func (m *MockClientInterface) GetEvents(params *GetEventsParams, authInfo runtime.ClientAuthInfoWriter) (*GetEventsOK, error) {
	ret := m.ctrl.Call(m, "GetEvents", params, authInfo)
	ret0, _ := ret[0].(*GetEventsOK)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEvents indicates an expected call of GetEvents
func (mr *MockClientInterfaceMockRecorder) GetEvents(params, authInfo interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEvents", reflect.TypeOf((*MockClientInterface)(nil).GetEvents), params, authInfo)
}

// Events mocks base method
func (m *MockClientInterface) Events(params *EventsParams, authInfo runtime.ClientAuthInfoWriter) (*EventsOK, error) {
	ret := m.ctrl.Call(m, "Events", params, authInfo)
	ret0, _ := ret[0].(*EventsOK)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Events indicates an expected call of Events
func (mr *MockClientInterfaceMockRecorder) Events(params, authInfo interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Events", reflect.TypeOf((*MockClientInterface)(nil).Events), params, authInfo)
}

// SetTransport mocks base method
func (m *MockClientInterface) SetTransport(transport runtime.ClientTransport) {
	m.ctrl.Call(m, "SetTransport", transport)
}

// SetTransport indicates an expected call of SetTransport
func (mr *MockClientInterfaceMockRecorder) SetTransport(transport interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTransport", reflect.TypeOf((*MockClientInterface)(nil).SetTransport), transport)
}
