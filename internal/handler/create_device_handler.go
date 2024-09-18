package handler

import (
	"context"

	"device-manager/api"
)

func (h *Handler) DevicesPost(ctx context.Context, req *api.DevicesPostReq) (api.DevicesPostRes, error) {
	device, err := h.devicesService.CreateDevice(ctx, req.DeviceType.Value, req.SerialNumber.Value, req.Status.Value)
	if err != nil {
		h.logger.Error(err)

		return &api.ErrorResponse{
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
