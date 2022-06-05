package units

import (
	"fmt"
)

// MassConcentration represents the smallest measurable concentration of a ratio of masses in picograms per cubic meter
type MassConcentration Concentration

const (
	PicogramPerCubicMeter  = MassConcentration(1 * PartPerQuintillion)
	NanogramPerCubicMeter  = 1000 * PicogramPerCubicMeter
	MicrogramPerCubicMeter = 1000 * NanogramPerCubicMeter
	MilligramPerCubicMeter = 1000 * MicrogramPerCubicMeter
	GramPerCubicMeter      = 1000 * MilligramPerCubicMeter
)

func (concentration MassConcentration) String() string {
	if concentration >= GramPerCubicMeter {
		return fmt.Sprintf("%g g/㎥", concentration.GramsPerCubicMeter())
	}
	return fmt.Sprintf("%g µg/㎥", concentration.MicrogramsPerCubicMeter())
}

func (concentration MassConcentration) Concentration() Concentration {
	return Concentration(concentration)
}

func (concentration MassConcentration) PicogramsPerCubicMeter() float64 {
	return float64(concentration)
}

func (concentration MassConcentration) NanogramsPerCubicMeter() float64 {
	return float64(concentration) / float64(NanogramPerCubicMeter)
}

func (concentration MassConcentration) MicrogramsPerCubicMeter() float64 {
	return float64(concentration) / float64(MicrogramPerCubicMeter)
}

func (concentration MassConcentration) MilligramsPerCubicMeter() float64 {
	return float64(concentration) / float64(MilligramPerCubicMeter)
}

func (concentration MassConcentration) GramsPerCubicMeter() float64 {
	return float64(concentration) / float64(GramPerCubicMeter)
}
