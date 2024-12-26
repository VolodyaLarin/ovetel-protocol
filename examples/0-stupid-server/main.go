package main

import (
	"context"
	"crypto/tls"
	"github.com/VolodyaLarin/ovetel-protocol/pkg/ovetel0"
	"github.com/VolodyaLarin/ovetel-protocol/pkg/ovetel0/ovetel0_if"
	"log/slog"
	"net/http"
	"net/url"
)

// MockUsecase реализует интерфейс IServerUsecase и возвращает фиктивные данные
type MockUsecase struct{}

func (m *MockUsecase) HandleVehicle(ctx context.Context, o *ovetel0_if.Request[ovetel0_if.Vehicle]) (*ovetel0_if.DefaultResp, error) {
	return &ovetel0_if.DefaultResp{}, nil
}

func (m *MockUsecase) HandleSendTelemetry(ctx context.Context, o *ovetel0_if.Request[ovetel0_if.SendTelemetryReq]) (*ovetel0_if.DefaultResp, error) {
	return &ovetel0_if.DefaultResp{}, nil
}

var vendor = "x-bmstu"

func (m *MockUsecase) GetConfig(ctx context.Context, o *ovetel0_if.Request[ovetel0_if.DefaultReq]) (*ovetel0_if.Config, error) {
	baseUrl, _ := url.Parse("https://localhost")
	return &ovetel0_if.Config{
		Endpoint:        *baseUrl,
		TelemetryPeriod: 30,
		Filters: []ovetel0_if.TelemetryFilter{{
			DeviceTypes: []ovetel0_if.DeviceTypeEnum{ovetel0_if.EngineLoadDeviceType},
			Vendor:      &vendor,
			DataSource:  nil,
			Level:       0,
			Measurands: []ovetel0_if.MeasurandFilter{
				{
					Measurand: "hel",
				},
			},
		}},
	}, nil
}

func (m *MockUsecase) GetChanges(ctx context.Context, o *ovetel0_if.Request[ovetel0_if.DefaultReq]) (*ovetel0_if.Changes, error) {
	return &ovetel0_if.Changes{
		Requests: []ovetel0_if.TelemetryRequest{},
	}, nil
}

func main() {
	mockUsecase := &MockUsecase{}
	server := ovetel0.NewServer(mockUsecase)

	mux := http.NewServeMux()
	server.RegisterRoutes(mux)

	certFile := "./examples/cert/cert.pem" // Путь к вашему сертификату
	keyFile := "./examples/cert/key.pem"   // Путь к вашему приватному ключу
	srv := &http.Server{
		Addr:    ":8443",
		Handler: mux,
		TLSConfig: &tls.Config{
			
			MinVersion:       tls.VersionTLS12,
			CurvePreferences: []tls.CurveID{tls.X25519, tls.CurveP256},
		},
	}

	slog.Info("Starting HTTP/2 server", "url", srv.Addr)

	err := srv.ListenAndServeTLS(certFile, keyFile)
	if err != nil {
		slog.Error("error in server creation", "error", err)
	}
}
