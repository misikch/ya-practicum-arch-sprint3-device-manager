package handler

import (
	"context"
	"device-manager/internal/service/devices"
	"errors"

	"device-manager/api"
)

func (h *Handler) DevicesDeviceIDGet(ctx context.Context, params api.DevicesDeviceIDGetParams) (api.DevicesDeviceIDGetRes, error) {
	device, err := h.devicesService.GetDeviceInfo(ctx, params.DeviceID)
	if err != nil {
		if errors.Is(err, devices.ErrDeviceNotFound) {
			return &api.DevicesDeviceIDGetNotFound{
				Code:    api.NewOptInt(404),
				Message: api.NewOptString("device with id not found"),
			}, nil
		}

		return &api.DevicesDeviceIDGetInternalServerError{
			Code:    api.NewOptInt(500),
			Message: api.NewOptString("internal server error"),
		}, nil
	}

	return &api.Device{
		DeviceID:     api.NewOptString(device.DeviceID),
		DeviceType:   api.NewOptString(device.DeviceType),
		SerialNumber: api.NewOptString(device.SerialNumber),
		Status:       api.NewOptString(device.Status),
	}, nil
}
