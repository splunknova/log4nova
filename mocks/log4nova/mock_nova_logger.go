// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/splunknova/log4nova (interfaces: INovaLogger)

// Package mock_log4nova is a generated GoMock package.
package mock_log4nova

import (
	gomock "github.com/golang/mock/gomock"
	log4nova "github.com/splunknova/log4nova"
	logrus "github.com/sirupsen/logrus"
	reflect "reflect"
)

// MockINovaLogger is a mock of INovaLogger interface
type MockINovaLogger struct {
	ctrl     *gomock.Controller
	recorder *MockINovaLoggerMockRecorder
}

// MockINovaLoggerMockRecorder is the mock recorder for MockINovaLogger
type MockINovaLoggerMockRecorder struct {
	mock *MockINovaLogger
}

// NewMockINovaLogger creates a new mock instance
func NewMockINovaLogger(ctrl *gomock.Controller) *MockINovaLogger {
	mock := &MockINovaLogger{ctrl: ctrl}
	mock.recorder = &MockINovaLoggerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockINovaLogger) EXPECT() *MockINovaLoggerMockRecorder {
	return m.recorder
}

// Debug mocks base method
func (m *MockINovaLogger) Debug(arg0 ...interface{}) {
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Debug", varargs...)
}

// Debug indicates an expected call of Debug
func (mr *MockINovaLoggerMockRecorder) Debug(arg0 ...interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Debug", reflect.TypeOf((*MockINovaLogger)(nil).Debug), arg0...)
}

// Debugf mocks base method
func (m *MockINovaLogger) Debugf(arg0 string, arg1 ...interface{}) {
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Debugf", varargs...)
}

// Debugf indicates an expected call of Debugf
func (mr *MockINovaLoggerMockRecorder) Debugf(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Debugf", reflect.TypeOf((*MockINovaLogger)(nil).Debugf), varargs...)
}

// Error mocks base method
func (m *MockINovaLogger) Error(arg0 ...interface{}) {
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Error", varargs...)
}

// Error indicates an expected call of Error
func (mr *MockINovaLoggerMockRecorder) Error(arg0 ...interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Error", reflect.TypeOf((*MockINovaLogger)(nil).Error), arg0...)
}

// Errorf mocks base method
func (m *MockINovaLogger) Errorf(arg0 string, arg1 ...interface{}) {
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Errorf", varargs...)
}

// Errorf indicates an expected call of Errorf
func (mr *MockINovaLoggerMockRecorder) Errorf(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Errorf", reflect.TypeOf((*MockINovaLogger)(nil).Errorf), varargs...)
}

// Info mocks base method
func (m *MockINovaLogger) Info(arg0 ...interface{}) {
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Info", varargs...)
}

// Info indicates an expected call of Info
func (mr *MockINovaLoggerMockRecorder) Info(arg0 ...interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Info", reflect.TypeOf((*MockINovaLogger)(nil).Info), arg0...)
}

// Infof mocks base method
func (m *MockINovaLogger) Infof(arg0 string, arg1 ...interface{}) {
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Infof", varargs...)
}

// Infof indicates an expected call of Infof
func (mr *MockINovaLoggerMockRecorder) Infof(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Infof", reflect.TypeOf((*MockINovaLogger)(nil).Infof), varargs...)
}

// SetDebug mocks base method
func (m *MockINovaLogger) SetDebug() {
	m.ctrl.Call(m, "SetDebug")
}

// SetDebug indicates an expected call of SetDebug
func (mr *MockINovaLoggerMockRecorder) SetDebug() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDebug", reflect.TypeOf((*MockINovaLogger)(nil).SetDebug))
}

// SetError mocks base method
func (m *MockINovaLogger) SetError() {
	m.ctrl.Call(m, "SetError")
}

// SetError indicates an expected call of SetError
func (mr *MockINovaLoggerMockRecorder) SetError() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetError", reflect.TypeOf((*MockINovaLogger)(nil).SetError))
}

// SetInfo mocks base method
func (m *MockINovaLogger) SetInfo() {
	m.ctrl.Call(m, "SetInfo")
}

// SetInfo indicates an expected call of SetInfo
func (mr *MockINovaLoggerMockRecorder) SetInfo() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetInfo", reflect.TypeOf((*MockINovaLogger)(nil).SetInfo))
}

// SetWarn mocks base method
func (m *MockINovaLogger) SetWarn() {
	m.ctrl.Call(m, "SetWarn")
}

// SetWarn indicates an expected call of SetWarn
func (mr *MockINovaLoggerMockRecorder) SetWarn() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetWarn", reflect.TypeOf((*MockINovaLogger)(nil).SetWarn))
}

// Start mocks base method
func (m *MockINovaLogger) Start() {
	m.ctrl.Call(m, "Start")
}

// Start indicates an expected call of Start
func (mr *MockINovaLoggerMockRecorder) Start() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockINovaLogger)(nil).Start))
}

// Stop mocks base method
func (m *MockINovaLogger) Stop() {
	m.ctrl.Call(m, "Stop")
}

// Stop indicates an expected call of Stop
func (mr *MockINovaLoggerMockRecorder) Stop() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockINovaLogger)(nil).Stop))
}

// Warning mocks base method
func (m *MockINovaLogger) Warning(arg0 ...interface{}) {
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Warning", varargs...)
}

// Warning indicates an expected call of Warning
func (mr *MockINovaLoggerMockRecorder) Warning(arg0 ...interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Warning", reflect.TypeOf((*MockINovaLogger)(nil).Warning), arg0...)
}

// Warningf mocks base method
func (m *MockINovaLogger) Warningf(arg0 string, arg1 ...interface{}) {
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Warningf", varargs...)
}

// Warningf indicates an expected call of Warningf
func (mr *MockINovaLoggerMockRecorder) Warningf(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Warningf", reflect.TypeOf((*MockINovaLogger)(nil).Warningf), varargs...)
}

// WithError mocks base method
func (m *MockINovaLogger) WithError(arg0 error) *logrus.Entry {
	ret := m.ctrl.Call(m, "WithError", arg0)
	ret0, _ := ret[0].(*logrus.Entry)
	return ret0
}

// WithError indicates an expected call of WithError
func (mr *MockINovaLoggerMockRecorder) WithError(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithError", reflect.TypeOf((*MockINovaLogger)(nil).WithError), arg0)
}

// WithField mocks base method
func (m *MockINovaLogger) WithField(arg0 string, arg1 interface{}) *logrus.Entry {
	ret := m.ctrl.Call(m, "WithField", arg0, arg1)
	ret0, _ := ret[0].(*logrus.Entry)
	return ret0
}

// WithField indicates an expected call of WithField
func (mr *MockINovaLoggerMockRecorder) WithField(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithField", reflect.TypeOf((*MockINovaLogger)(nil).WithField), arg0, arg1)
}

// WithFields mocks base method
func (m *MockINovaLogger) WithFields(arg0 log4nova.Fields) *logrus.Entry {
	ret := m.ctrl.Call(m, "WithFields", arg0)
	ret0, _ := ret[0].(*logrus.Entry)
	return ret0
}

// WithFields indicates an expected call of WithFields
func (mr *MockINovaLoggerMockRecorder) WithFields(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithFields", reflect.TypeOf((*MockINovaLogger)(nil).WithFields), arg0)
}

// Write mocks base method
func (m *MockINovaLogger) Write(arg0 []byte) (int, error) {
	ret := m.ctrl.Call(m, "Write", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Write indicates an expected call of Write
func (mr *MockINovaLoggerMockRecorder) Write(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockINovaLogger)(nil).Write), arg0)
}
