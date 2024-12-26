package ovetel0_uc

import (
	"context"
	"github.com/VolodyaLarin/ovetel-protocol/pkg/ovetel0/ovetel0_if"
	"net/http"
	"net/url"
)

// IClient определяет методы, которые должен реализовать клиент.
type IClient interface {
	SetBaseUrl(url url.URL)
	BaseUrl() (url url.URL)
	SetHttpClient(client *http.Client)
	HttpClient() (client *http.Client)

	SendVehicle(ctx context.Context, v *ovetel0_if.Request[ovetel0_if.Vehicle]) (*ovetel0_if.DefaultResp, error)
	SendTelemetry(ctx context.Context, s *ovetel0_if.Request[ovetel0_if.SendTelemetryReq]) (*ovetel0_if.DefaultResp, error)
	GetConfig(ctx context.Context, r *ovetel0_if.Request[ovetel0_if.DefaultReq]) (*ovetel0_if.Config, error)
	GetChanges(ctx context.Context, r *ovetel0_if.Request[ovetel0_if.DefaultReq]) (*ovetel0_if.Changes, error)
}
