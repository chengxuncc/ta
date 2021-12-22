package ta

import (
	"github.com/rs/zerolog/log"
	"github.com/shopspring/decimal"
)

var (
	Zero   = decimal.Zero
	One    = decimal.NewFromInt(1)
	NegOne = decimal.NewFromInt(-1)
)

func Decimal(s string) decimal.Decimal {
	d, err := decimal.NewFromString(s)
	if err != nil {
		log.Warn().Str("s", s).Msg("Decimal")
		return decimal.Zero
	}
	return d
}

type Indicator interface {
	Update(value decimal.Decimal) decimal.Decimal
	DryUpdate(value decimal.Decimal) decimal.Decimal
}

var _ Indicator = (Indicators)(nil)

type Indicators []Indicator

func (i Indicators) Update(value decimal.Decimal) decimal.Decimal {
	for _, indicator := range i {
		value = indicator.Update(value)
	}
	return value
}

func (i Indicators) DryUpdate(value decimal.Decimal) decimal.Decimal {
	for _, indicator := range i {
		value = indicator.DryUpdate(value)
	}
	return value
}

func BatchUpdate(indicator Indicator, values []decimal.Decimal) []decimal.Decimal {
	valuesLen := len(values)
	result := make([]decimal.Decimal, valuesLen)
	for i := 0; i < valuesLen; i++ {
		result[i] = indicator.Update(values[i])
	}
	return result
}

func BatchUpdateFloat64(indicator Indicator, values []float64) []float64 {
	valuesLen := len(values)
	result := make([]float64, valuesLen)
	for i := 0; i < valuesLen; i++ {
		result[i] = indicator.Update(decimal.NewFromFloat(values[i])).InexactFloat64()
	}
	return result
}
