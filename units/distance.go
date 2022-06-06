package units

import "fmt"

// Distance represents the smallest measurable distance as an int64 nanometer count
type Distance uint64

const (
	Nanometer  Distance = 1
	Micrometer          = 1000 * Nanometer
	Millimeter          = 1000 * Micrometer
	Centimeter          = 10 * Millimeter
	Decimeter           = 10 * Centimeter
	Meter               = 10 * Decimeter
)

func (distance Distance) String() string {
	if distance >= Meter {
		return fmt.Sprintf("%g m", distance.Meters())
	}
	if distance >= Centimeter {
		return fmt.Sprintf("%g cm", distance.Centimeters())
	}
	if distance >= Millimeter {
		return fmt.Sprintf("%g mm", distance.Millimeters())
	}
	if distance >= Micrometer {
		return fmt.Sprintf("%g µm", distance.Micrometers())
	}
	return fmt.Sprintf("%d nm", distance.Nanometers())
}

func (distance Distance) Nanometers() int64 {
	return int64(distance)
}

func (distance Distance) Micrometers() float64 {
	return float64(distance) / float64(Micrometer)
}

func (distance Distance) Millimeters() float64 {
	return float64(distance) / float64(Millimeter)
}

func (distance Distance) Centimeters() float64 {
	return float64(distance) / float64(Centimeter)
}

func (distance Distance) Decimeters() float64 {
	return float64(distance) / float64(Decimeter)
}

func (distance Distance) Meters() float64 {
	return float64(distance) / float64(Meter)
}
