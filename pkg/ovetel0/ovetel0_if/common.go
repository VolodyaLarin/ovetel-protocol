package ovetel0_if

const AuthHeader = "x-vehicle-id"

// DefaultResp описывает стандартный ответ.
type DefaultResp struct {
	Status int     `bson:"status"`
	Error  *string `bson:"error,omitempty"`
}

// DefaultReq описывает стандартный запрос.
type DefaultReq struct{}

// Request описывает любой запрос
type Request[T interface{}] struct {
	VehicleID string
	Data      T
}
