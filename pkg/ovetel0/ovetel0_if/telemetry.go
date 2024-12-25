package ovetel0_if

import "time"

// Telemetry описывает структуру телеметрического измерения.
type Telemetry struct {
	StartTimestamp   time.Time      `bson:"start_timestamp"`
	EndTimestamp     time.Time      `bson:"end_timestamp"`
	InternalDeviceID string         `bson:"internal_device_id"`
	Measurand        MeasurandEnum  `bson:"measurand"`
	UOM              UomEnum        `bson:"uom"`
	Value            float64        `bson:"value"`
	DataType         DataTypeEnum   `bson:"data_type"`
	DataSource       DataSourceEnum `bson:"data_source"`
	Level            int            `bson:"level"`
}

// UomEnum представляет единицы измерения для параметра.
type UomEnum string

const (
	Celsius     UomEnum = "Celsius"
	Fahrenheit  UomEnum = "Fahrenheit"
	Kelvin      UomEnum = "Kelvin"
	Kilopascals UomEnum = "Kilopascals"
	Volts       UomEnum = "Volts"
	Amperes     UomEnum = "Amperes"
	Litres      UomEnum = "Litres"
	Percent     UomEnum = "Percent"
	MPS         UomEnum = "MPS"
	RPM         UomEnum = "RPM"
	Boolean     UomEnum = "Boolean"
)

// DataTypeEnum представляет способ представления данных.
type DataTypeEnum string

const (
	Average  DataTypeEnum = "Average"
	Minimum  DataTypeEnum = "Minimum"
	Maximum  DataTypeEnum = "Maximum"
	Sum      DataTypeEnum = "Sum"
	Variance DataTypeEnum = "Variance"
	Median   DataTypeEnum = "Median"
	Mode     DataTypeEnum = "Mode"
	Range    DataTypeEnum = "Range"
	Count    DataTypeEnum = "Count"
	First    DataTypeEnum = "First"
	Last     DataTypeEnum = "Last"
)

// DataSourceEnum представляет источник данных.
type DataSourceEnum string

const (
	Raw        DataSourceEnum = "Raw"
	Calculated DataSourceEnum = "Calculated"
)
