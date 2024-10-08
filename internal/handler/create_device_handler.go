package handler

import (
	"context"
	"errors"

	"device-manager/api"
	"device-manager/internal/service/devices"
)

func (h *Handler) DevicesPost(ctx context.Context, req *api.DevicesPostReq) (api.DevicesPostRes, error) {
	device, err := h.devicesService.CreateDevice(ctx, req.DeviceType.Value, req.SerialNumber.Value, req.Status.Value)
	if err != nil {
		if errors.Is(err, devices.ErrDeviceWithSerialAlreadyExists) {
			return &api.DevicesPostConflict{
				Code:    api.NewOptInt(409),
				Message: api.NewOptString("device with serial number already exists"),
			}, nil
		}

		h.logger.Error(err)

		return &api.DevicesPostInternalServerError{
			Code:    api.NewOptInt(500),
			Message: api.NewOptString("internal server error"),
		}, nil
	}

	return &api.Device{
		DeviceId:     api.NewOptString(device.DeviceID),
		DeviceType:   api.NewOptString(device.DeviceType),
		SerialNumber: api.NewOptString(device.SerialNumber),
		Status:       api.NewOptString(device.Status),
	}, nil
}
