// Code generated by MockGen. DO NOT EDIT.
// Source: contract.go
//
// Generated by this command:
//
//	mockgen -source=contract.go -destination=mock_contract.go -package=devices
//

// Package devices is a generated GoMock package.
package devices

import (
	context "context"
	entity "device-manager/internal/entity"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockStorage is a mock of Storage interface.
type MockStorage struct {
	ctrl     *gomock.Controller
	recorder *MockStorageMockRecorder
}

// MockStorageMockRecorder is the mock recorder for MockStorage.
type MockStorageMockRecorder struct {
	mock *MockStorage
}

// NewMockStorage creates a new mock instance.
func NewMockStorage(ctrl *gomock.Controller) *MockStorage {
	mock := &MockStorage{ctrl: ctrl}
	mock.recorder = &MockStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStorage) EXPECT() *MockStorageMockRecorder {
	return m.recorder
}

// AddDevice mocks base method.
func (m *MockStorage) AddDevice(ctx context.Context, deviceType, serialNumber, status string) (*entity.Device, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddDevice", ctx, deviceType, serialNumber, status)
	ret0, _ := ret[0].(*entity.Device)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddDevice indicates an expected call of AddDevice.
func (mr *MockStorageMockRecorder) AddDevice(ctx, deviceType, serialNumber, status any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddDevice", reflect.TypeOf((*MockStorage)(nil).AddDevice), ctx, deviceType, serialNumber, status)
}

// GetDevice mocks base method.
func (m *MockStorage) GetDevice(ctx context.Context, deviceID string) (*entity.Device, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDevice", ctx, deviceID)
	ret0, _ := ret[0].(*entity.Device)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDevice indicates an expected call of GetDevice.
func (mr *MockStorageMockRecorder) GetDevice(ctx, deviceID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDevice", reflect.TypeOf((*MockStorage)(nil).GetDevice), ctx, deviceID)
}

// LogCommandToDevice mocks base method.
func (m *MockStorage) LogCommandToDevice(ctx context.Context, deviceID, command, value string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LogCommandToDevice", ctx, deviceID, command, value)
	ret0, _ := ret[0].(error)
	return ret0
}

// LogCommandToDevice indicates an expected call of LogCommandToDevice.
func (mr *MockStorageMockRecorder) LogCommandToDevice(ctx, deviceID, command, value any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LogCommandToDevice", reflect.TypeOf((*MockStorage)(nil).LogCommandToDevice), ctx, deviceID, command, value)
}

// UpdateDeviceStatus mocks base method.
func (m *MockStorage) UpdateDeviceStatus(ctx context.Context, deviceID, status string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateDeviceStatus", ctx, deviceID, status)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateDeviceStatus indicates an expected call of UpdateDeviceStatus.
func (mr *MockStorageMockRecorder) UpdateDeviceStatus(ctx, deviceID, status any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateDeviceStatus", reflect.TypeOf((*MockStorage)(nil).UpdateDeviceStatus), ctx, deviceID, status)
}

// MockDatabusProducer is a mock of DatabusProducer interface.
type MockDatabusProducer struct {
	ctrl     *gomock.Controller
	recorder *MockDatabusProducerMockRecorder
}

// MockDatabusProducerMockRecorder is the mock recorder for MockDatabusProducer.
type MockDatabusProducerMockRecorder struct {
	mock *MockDatabusProducer
}

// NewMockDatabusProducer creates a new mock instance.
func NewMockDatabusProducer(ctrl *gomock.Controller) *MockDatabusProducer {
	mock := &MockDatabusProducer{ctrl: ctrl}
	mock.recorder = &MockDatabusProducerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDatabusProducer) EXPECT() *MockDatabusProducerMockRecorder {
	return m.recorder
}

// PublishDevice mocks base method.
func (m *MockDatabusProducer) PublishDevice(ctx context.Context, device *entity.Device) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PublishDevice", ctx, device)
	ret0, _ := ret[0].(error)
	return ret0
}

// PublishDevice indicates an expected call of PublishDevice.
func (mr *MockDatabusProducerMockRecorder) PublishDevice(ctx, device any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PublishDevice", reflect.TypeOf((*MockDatabusProducer)(nil).PublishDevice), ctx, device)
}
