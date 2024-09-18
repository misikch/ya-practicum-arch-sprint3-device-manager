package storage

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"device-manager/internal/entity"
)

// GetDevice получает информацию об устройстве по его device_id.
func (s *Storage) GetDevice(ctx context.Context, deviceID string) (*entity.Device, error) {
	collection := s.Database().Collection("devices")
	var device entity.Device
	err := collection.FindOne(ctx, bson.M{"device_id": deviceID}).Decode(&device)
	if err != nil {
		s.logger.Errorf("Ошибка при получении устройства: %v", err)

		return nil, err
	}

	return &device, nil
}

// UpdateDeviceStatus обновляет статус устройства.
func (s *Storage) UpdateDeviceStatus(ctx context.Context, deviceID string, status string) error {
	collection := s.Database().Collection("devices")
	_, err := collection.UpdateOne(
		ctx,
		bson.M{"device_id": deviceID},
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
		DeviceType:   deviceType,
		SerialNumber: serialNumber,
		Status:       status,
	}
	res, err := collection.InsertOne(ctx, device)
	if err != nil {
		s.logger.Errorf("Ошибка при добавлении устройства: %v", err)

		return nil, err
	}

	device.DeviceID = res.InsertedID.(primitive.ObjectID).Hex()

	return &device, nil
}
