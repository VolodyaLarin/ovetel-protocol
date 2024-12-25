package ovetel0_uc

import (
	"context"
	"github.com/VolodyaLarin/ovetel-protocol/pkg/ovetel0/ovetel0_if"
)

// IServerUsecase определяет методы, которые должен реализовать сервер.
type IServerUsecase interface {
	HandleVehicle(context.Context, *ovetel0_if.Request[ovetel0_if.Vehicle]) (*ovetel0_if.DefaultResp, error)
	HandleSendTelemetry(context.Context, *ovetel0_if.Request[ovetel0_if.SendTelemetryReq]) (*ovetel0_if.DefaultResp, error)
	GetConfig(context.Context, *ovetel0_if.Request[ovetel0_if.DefaultReq]) (*ovetel0_if.Config, error)
	GetChanges(context.Context, *ovetel0_if.Request[ovetel0_if.DefaultReq]) (*ovetel0_if.Changes, error)
}
