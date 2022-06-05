package gas

import (
	"context"

	"github.com/go-sensors/core/units"
)

//go:generate mockgen -source=./definition.go -destination=./mocks/mock_definition.go -package=mocks

type Concentration struct {
	Gas    string
	Amount units.Concentration
}

type ConcentrationSpec struct {
	Gas              string
	Resolution       units.Concentration
	MinConcentration units.Concentration
	MaxConcentration units.Concentration
}

type GasSensor interface {
	// Run begins reading from the sensor and blocks until either an error occurs or the context is completed
	Run(context.Context) error
	// Concentrations returns a channel of gas concentration readings as they become available from the sensor
	Concentrations() <-chan *Concentration
	// ConcentrationSpecs returns a collection of specified measurement ranges supported by the sensor
	ConcentrationSpecs() []*ConcentrationSpec
}
