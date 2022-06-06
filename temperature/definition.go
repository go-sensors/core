package temperature

import (
	"context"

	"github.com/go-sensors/core/units"
)

//go:generate mockgen -source=./definition.go -destination=./mocks/mock_definition.go -package=mocks

// TemperatureSpec is a specification of valid measurement ranges supported by the sensor
type TemperatureSpec struct {
	// Resolution is the smallest increment of measurement from the sensor
	Resolution units.Temperature
	// MinTemperature is the minimum valid measurement from the sensor
	MinTemperature units.Temperature
	// MaxTemperature is the maximum valid measurement from the sensor
	MaxTemperature units.Temperature
}

// TemperatureSensor defines a sensor for measuring temperature
type TemperatureSensor interface {
	// Run begins reading from the sensor and blocks until either an error occurs or the context is completed
	Run(context.Context) error
	// Temperatures returns a channel of temperature readings as they become available from the sensor
	Temperatures() <-chan *units.Temperature
	// TemperatureSpecs returns a collection of specified measurement ranges supported by the sensor
	TemperatureSpecs() []*TemperatureSpec
}
