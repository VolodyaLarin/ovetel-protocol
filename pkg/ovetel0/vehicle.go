package ovetel0

// Vehicle описывает информацию о транспортном средстве.
type Vehicle struct {
	ID              string   `bson:"id"`
	VIN             string   `bson:"vin"`
	Model           *string  `bson:"model,omitempty"`
	Year            *int     `bson:"year,omitempty"`
	FirmwareVersion *string  `bson:"firmware_version,omitempty"`
	Devices         []Device `bson:"devices"`
}
