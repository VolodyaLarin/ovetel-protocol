package ovetel0

import "context"

// ClientInterface определяет методы, которые должен реализовать клиент.
type ClientInterface interface {
	SendVehicle(ctx context.Context, v *Vehicle) (*DefaultResp, error)
	SendTelemetry(ctx context.Context, s *SendTelemetryReq) (*DefaultResp, error)
	GetConfig(ctx context.Context) (*Config, error)
	GetChanges(ctx context.Context) (*Changes, error)
}
