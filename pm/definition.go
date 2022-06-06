package pm

import (
	"context"

	"github.com/go-sensors/core/units"
)

//go:generate mockgen -source=./definition.go -destination=./mocks/mock_definition.go -package=mocks

// Concentration is a measurement of particulate matter concentrations
type Concentration struct {
	// UpperBoundSize is the upper bound of the size of particles included in the concentration
	UpperBoundSize units.Distance
	// Amount is the mass concentration of particulate matter in a volume of air
	Amount units.MassConcentration
}

// ConcentrationSpec is a specification of valid measurement ranges supported by the sensor
type ConcentrationSpec struct {
	// UpperBoundSize is the upper bound of the size of particles included in the concentration
	UpperBoundSize units.Distance
	// Resolution is the smallest increment of measurement from the sensor
	Resolution units.MassConcentration
	// MinConcentration is the minimum valid measurement from the sensor
	MinConcentration units.MassConcentration
	// MaxConcentration is the maximum valid measurement from the sensor
	MaxConcentration units.MassConcentration
}

// ParticulateMatterSensor defines a sensor for measuring concentrations of particulate matter in air
type ParticulateMatterSensor interface {
	// Run begins reading from the sensor and blocks until either an error occurs or the context is completed
	Run(context.Context) error
	// Concentrations returns a channel of PM concentration readings as they become available from the sensor
	Concentrations() <-chan *Concentration
	// ConcentrationSpecs returns a collection of specified measurement ranges supported by the sensor
	ConcentrationSpecs() []*ConcentrationSpec
}
