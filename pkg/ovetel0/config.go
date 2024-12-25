package ovetel0

import "net/url"

// Config описывает конфигурацию устройства.
type Config struct {
	Endpoint        url.URL           `bson:"endpoint"`
	TelemetryPeriod int               `bson:"telemetry_period"`
	Filters         []TelemetryFilter `bson:"filters"`
}
