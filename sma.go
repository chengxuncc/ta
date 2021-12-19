package ta

import (
	"github.com/shopspring/decimal"
)

// A simple moving average is formed by computing the average price of a security
// over a specific number of periods. Most moving averages are based on closing
// prices; for example, a 5-day simple moving average is the five-day sum of closing
// prices divided by five. As its name implies, a moving average is an average that
// moves. Old data is dropped as new data becomes available, causing the average
// to move along the time scale. The example below shows a 5-day moving average
// evolving over three days.
//  https://school.stockcharts.com/doku.php?id=technical_indicators:moving_averages
//  https://www.investopedia.com/terms/s/sma.asp
//  https://www.fidelity.com/learning-center/trading-investing/technical-analysis/technical-indicator-guide/sma

func NewSMA(window int) *SimpleMovingAverage {
	return &SimpleMovingAverage{
		window:        window,
		windowDecimal: decimal.NewFromInt(int64(window)),
		history:       NewHistory(window),
	}
}

var _ Indicator = (*SimpleMovingAverage)(nil)

type SimpleMovingAverage struct {
	window        int
	windowDecimal decimal.Decimal
	history       *History
	sum           decimal.Decimal
	count         int
}

func (s *SimpleMovingAverage) WindowSize() int {
	return s.window
}

func (s *SimpleMovingAverage) Update(value decimal.Decimal) decimal.Decimal {
	old := s.history.Update(value)
	s.sum = s.sum.Add(value.Sub(old))
	if s.history.count < s.window {
		return Zero
	}
	return s.sum.Div(s.windowDecimal)
}

func (s *SimpleMovingAverage) DryUpdate(value decimal.Decimal) decimal.Decimal {
	old := s.history.DryUpdate(value)
	sum := s.sum.Add(value.Sub(old))
	if s.history.count < s.window {
		return Zero
	}
	return sum.Div(s.windowDecimal)
}
