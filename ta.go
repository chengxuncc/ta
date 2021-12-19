package ta

import (
	"github.com/shopspring/decimal"
)

var (
	Zero   = decimal.Zero
	One    = decimal.NewFromInt(1)
	NegOne = decimal.NewFromInt(-1)
)

func NewDecimal(value string) decimal.Decimal {
	d, err := decimal.NewFromString(value)
	if err != nil {
		return decimal.Zero
	}
	return d
}

type Indicator interface {
	WindowSize() int
	Update(value decimal.Decimal) decimal.Decimal
	DryUpdate(value decimal.Decimal) decimal.Decimal
}

var _ Indicator = (Indicators)(nil)

type Indicators []Indicator

func (i Indicators) WindowSize() int {
	return 0
}

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

func Float64sToDecimals(values []float64) []decimal.Decimal {
	valuesLen := len(values)
	result := make([]decimal.Decimal, valuesLen)
	for i := 0; i < valuesLen; i++ {
		result[i] = decimal.NewFromFloat(values[i])
	}
	return result
}
