package ta

import (
	"github.com/shopspring/decimal"
)

// Exponential moving averages (EMAs) reduce the lag by
// applying more weight to recent prices. The weighting
// applied to the most recent price depends on the number
// of periods in the moving average. EMAs differ from
// simple moving averages in that a given day's EMA
// calculation depends on the EMA calculations for all
// the days prior to that day. You need far more than 10
// days of data to calculate a reasonably accurate 10-day EMA.
//  https://school.stockcharts.com/doku.php?id=technical_indicators:moving_averages

func NewEMA(window int) *ExponentialMovingAverage {
	return NewEMAWithSmoothing(window, decimal.NewFromInt(2))
}

func NewEMAWithSmoothing(window int, smoothing decimal.Decimal) *ExponentialMovingAverage {
	k1 := smoothing.Div(decimal.NewFromInt(int64(window + 1)))
	k2 := One.Sub(k1)
	return &ExponentialMovingAverage{
		window: window,
		sma:    NewSMA(window),
		k1:     k1,
		k2:     k2,
	}
}

var _ Indicator = (*ExponentialMovingAverage)(nil)

type ExponentialMovingAverage struct {
	window int
	sma    *SimpleMovingAverage
	k1     decimal.Decimal
	k2     decimal.Decimal
	ema    decimal.Decimal
	count  int
}

func (e *ExponentialMovingAverage) WindowSize() int {
	return e.window
}

func (e *ExponentialMovingAverage) Update(value decimal.Decimal) decimal.Decimal {
	e.count++
	switch {
	case e.count <= e.window:
		e.ema = e.sma.Update(value)
	case e.count > e.window:
		e.ema = value.Mul(e.k1).Add(e.ema.Mul(e.k2))
	}
	return e.ema
}

func (e *ExponentialMovingAverage) DryUpdate(value decimal.Decimal) decimal.Decimal {
	var ema decimal.Decimal
	switch {
	case e.count <= e.window:
		ema = e.sma.DryUpdate(value)
	case e.count > e.window:
		ema = value.Mul(e.k1).Add(e.ema.Mul(e.k2))
	}
	return ema
}
