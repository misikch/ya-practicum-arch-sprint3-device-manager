package devices

import (
	"context"
	"errors"

	"device-manager/internal/entity"
)

type Service struct {
	storage Storage
	databus DatabusProducer
}

var (
	ErrDeviceNotFound = errors.New("device not found")
)

func NewService(storage Storage, databus DatabusProducer) *Service {
	return &Service{
		storage: storage,
		databus: databus,
	}
}

func (s *Service) GetDeviceInfo(ctx context.Context, deviceID string) (entity.Device, error) {
	device, err := s.storage.GetDevice(ctx, deviceID)
	if err != nil {
		return entity.Device{}, err
	}

	if device == nil {
		return entity.Device{}, ErrDeviceNotFound
	}

	return *device, nil
}

func (s *Service) CreateDevice(ctx context.Context, deviceType string, serialNumber string, status string) (*entity.Device, error) {
	device, err := s.storage.AddDevice(ctx, deviceType, serialNumber, status)
	if err != nil {
		return device, err
	}

	// publish msg to databus
	err = s.databus.PublishDevice(ctx, device)
	if err != nil {
		return device, err
	}

	return device, nil
}

func (s *Service) RunCommand(ctx context.Context, deviceID string, command string, value string) error {
	device, err := s.storage.GetDevice(ctx, deviceID)
	if err != nil {
		return err
	}

	if device == nil {
		return ErrDeviceNotFound
	}

	//TODO: implement logic for device running command

	err = s.storage.LogCommandToDevice(ctx, deviceID, command, value)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) UpdateDeviceStatus(ctx context.Context, deviceID string, status string) error {
	device, err := s.storage.GetDevice(ctx, deviceID)
	if err != nil {
		return err
	}

	if device == nil {
		return ErrDeviceNotFound
	}

	err = s.storage.UpdateDeviceStatus(ctx, deviceID, status)
	if err != nil {
		return err
	}

	device.Status = status

	// publish msg to databus
	err = s.databus.PublishDevice(ctx, device)
	if err != nil {
		return err
	}

	return nil
}
