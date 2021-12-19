package ta

import (
	"github.com/shopspring/decimal"
)

func NewGain() *GainLoss {
	return NewGainLoss(One)
}

func NewLoss() *GainLoss {
	return NewGainLoss(NegOne)
}

func NewGainLoss(coefficient decimal.Decimal) *GainLoss {
	return &GainLoss{
		coefficient: coefficient,
		firstOne:    true,
	}
}

var _ Indicator = (*GainLoss)(nil)

type GainLoss struct {
	coefficient decimal.Decimal
	lastValue   decimal.Decimal
	firstOne    bool
}

func (g *GainLoss) WindowSize() int {
	return 0
}

func (g *GainLoss) Update(value decimal.Decimal) decimal.Decimal {
	if g.firstOne {
		g.lastValue = value
		g.firstOne = false
		return Zero
	}
	delta := value.Sub(g.lastValue).Mul(g.coefficient)
	g.lastValue = value
	if delta.GreaterThan(Zero) {
		return delta
	}
	return Zero
}

func (g *GainLoss) DryUpdate(value decimal.Decimal) decimal.Decimal {
	delta := value.Sub(g.lastValue).Mul(g.coefficient)
	if delta.GreaterThan(Zero) {
		return delta
	}
	return Zero
}
