package handler

import (
	"context"

	"device-manager/internal/entity"
)

type DeviceService interface {
	GetDeviceInfo(ctx context.Context, deviceID string) (entity.Device, error)
	CreateDevice(ctx context.Context, deviceType string, serialNumber string, status string) (*entity.Device, error)
	RunCommand(ctx context.Context, deviceID string, command string, value string) error
	UpdateDeviceStatus(ctx context.Context, deviceID string, status string) error
}

type Log interface {
	Error(args ...interface{})
}
