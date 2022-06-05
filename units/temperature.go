package units

import "fmt"

// Temperature represents the smallest measurable temperature in billionths of a degree Celsius
type Temperature int64

const (
	BillionthDegreeCelsius  Temperature = 1
	MillionthDegreeCelsius  Temperature = 1000 * BillionthDegreeCelsius
	ThousandthDegreeCelsius Temperature = 1000 * MillionthDegreeCelsius
	DegreeCelsius           Temperature = 1000 * ThousandthDegreeCelsius
)

func (temperature Temperature) String() string {
	return fmt.Sprintf("%f °C", temperature.DegreesCelsius())
}

func (temperature Temperature) DegreesCelsius() float64 {
	return float64(temperature) / float64(DegreeCelsius)
}
