package ovetel0_uc

import (
	"context"
	"github.com/VolodyaLarin/ovetel-protocol/pkg/ovetel0/ovetel0_if"
)

// IClientUsecase определяет методы, которые должен реализовать клиент.
type IClientUsecase interface {
	SendVehicle(ctx context.Context, v *ovetel0_if.Vehicle) (*ovetel0_if.DefaultResp, error)
	SendTelemetry(ctx context.Context, s *ovetel0_if.SendTelemetryReq) (*ovetel0_if.DefaultResp, error)
	GetConfig(ctx context.Context) (*ovetel0_if.Config, error)
	GetChanges(ctx context.Context) (*ovetel0_if.Changes, error)
}
