package main

import (
	"context"
	"errors"
	"github.com/VolodyaLarin/ovetel-protocol/pkg/ovetel0/ovetel0_if"
	"github.com/VolodyaLarin/ovetel-protocol/pkg/ovetel0/ovetel0_uc"
	"log/slog"
	"math/rand"
	"net/url"
	"sync"
	"time"

	"github.com/google/uuid"
	"golang.org/x/sync/semaphore"
)

var (
	ErrInvalidConfiguration = errors.New("invalid configuration received from server")
)

// Simulator представляет собой эмулятор телеметрии автомобиля.
type Simulator struct {
	client       ovetel0_uc.IClient
	vehicle      *ovetel0_if.Vehicle
	config       *ovetel0_if.Config
	telemetrySem *semaphore.Weighted
	stopCh       chan struct{}
	wg           sync.WaitGroup
}

// NewSimulator создает новый экземпляр эмулятора.
func NewSimulator(client ovetel0_uc.IClient, vehicle *ovetel0_if.Vehicle) *Simulator {
	return &Simulator{
		client:       client,
		vehicle:      vehicle,
		telemetrySem: semaphore.NewWeighted(500),
		stopCh:       make(chan struct{}),
	}
}

// Start запускает процесс эмуляции телеметрии.
func (e *Simulator) Start(ctx context.Context) error {
	e.wg.Add(1)
	go e.run(ctx)
	return nil
}

// Stop останавливает процесс эмуляции телеметрии.
func (e *Simulator) Stop() {
	close(e.stopCh)
	e.wg.Wait()
}

// run выполняет основной цикл эмуляции телеметрии.
func (e *Simulator) run(ctx context.Context) {
	defer e.wg.Done()

	slog.Info("Starts emulator")

	slog.Info("sendVehicleTopology")
	// Отправляем информацию о топологии автомобиля
	if err := e.sendVehicleTopology(ctx); err != nil {
		slog.Error("Error sendVehicleTopology", "error", err)
		return
	}

	slog.Info("get server configuration")

	// Получаем конфигурацию от сервера
	if err := e.getServerConfiguration(ctx); err != nil {
		slog.Error("Error get server configuration", "error", err)
		return
	}

	// Начинаем отправлять случайную телеметрию
	for {
		select {
		case <-time.After(time.Duration(e.config.TelemetryPeriod) * time.Second):
			slog.Info("sendRandomTelemetry")
			if err := e.sendRandomTelemetry(ctx); err != nil {
				slog.Error("Error sendRandomTelemetry", "error", err)
				return
			}
		case <-e.stopCh:
			return
		}
	}
}

// sendVehicleTopology отправляет информацию о топологии автомобиля на сервер.
func (e *Simulator) sendVehicleTopology(ctx context.Context) error {
	req := &ovetel0_if.Request[ovetel0_if.Vehicle]{
		VehicleID: e.vehicle.ID,
		Data:      *e.vehicle,
	}

	_, err := e.client.SendVehicle(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

// getServerConfiguration получает конфигурацию от сервера.
func (e *Simulator) getServerConfiguration(ctx context.Context) error {
	req := &ovetel0_if.Request[ovetel0_if.DefaultReq]{}

	config, err := e.client.GetConfig(ctx, req)
	if err != nil {
		return err
	}

	if config.Endpoint == (url.URL{}) || config.TelemetryPeriod <= 0 {
		return ErrInvalidConfiguration
	}

	e.config = config
	return nil
}

// sendRandomTelemetry отправляет случайную телеметрию на сервер.
func (e *Simulator) sendRandomTelemetry(ctx context.Context) error {
	if err := e.telemetrySem.Acquire(ctx, 1); err != nil {
		return err
	}
	defer e.telemetrySem.Release(1)

	// Генерируем случайную телеметрию
	telemetry := generateRandomTelemetry(e.vehicle)

	req := &ovetel0_if.Request[ovetel0_if.SendTelemetryReq]{
		VehicleID: e.vehicle.ID,
		Data: ovetel0_if.SendTelemetryReq{
			RequestID: uuid.New().String(),
			Reason:    ovetel0_if.Regular,
			Telemetry: telemetry,
		},
	}

	_, err := e.client.SendTelemetry(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

// generateRandomTelemetry генерирует случайную телеметрию для заданного автомобиля.
func generateRandomTelemetry(vehicle *ovetel0_if.Vehicle) []ovetel0_if.Telemetry {
	rand.Seed(time.Now().UnixNano())

	now := time.Now()
	startTime := now.Add(-time.Duration(rand.Intn(600)) * time.Second)
	endTime := startTime.Add(time.Duration(rand.Intn(300)+30) * time.Second)

	telemetry := make([]ovetel0_if.Telemetry, len(vehicle.Devices))
	for i, device := range vehicle.Devices {
		telemetry[i] = ovetel0_if.Telemetry{
			StartTimestamp:   startTime,
			EndTimestamp:     endTime,
			InternalDeviceID: device.InternalDeviceID,
			Measurand:        device.SupportedMeasurands[rand.Intn(len(device.SupportedMeasurands))],
			UOM:              UomEnums[rand.Intn(len(UomEnums))],
			Value:            rand.Float64(),
			DataType:         DataTypeEnums[rand.Intn(len(DataTypeEnums))],
			DataSource:       DataSourceEnums[rand.Intn(len(DataSourceEnums))],
			Level:            rand.Intn(3),
		}
	}

	return telemetry
}

// UomEnums содержит список возможных единиц измерения.
var UomEnums = []ovetel0_if.UomEnum{ovetel0_if.Celsius, ovetel0_if.Kilopascals, ovetel0_if.Percent}

// DataTypeEnums содержит список возможных способов представления данных.
var DataTypeEnums = []ovetel0_if.DataTypeEnum{ovetel0_if.Average, ovetel0_if.Last, ovetel0_if.Maximum}

// DataSourceEnums содержит список возможных источников данных.
var DataSourceEnums = []ovetel0_if.DataSourceEnum{ovetel0_if.Raw, ovetel0_if.Calculated}
