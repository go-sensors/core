package units

import (
	"fmt"
)

// Pressure represents the smallest measurable pressure in billionths of a Pascal
type Pressure uint64

const (
	Nanopascal  Pressure = 1
	Micropascal Pressure = 1000 * Nanopascal
	Millipascal Pressure = 1000 * Micropascal
	Centipascal Pressure = 10 * Millipascal
	Decipascal  Pressure = 10 * Centipascal
	Pascal      Pressure = 10 * Decipascal
	Decapascal  Pressure = 10 * Pascal
	Hectopascal Pressure = 10 * Decapascal
	Kilopascal  Pressure = 10 * Hectopascal
)

func (pressure Pressure) String() string {
	if pressure >= Decapascal {
		return fmt.Sprintf("%g kPa", pressure.Kilopascals())
	}
	if pressure >= Pascal {
		return fmt.Sprintf("%g Pa", pressure.Pascals())
	}
	if pressure >= Decipascal {
		return fmt.Sprintf("%g dPa", pressure.Decipascals())
	}
	if pressure >= Centipascal {
		return fmt.Sprintf("%g cPa", pressure.Centipascals())
	}
	if pressure >= Millipascal {
		return fmt.Sprintf("%g mPa", pressure.Millipascals())
	}
	if pressure >= Micropascal {
		return fmt.Sprintf("%g µPa", pressure.Micropascals())
	}
	return fmt.Sprintf("%d nPa", pressure.Nanopascals())
}

func (distance Pressure) Nanopascals() int64 {
	return int64(distance)
}

func (distance Pressure) Micropascals() float64 {
	return float64(distance) / float64(Micropascal)
}

func (distance Pressure) Millipascals() float64 {
	return float64(distance) / float64(Millipascal)
}

func (distance Pressure) Centipascals() float64 {
	return float64(distance) / float64(Centipascal)
}

func (distance Pressure) Decipascals() float64 {
	return float64(distance) / float64(Decipascal)
}

func (distance Pressure) Pascals() float64 {
	return float64(distance) / float64(Pascal)
}

func (distance Pressure) Hectopascals() float64 {
	return float64(distance) / float64(Hectopascal)
}

func (distance Pressure) Kilopascals() float64 {
	return float64(distance) / float64(Kilopascal)
}
