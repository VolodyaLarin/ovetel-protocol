package ovetel0

import "context"

// ServerInterface определяет методы, которые должен реализовать сервер.
type ServerInterface interface {
	HandleVehicle(context.Context, *Request[Vehicle]) (*DefaultResp, error)
	HandleSendTelemetry(context.Context, *Request[SendTelemetryReq]) (*DefaultResp, error)
	GetConfig(context.Context, *Request[DefaultReq]) (*Config, error)
	GetChanges(context.Context, *Request[DefaultReq]) (*Changes, error)
}
