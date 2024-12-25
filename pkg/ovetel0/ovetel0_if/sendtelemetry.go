package ovetel0_if

// SendTelemetryReq описывает запрос на отправку телеметрии.
type SendTelemetryReq struct {
	RequestID string      `bson:"request_id,omitempty"` // Внутренний уникальный идентификатор MongoDB
	Reason    ReasonEnum  `bson:"reason"`
	Telemetry []Telemetry `bson:"telemetry"`
}

// ReasonEnum представляет возможные причины отправки телеметрии.
type ReasonEnum string

const (
	Regular ReasonEnum = "Regular"
	Alert   ReasonEnum = "Alert"
	Manual  ReasonEnum = "Manual"
	Debug   ReasonEnum = "Debug"
)
