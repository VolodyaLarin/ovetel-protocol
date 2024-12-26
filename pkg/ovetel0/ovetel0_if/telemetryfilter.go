package ovetel0_if

// TelemetryFilter описывает общие параметры фильтра телеметрии.
type TelemetryFilter struct {
	DeviceTypes []DeviceTypeEnum  `bson:"device_types"`
	Vendor      *string           `bson:"vendor,omitempty"`
	DataSource  *DataSourceEnum   `bson:"data_source,omitempty"`
	Level       int               `bson:"level"`
	Measurands  []MeasurandFilter `bson:"measurands,omitempty"`
}

// MeasurandFilter описывает фильтры для конкретного измеряемого параметра.
type MeasurandFilter struct {
	Measurand MeasurandEnum `bson:"measurand"`
	MinValue  *float64      `bson:"min_value,omitempty"`
	MaxValue  *float64      `bson:"max_value,omitempty"`
}
