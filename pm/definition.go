package pm

import (
	"context"

	"github.com/go-sensors/core/units"
)

//go:generate mockgen -source=./definition.go -destination=./mocks/mock_definition.go -package=mocks

type Concentration struct {
	UpperBoundSize units.Distance
	Amount         units.MassConcentration
}

type ConcentrationSpec struct {
	UpperBoundSize   units.Distance
	Resolution       units.MassConcentration
	MinConcentration units.MassConcentration
	MaxConcentration units.MassConcentration
}

type ParticulateMatterSensor interface {
	// Run begins reading from the sensor and blocks until either an error occurs or the context is completed
	Run(context.Context) error
	// Concentrations returns a channel of PM concentration readings as they become available from the sensor
	Concentrations() <-chan *Concentration
	// ConcentrationSpecs returns a collection of specified measurement ranges supported by the sensor
	ConcentrationSpecs() []*ConcentrationSpec
}
