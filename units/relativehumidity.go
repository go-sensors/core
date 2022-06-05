package units

import (
	"fmt"
	"math"
)

// RelativeHumidity represents the amount of water vapor present in air expressed as a percentage of the amount needed for saturation at the given temperature
type RelativeHumidity struct {
	Temperature Temperature
	// Percentage is a value typically in the range of 0.0 (0%) to 1.0 (100%), though higher values may be feasible depending on conditions
	Percentage float64
}

func (relativeHumidity RelativeHumidity) String() string {
	rh := relativeHumidity.Percentage * 100
	return fmt.Sprintf("%2.2f %%RH @ %v", rh, relativeHumidity.Temperature.String())
}

// AbsoluteHumidity derives the absolute humidity (concentration of mass of humidity in air) from the relative humidity
func (relativeHumidity RelativeHumidity) AbsoluteHumidity() MassConcentration {
	// Adapted from https://github.com/skgrange/threadr/blob/fd42380883133fe7a47c479e778afe644a507334/R/absolute_humidity.R
	rh := relativeHumidity.Percentage * 100
	t := relativeHumidity.Temperature.DegreesCelsius()
	h := (6.112 * math.Exp((17.67*t)/(t+243.5)) * rh * 2.1674) / (273.15 + t)

	humidity := MassConcentration(h * float64(GramPerCubicMeter))
	return humidity
}
