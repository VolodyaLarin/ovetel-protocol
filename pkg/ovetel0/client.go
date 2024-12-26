package ovetel0

import (
	"bytes"
	"context"
	"github.com/VolodyaLarin/ovetel-protocol/pkg/ovetel0/ovetel0_if"
	"github.com/pkg/errors"
	"gopkg.in/bson.v2"
	"io"
	"log/slog"
	"net/http"
	"net/url"
)

const BsonContentType = "application/bson"

type DefaultOvetel0Client struct {
	httpClient *http.Client
	baseUrl    url.URL
}

func (c *DefaultOvetel0Client) BaseUrl() url.URL {
	return c.baseUrl
}

func (c *DefaultOvetel0Client) HttpClient() *http.Client {
	return c.httpClient
}

func (c *DefaultOvetel0Client) SetHttpClient(httpClient *http.Client) {
	c.httpClient = httpClient
}

func (c *DefaultOvetel0Client) SetBaseUrl(baseUrl url.URL) {
	c.baseUrl = baseUrl
}

func clientSendData[TIN interface{}, TOUT interface{}](ctx context.Context, c *DefaultOvetel0Client, endpoint string, v *ovetel0_if.Request[TIN]) (*TOUT, error) {
	bsonData, err := bson.Marshal(v.Data)
	if err != nil {
		slog.Error("BSON marshal error", "err", errors.WithStack(err))
	}

	reqBody := bytes.NewBuffer(bsonData)

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseUrl.String()+endpoint, reqBody)
	if err != nil {
		return nil, errors.Wrap(errors.WithStack(err), "create request error")
	}
	req.Header.Set(ovetel0_if.AuthHeader, v.VehicleID)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(errors.WithStack(err), "do request error")
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			slog.Error("can't close body", "err", errors.WithStack(err))
			return
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(errors.WithStack(err), "read response body error")
	}

	var data TOUT

	err = bson.Unmarshal(body, &data)
	if err != nil {
		return nil, errors.Wrap(errors.WithStack(err), "can't parse body")
	}

	return &data, nil
}

func (c *DefaultOvetel0Client) SendVehicle(ctx context.Context, v *ovetel0_if.Request[ovetel0_if.Vehicle]) (*ovetel0_if.DefaultResp, error) {
	return clientSendData[ovetel0_if.Vehicle, ovetel0_if.DefaultResp](ctx, c, "vehicle", v)
}

func (c *DefaultOvetel0Client) SendTelemetry(ctx context.Context, s *ovetel0_if.Request[ovetel0_if.SendTelemetryReq]) (*ovetel0_if.DefaultResp, error) {
	return clientSendData[ovetel0_if.SendTelemetryReq, ovetel0_if.DefaultResp](ctx, c, "telemetry", s)
}

func (c *DefaultOvetel0Client) GetConfig(ctx context.Context, r *ovetel0_if.Request[ovetel0_if.DefaultReq]) (*ovetel0_if.Config, error) {
	return clientSendData[ovetel0_if.DefaultReq, ovetel0_if.Config](ctx, c, "config", r)
}

func (c *DefaultOvetel0Client) GetChanges(ctx context.Context, r *ovetel0_if.Request[ovetel0_if.DefaultReq]) (*ovetel0_if.Changes, error) {
	return clientSendData[ovetel0_if.DefaultReq, ovetel0_if.Changes](ctx, c, "changes", r)

}
