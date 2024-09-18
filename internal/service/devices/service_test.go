package devices

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"

	"device-manager/internal/entity"
)

type TestCase struct {
	name    string
	setup   func(*MockStorage, *MockDatabusProducer)
	execute func(*Service) error
	check   func(*testing.T, error)
}

func TestService_GetDeviceInfo(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := NewMockStorage(ctrl)
	mockDatabus := NewMockDatabusProducer(ctrl)
	service := NewService(mockStorage, mockDatabus)

	tests := []TestCase{
		{
			name: "Device found",
			setup: func(storage *MockStorage, _ *MockDatabusProducer) {
				storage.EXPECT().GetDevice(gomock.Any(), "1").Return(&entity.Device{DeviceID: "1"}, nil)
			},
			execute: func(s *Service) error {
				_, err := s.GetDeviceInfo(context.Background(), "1")
				return err
			},
			check: func(t *testing.T, err error) {
				assert.NoError(t, err)
			},
		},
		{
			name: "Device not found",
			setup: func(storage *MockStorage, _ *MockDatabusProducer) {
				storage.EXPECT().GetDevice(gomock.Any(), "1").Return(nil, nil)
			},
			execute: func(s *Service) error {
				_, err := s.GetDeviceInfo(context.Background(), "1")
				return err
			},
			check: func(t *testing.T, err error) {
				assert.Error(t, err)
				assert.Equal(t, ErrDeviceNotFound, err)
			},
		},
		{
			name: "Storage error",
			setup: func(storage *MockStorage, _ *MockDatabusProducer) {
				storage.EXPECT().GetDevice(gomock.Any(), "1").Return(nil, errors.New("some error"))
			},
			execute: func(s *Service) error {
				_, err := s.GetDeviceInfo(context.Background(), "1")
				return err
			},
			check: func(t *testing.T, err error) {
				assert.Error(t, err)
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.setup(mockStorage, mockDatabus)
			err := tc.execute(service)
			tc.check(t, err)
		})
	}
}

func TestService_CreateDevice(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStorage := NewMockStorage(ctrl)
	mockDatabus := NewMockDatabusProducer(ctrl)
	service := NewService(mockStorage, mockDatabus)

	tests := []TestCase{
		{
			name: "Device created and published",
			setup: func(storage *MockStorage, databus *MockDatabusProducer) {
				storage.EXPECT().AddDevice(gomock.Any(), "type", "serial", "status").Return(&entity.Device{DeviceID: "1"}, nil)
				databus.EXPECT().PublishDevice(gomock.Any(), &entity.Device{DeviceID: "1"}).Return(nil)
			},
			execute: func(s *Service) error {
				_, err := s.CreateDevice(context.Background(), "type", "serial", "status")
				return err
			},
			check: func(t *testing.T, err error) {
				assert.NoError(t, err)
			},
		},
		{
			name: "Storage returns error",
			setup: func(storage *MockStorage, databus *MockDatabusProducer) {
				storage.EXPECT().AddDevice(gomock.Any(), "type", "serial", "status").Return(nil, errors.New("some error"))
			},
			execute: func(s *Service) error {
				_, err := s.CreateDevice(context.Background(), "type", "serial", "status")
				return err
			},
			check: func(t *testing.T, err error) {
				assert.Error(t, err)
			},
		},
		{
			name: "Publish returns error",
			setup: func(storage *MockStorage, databus *MockDatabusProducer) {
				storage.EXPECT().AddDevice(gomock.Any(), "type", "serial", "status").Return(&entity.Device{DeviceID: "1"}, nil)
				databus.EXPECT().PublishDevice(gomock.Any(), &entity.Device{DeviceID: "1"}).Return(errors.New("publish error"))
			},
			execute: func(s *Service) error {
				_, err := s.CreateDevice(context.Background(), "type", "serial", "status")
				return err
			},
			check: func(t *testing.T, err error) {
				assert.Error(t, err)
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tc.setup(mockStorage, mockDatabus)
			err := tc.execute(service)
			tc.check(t, err)
		})
	}
}
