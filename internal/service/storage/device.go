package storage

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"device-manager/internal/entity"
)

// GetDevice получает информацию об устройстве по его device_id.
func (s *Storage) GetDevice(ctx context.Context, deviceID string) (*entity.Device, error) {
	collection := s.Database().Collection("devices")
	var device entity.Device
	err := collection.FindOne(ctx, bson.M{"deviceId": deviceID}).Decode(&device)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}

		s.logger.Errorf("Ошибка при получении устройства: %v", err)

		return nil, err
	}

	return &device, nil
}

// GetDeviceBySerial получает информацию об устройстве по его serial_number.
func (s *Storage) GetDeviceBySerial(ctx context.Context, serialNumber string) (*entity.Device, error) {
	collection := s.Database().Collection("devices")
	var device entity.Device
	err := collection.FindOne(ctx, bson.M{"serialNumber": serialNumber}).Decode(&device)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, nil
		}

		s.logger.Errorf("failed to get defice by serial %q: %v", serialNumber, err)

		return nil, err
	}

	return &device, nil
}

// UpdateDeviceStatus обновляет статус устройства.
func (s *Storage) UpdateDeviceStatus(ctx context.Context, deviceID string, status string) error {
	collection := s.Database().Collection("devices")
	_, err := collection.UpdateOne(
		ctx,
		bson.M{"deviceId": deviceID},
		bson.M{"$set": bson.M{"status": status}},
	)
	if err != nil {
		s.logger.Errorf("Ошибка при обновлении статуса устройства: %v", err)

		return err
	}

	return nil
}

// SendCommandToDevice отправляет команду устройству.
func (s *Storage) LogCommandToDevice(ctx context.Context, deviceID string, command string, value string) error {
	collection := s.Database().Collection("commands_log")
	_, err := collection.InsertOne(ctx, bson.M{
		"device_id": deviceID,
		"command":   command,
		"value":     value,
	})
	if err != nil {
		s.logger.Errorf("Ошибка при отправке команды устройству: %v", err)

		return err
	}

	return nil
}

// AddDevice добавляет новое устройство.
func (s *Storage) AddDevice(ctx context.Context, deviceType string, serialNumber string, status string) (*entity.Device, error) {
	collection := s.Database().Collection("devices")

	device := entity.Device{
		DeviceID:     uuid.New().String(), // Проставляем uuid
		DeviceType:   deviceType,
		SerialNumber: serialNumber,
		Status:       status,
	}

	_, err := collection.InsertOne(ctx, device)
	if err != nil {
		s.logger.Errorf("Ошибка при добавлении устройства: %v", err)

		return nil, err
	}

	return &device, nil
}
