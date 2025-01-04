package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"github.com/VolodyaLarin/ovetel-protocol/pkg/ovetel0"
	"github.com/VolodyaLarin/ovetel-protocol/pkg/ovetel0/ovetel0_if"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/google/uuid"
)

func retPtr[T any](d T) *T {
	return &d
}

func main() {
	certPool := x509.NewCertPool()
	pemData, err := ioutil.ReadFile("./examples/cert/cert.pem")
	if err != nil {
		log.Fatalf("Failed to read certificate file: %v", err)
	}
	ok := certPool.AppendCertsFromPEM(pemData)
	if !ok {
		log.Fatal("Failed to append certificates from PEM data")
	}

	// Создаем TLS-конфигурацию, которая доверяет вашему сертификату
	tlsConfig := &tls.Config{
		RootCAs:            certPool,
		InsecureSkipVerify: true,
	}

	// Создаем транспорт с настроенной TLS-конфигурацией
	transport := &http.Transport{TLSClientConfig: tlsConfig}

	// Создаем HTTP-клиент с вашим транспортом
	httpClient := &http.Client{Transport: transport}

	// Создаем клиент для взаимодействия с сервером
	baseURL, _ := url.Parse("http://127.0.0.1:8443/")
	client := ovetel0.NewDefaultOvetel0Client(*baseURL)
	client.SetHttpClient(httpClient)

	// Создаем информацию о топологии автомобиля
	vehicle := &ovetel0_if.Vehicle{
		ID:              uuid.New().String(),
		VIN:             "12345678901234567",
		Model:           retPtr("Toyota Corolla"),
		Year:            retPtr(2023),
		FirmwareVersion: retPtr("v1.0.0"),
		Devices: []ovetel0_if.Device{
			{
				InternalDeviceID: "ff00::device_1",
				DeviceType:       ovetel0_if.GPSModuleDeviceType,
				Vendor:           "VendorX",
				FirmwareVersion:  retPtr("v2.1.4"),
				SupportedMeasurands: []ovetel0_if.MeasurandEnum{
					ovetel0_if.MeasurandEnum("Speed"),
					ovetel0_if.MeasurandEnum("Location"),
				},
			},
		},
	}

	// Создаем и запускаем эмулятор
	emulator := NewEmulator(client, vehicle)
	if err := emulator.Start(context.Background()); err != nil {
		log.Fatalf("Failed to start emulator: %v", err)
	}

	// Ожидаем завершения эмулятора
	time.Sleep(5 * time.Minute)
	emulator.Stop()
}
