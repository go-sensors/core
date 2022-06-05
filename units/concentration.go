package units

import (
	"fmt"
)

// Concentration represents the smallest measurable concentration in parts per quintillion
type Concentration uint64

const (
	PartPerQuintillion Concentration = 1
	PartPerQuadrillion               = 1000 * PartPerQuintillion
	PartPerTrillion                  = 1000 * PartPerQuadrillion
	PartPerBillion                   = 1000 * PartPerTrillion
	PartPerMillion                   = 1000 * PartPerBillion
	PartPerThousand                  = 1000 * PartPerMillion
)

func (concentration Concentration) String() string {
	if concentration >= PartPerThousand {
		return fmt.Sprintf("%g parts per thousand", concentration.PartsPerThousand())
	}
	if concentration >= PartPerMillion {
		return fmt.Sprintf("%g parts per million", concentration.PartsPerMillion())
	}
	if concentration >= PartPerBillion {
		return fmt.Sprintf("%g parts per billion", concentration.PartsPerBillion())
	}
	if concentration >= PartPerTrillion {
		return fmt.Sprintf("%g parts per trillion", concentration.PartsPerTrillion())
	}
	if concentration >= PartPerQuadrillion {
		return fmt.Sprintf("%g parts per quadrillion", concentration.PartsPerQuadrillion())
	}
	return fmt.Sprintf("%g parts per quintillion", concentration.PartsPerQuintillion())
}

func (concentration Concentration) PartsPerQuintillion() float64 {
	return float64(concentration)
}

func (concentration Concentration) PartsPerQuadrillion() float64 {
	return float64(concentration) / float64(PartPerQuadrillion)
}

func (concentration Concentration) PartsPerTrillion() float64 {
	return float64(concentration) / float64(PartPerTrillion)
}

func (concentration Concentration) PartsPerBillion() float64 {
	return float64(concentration) / float64(PartPerBillion)
}

func (concentration Concentration) PartsPerMillion() float64 {
	return float64(concentration) / float64(PartPerMillion)
}

func (concentration Concentration) PartsPerThousand() float64 {
	return float64(concentration) / float64(PartPerThousand)
}
