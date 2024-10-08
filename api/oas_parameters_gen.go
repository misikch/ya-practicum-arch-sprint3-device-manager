// Code generated by ogen, DO NOT EDIT.

package api

import (
	"net/http"
	"net/url"

	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen/conv"
	"github.com/ogen-go/ogen/middleware"
	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
)

// DevicesDeviceIDCommandsPostParams is parameters of POST /devices/{device_id}/commands operation.
type DevicesDeviceIDCommandsPostParams struct {
	DeviceID string
}

func unpackDevicesDeviceIDCommandsPostParams(packed middleware.Parameters) (params DevicesDeviceIDCommandsPostParams) {
	{
		key := middleware.ParameterKey{
			Name: "device_id",
			In:   "path",
		}
		params.DeviceID = packed[key].(string)
	}
	return params
}

func decodeDevicesDeviceIDCommandsPostParams(args [1]string, argsEscaped bool, r *http.Request) (params DevicesDeviceIDCommandsPostParams, _ error) {
	// Decode path: device_id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "device_id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.DeviceID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "device_id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// DevicesDeviceIDGetParams is parameters of GET /devices/{device_id} operation.
type DevicesDeviceIDGetParams struct {
	DeviceID string
}

func unpackDevicesDeviceIDGetParams(packed middleware.Parameters) (params DevicesDeviceIDGetParams) {
	{
		key := middleware.ParameterKey{
			Name: "device_id",
			In:   "path",
		}
		params.DeviceID = packed[key].(string)
	}
	return params
}

func decodeDevicesDeviceIDGetParams(args [1]string, argsEscaped bool, r *http.Request) (params DevicesDeviceIDGetParams, _ error) {
	// Decode path: device_id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "device_id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.DeviceID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "device_id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// DevicesDeviceIDStatusPutParams is parameters of PUT /devices/{device_id}/status operation.
type DevicesDeviceIDStatusPutParams struct {
	DeviceID string
}

func unpackDevicesDeviceIDStatusPutParams(packed middleware.Parameters) (params DevicesDeviceIDStatusPutParams) {
	{
		key := middleware.ParameterKey{
			Name: "device_id",
			In:   "path",
		}
		params.DeviceID = packed[key].(string)
	}
	return params
}

func decodeDevicesDeviceIDStatusPutParams(args [1]string, argsEscaped bool, r *http.Request) (params DevicesDeviceIDStatusPutParams, _ error) {
	// Decode path: device_id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "device_id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.DeviceID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "device_id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}
