package ovetel0

// Device описывает устройство, установленное на транспортное средство.
type Device struct {
	InternalDeviceID    string          `bson:"internal_device_id"`
	DeviceType          DeviceTypeEnum  `bson:"device_type"`
	VendorDeviceType    *string         `bson:"vendor_device_type,omitempty"`
	Vendor              string          `bson:"vendor"`
	FirmwareVersion     *string         `bson:"firmware_version,omitempty"`
	SupportedMeasurands []MeasurandEnum `bson:"supported_measurands,omitempty"`
}

// DeviceTypeEnum представляет возможные типы устройств.
type DeviceTypeEnum string

const (
	TemperatureSensorDeviceType DeviceTypeEnum = "TemperatureSensor"
	PressureSensorDeviceType    DeviceTypeEnum = "PressureSensor"
	HumiditySensorDeviceType    DeviceTypeEnum = "HumiditySensor"
	GPSModuleDeviceType         DeviceTypeEnum = "GPSModule"
	RPMMeterDeviceType          DeviceTypeEnum = "RPMMeter"
	FuelLevelSensorDeviceType   DeviceTypeEnum = "FuelLevelSensor"
	BatteryVoltageDeviceType    DeviceTypeEnum = "BatteryVoltage"
	OilPressureDeviceType       DeviceTypeEnum = "OilPressure"
	CoolantTempDeviceType       DeviceTypeEnum = "CoolantTemp"
	EngineLoadDeviceType        DeviceTypeEnum = "EngineLoad"
	AirFlowRateDeviceType       DeviceTypeEnum = "AirFlowRate"
	ThrottlePositionDeviceType  DeviceTypeEnum = "ThrottlePosition"
)

// MeasurandEnum представляет возможные измеряемые параметры.
type MeasurandEnum string
