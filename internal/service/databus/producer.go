package databus

import (
	"context"
	"encoding/json"

	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"

	"device-manager/internal/entity"
)

type Producer struct {
	writer *kafka.Writer
	logger *zap.SugaredLogger
}

type DeviceMessage struct {
	DeviceId string `json:"deviceId"`
	Status   string `json:"status"`
}

func NewProducer(writer *kafka.Writer, logger *zap.SugaredLogger) *Producer {
	return &Producer{
		writer: writer,
		logger: logger,
	}
}

// PublishDevice публикует событие по созданию/обновлению устройства в Kafka.
func (p *Producer) PublishDevice(ctx context.Context, device *entity.Device) error {
	message, err := json.Marshal(DeviceMessage{
		DeviceId: device.DeviceID,
		Status:   device.Status,
	})
	if err != nil {
		p.logger.Errorw("failed to marshal telemetry message", "error", err)
		return err
	}

	err = p.writer.WriteMessages(ctx, kafka.Message{
		Value: message,
	})
	if err != nil {
		p.logger.Errorw("failed to send deviceInfo message to kafka", "error", err)
		return err
	}

	p.logger.Infow("successfully sent deviceInfo message to kafka", "deviceId", device.DeviceID)
	return nil
}
