package humidity

import (
	"context"

	"github.com/go-sensors/core/units"
)

//go:generate mockgen -source=./definition.go -destination=./mocks/mock_definition.go -package=mocks

type RelativeHumiditySpec struct {
	// PercentageResolution is a value typically in the range of 0.0 (0%) to 1.0 (100%), though higher values may be feasible depending on conditions
	PercentageResolution float64
	// MinPercentage is a value typically in the range of 0.0 (0%) to 1.0 (100%), though higher values may be feasible depending on conditions
	MinPercentage float64
	// MaxPercentage is a value typically in the range of 0.0 (0%) to 1.0 (100%), though higher values may be feasible depending on conditions
	MaxPercentage float64
}

type RelativeHumiditySensor interface {
	// Run begins reading from the sensor and blocks until either an error occurs or the context is completed
	Run(context.Context) error
	// RelativeHumidities returns a channel of relative humidity readings as they become available from the sensor
	RelativeHumidities() <-chan *units.RelativeHumidity
	// HumiditySpecs returns a collection of specified measurement ranges supported by the sensor
	RelativeHumiditySpecs() []*RelativeHumiditySpec
}
