package handler

import (
	"context"
	"errors"

	"device-manager/api"
	"device-manager/internal/service/devices"
)

func (h *Handler) DevicesDeviceIDStatusPut(
	ctx context.Context,
	req *api.DevicesDeviceIDStatusPutReq,
	params api.DevicesDeviceIDStatusPutParams,
) (api.DevicesDeviceIDStatusPutRes, error) {
	err := h.devicesService.UpdateDeviceStatus(ctx, params.DeviceID, req.Status.Value)
	if err != nil {
		if errors.Is(err, devices.ErrDeviceNotFound) {
			return &api.DevicesDeviceIDStatusPutNotFound{
				Code:    api.NewOptInt(404),
				Message: api.NewOptString("device with id not found"),
			}, err
		}

		h.logger.Error(err)

		return &api.DevicesDeviceIDStatusPutInternalServerError{
			Code:    api.NewOptInt(500),
			Message: api.NewOptString("internal server error"),
		}, err
	}

	return &api.DevicesDeviceIDStatusPutOK{Status: req.Status}, nil
}
