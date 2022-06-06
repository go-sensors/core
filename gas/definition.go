package gas

import (
	"context"

	"github.com/go-sensors/core/units"
)

//go:generate mockgen -source=./definition.go -destination=./mocks/mock_definition.go -package=mocks

// Concentration is a measurement of gas concentrations
type Concentration struct {
	// Gas is a descriptive string identifying the gas or group of gasses included in the concentration
	Gas string
	// Amount is the mass concentration of gas in a volume of air
	Amount units.Concentration
}

// ConcentrationSpec is a specification of valid measurement ranges supported by the sensor
type ConcentrationSpec struct {
	// Gas is a descriptive string identifying the gas or group of gasses included in the concentration
	Gas string
	// Resolution is the smallest increment of measurement from the sensor
	Resolution units.Concentration
	// MinConcentration is the minimum valid measurement from the sensor
	MinConcentration units.Concentration
	// MaxConcentration is the maximum valid measurement from the sensor
	MaxConcentration units.Concentration
}

// GasSensor defines a sensor for measuring concentrations of gas in air
type GasSensor interface {
	// Run begins reading from the sensor and blocks until either an error occurs or the context is completed
	Run(context.Context) error
	// Concentrations returns a channel of gas concentration readings as they become available from the sensor
	Concentrations() <-chan *Concentration
	// ConcentrationSpecs returns a collection of specified measurement ranges supported by the sensor
	ConcentrationSpecs() []*ConcentrationSpec
}
