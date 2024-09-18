package handler

type Handler struct {
	logger         Log
	devicesService DeviceService
}

func NewHandler(s DeviceService, logger Log) *Handler {
	return &Handler{
		devicesService: s,
		logger:         logger,
	}
}
