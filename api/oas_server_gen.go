// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"
)

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	// DevicesDeviceIDCommandsPost implements POST /devices/{device_id}/commands operation.
	//
	// Отправляет команду устройству (например "открыть
	// ворота").
	//
	// POST /devices/{device_id}/commands
	DevicesDeviceIDCommandsPost(ctx context.Context, req *DevicesDeviceIDCommandsPostReq, params DevicesDeviceIDCommandsPostParams) (DevicesDeviceIDCommandsPostRes, error)
	// DevicesDeviceIDGet implements GET /devices/{device_id} operation.
	//
	// Возвращает подробную информацию о конкретном
	// устройстве по его ID.
	//
	// GET /devices/{device_id}
	DevicesDeviceIDGet(ctx context.Context, params DevicesDeviceIDGetParams) (DevicesDeviceIDGetRes, error)
	// DevicesDeviceIDStatusPut implements PUT /devices/{device_id}/status operation.
	//
	// Позволяет изменить состояние устройства (например,
	// включить/выключить).
	//
	// PUT /devices/{device_id}/status
	DevicesDeviceIDStatusPut(ctx context.Context, req *DevicesDeviceIDStatusPutReq, params DevicesDeviceIDStatusPutParams) (DevicesDeviceIDStatusPutRes, error)
	// DevicesPost implements POST /devices operation.
	//
	// Добавляет новое устройство.
	//
	// POST /devices
	DevicesPost(ctx context.Context, req *DevicesPostReq) (DevicesPostRes, error)
}

// Server implements http server based on OpenAPI v3 specification and
// calls Handler to handle requests.
type Server struct {
	h Handler
	baseServer
}

// NewServer creates new Server.
func NewServer(h Handler, opts ...ServerOption) (*Server, error) {
	s, err := newServerConfig(opts...).baseServer()
	if err != nil {
		return nil, err
	}
	return &Server{
		h:          h,
		baseServer: s,
	}, nil
}
