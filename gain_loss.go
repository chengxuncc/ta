package ta

import (
	"github.com/shopspring/decimal"
)

func NewGainLoss(coefficient decimal.Decimal) *GainLoss {
	return &GainLoss{
		coefficient: coefficient,
	}
}

func NewGain() *GainLoss {
	return NewGainLoss(One)
}

func NewLoss() *GainLoss {
	return NewGainLoss(NegOne)
}

var _ Indicator = (*GainLoss)(nil)

type GainLoss struct {
	coefficient decimal.Decimal
	lastValue   decimal.Decimal
}

func (g *GainLoss) WindowSize() int {
	return 0
}

func (g *GainLoss) Update(value decimal.Decimal) decimal.Decimal {
	delta := value.Sub(g.lastValue).Mul(g.coefficient)
	g.lastValue = value
	if delta.GreaterThan(decimal.Zero) {
		return delta
	}
	return decimal.Zero
}

func (g *GainLoss) DryUpdate(value decimal.Decimal) decimal.Decimal {
	delta := value.Sub(g.lastValue).Mul(g.coefficient)
	if delta.GreaterThan(decimal.Zero) {
		return delta
	}
	return decimal.Zero
}
