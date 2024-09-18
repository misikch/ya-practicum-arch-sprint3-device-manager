package handler

import (
	"context"
	"errors"

	"device-manager/api"
	"device-manager/internal/service/devices"
)

func (h *Handler) DevicesDeviceIDCommandsPost(
	ctx context.Context,
	req *api.DevicesDeviceIDCommandsPostReq,
	params api.DevicesDeviceIDCommandsPostParams,
) (api.DevicesDeviceIDCommandsPostRes, error) {
	err := h.devicesService.RunCommand(ctx, params.DeviceID, req.Command.Value, req.Command.Value)
	if err != nil {
		if errors.Is(err, devices.ErrDeviceNotFound) {
			return &api.DevicesDeviceIDCommandsPostNotFound{
				Code:    api.NewOptInt(404),
				Message: api.NewOptString("device with id not found"),
			}, nil
		}

		h.logger.Error(err)

		return &api.DevicesDeviceIDCommandsPostInternalServerError{
			Code:    api.NewOptInt(500),
			Message: api.NewOptString("internal server error"),
		}, nil
	}

	return &api.DevicesDeviceIDCommandsPostOK{Result: api.NewOptString("ok")}, nil
}
