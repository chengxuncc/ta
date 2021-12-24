package ta

import (
	"time"

	"github.com/shopspring/decimal"
)

// TimeSeries represents an array of candles
type TimeSeries struct {
	Candles []*Candle
}

// NewTimeSeries returns a new, empty, TimeSeries
func NewTimeSeries() (t *TimeSeries) {
	t = new(TimeSeries)
	t.Candles = make([]*Candle, 0)

	return t
}

// AddCandle adds the given candle to this TimeSeries if it is not nil and after the last candle in this timeseries.
// If the candle is added, AddCandle will return true, otherwise it will return false.
func (ts *TimeSeries) AddCandle(candle *Candle) bool {
	if candle == nil {
		return false
	}

	if ts.LastCandle() == nil || candle.Period.End.After(ts.LastCandle().Period.End) {
		ts.Candles = append(ts.Candles, candle)
		return true
	}

	return false
}

func (ts *TimeSeries) AddCandleRealtime(candle *Candle) bool {
	if candle == nil {
		return false
	}

	now := time.Now()
	if !(now.After(candle.Period.Start) && now.After(candle.Period.End)) {
		return false
	}
	if len(ts.Candles) == 0 {
		ts.Candles = append(ts.Candles, candle)
		return true
	} else if candle.Period.End.After(ts.LastCandle().Period.End) {
		ts.Candles = append(ts.Candles[1:], candle)
		return true
	}
	return false
}

// LastCandle will return the lastCandle in this series, or nil if this series is empty
func (ts *TimeSeries) LastCandle() *Candle {
	if len(ts.Candles) > 0 {
		return ts.Candles[len(ts.Candles)-1]
	}

	return nil
}

// LastIndex will return the index of the last candle in this series
func (ts *TimeSeries) LastIndex() int {
	return len(ts.Candles) - 1
}

func (ts *TimeSeries) Calculate(indicator Indicator) decimal.Decimal {
	var v decimal.Decimal
	for _, candle := range ts.Candles {
		v = indicator.Update(candle.ClosePrice)
	}
	return v
}
