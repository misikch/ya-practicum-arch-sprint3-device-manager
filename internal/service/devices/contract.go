package devices

//go:generate mockgen -source=contract.go -destination=mock_contract.go -package=devices

import (
	"context"
	"device-manager/internal/entity"
)

type Storage interface {
	GetDevice(ctx context.Context, deviceID string) (*entity.Device, error)
	GetDeviceBySerial(ctx context.Context, serialNumber string) (*entity.Device, error)
	UpdateDeviceStatus(ctx context.Context, deviceID string, status string) error
	LogCommandToDevice(ctx context.Context, deviceID string, command string, value string) error
	AddDevice(ctx context.Context, deviceType string, serialNumber string, status string) (*entity.Device, error)
}

type DatabusProducer interface {
	PublishDevice(ctx context.Context, device *entity.Device) error
}
