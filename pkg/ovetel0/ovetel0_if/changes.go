package ovetel0_if

import (
	"time"
)

// Changes
type Changes struct {
	Requests []TelemetryRequest `bson:"requests"`
}

// TelemetryRequest описывает запрос на уточнение телеметрии.
type TelemetryRequest struct {
	RequestId      string            `bson:"request_id"`
	StartTimestamp time.Time         `bson:"start_timestamp"`
	EndTimestamp   time.Time         `bson:"end_timestamp"`
	Filters        []TelemetryFilter `bson:"filters"`
}
