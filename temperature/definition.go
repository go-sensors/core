package temperature

import (
	"context"

	"github.com/go-sensors/core/units"
)

//go:generate mockgen -source=./definition.go -destination=./mocks/mock_definition.go -package=mocks

type TemperatureSpec struct {
	Resolution     units.Temperature
	MinTemperature units.Temperature
	MaxTemperature units.Temperature
}

type TemperatureSensor interface {
	// Run begins reading from the sensor and blocks until either an error occurs or the context is completed
	Run(context.Context) error
	// Temperatures returns a channel of temperature readings as they become available from the sensor
	Temperatures() <-chan *units.Temperature
	// TemperatureSpecs returns a collection of specified measurement ranges supported by the sensor
	TemperatureSpecs() []*TemperatureSpec
}
