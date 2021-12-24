package ta

import (
	"time"

	"github.com/shopspring/decimal"
)

// TimeSeries represents an array of candles
type TimeSeries struct {
	Candles []*Candle
	Limit   int
}

// NewTimeSeries returns a new, empty, TimeSeries
func NewTimeSeries(limit int) (t *TimeSeries) {
	return &TimeSeries{
		Candles: make([]*Candle, 0, limit),
		Limit:   limit,
	}
}

// AddCandle adds the given candle to this TimeSeries if it is not nil and after the last candle in this timeseries.
// If the candle is added, AddCandle will return true, otherwise it will return false.
func (ts *TimeSeries) AddCandle(candle *Candle) bool {
	if candle == nil {
		return false
	}
	if ts.LastCandle() != nil && !candle.Period.End.After(ts.LastCandle().Period.End) {
		return false
	}

	if ts.Limit > 0 && len(ts.Candles) == ts.Limit {
		ts.Candles = append(ts.Candles[1:], candle)
	} else {
		ts.Candles = append(ts.Candles, candle)
	}
	return true
}

func (ts *TimeSeries) AddCandleRealtime(candle *Candle) bool {
	if candle == nil {
		return false
	}
	now := time.Now()
	if !(now.After(candle.Period.Start) && now.After(candle.Period.End)) {
		return false
	}
	return ts.AddCandle(candle)
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
