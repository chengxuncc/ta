package ta

import (
	"github.com/shopspring/decimal"
)

// Developed by J. Welles Wilder, the Relative Strength Index (RSI) is a momentum
// oscillator that measures the speed and change of price movements. RSI oscillates
// between zero and 100. According to Wilder, RSI is considered overbought when
// above 70 and oversold when below 30. Signals can also be generated by looking
// for divergences, failure swings and centerline crossovers. RSI can also be used
// to identify the general trend.
// RSI is an extremely popular momentum indicator that has been featured in a number
// of articles, interviews and books over the years. In particular, Constance Brown's
// book, Technical Analysis for the Trading Professional, features the concept of
// bull market and bear market ranges for RSI. Andrew Cardwell, Brown's RSI mentor,
// introduced positive and negative reversals for RSI and, additionally, turned the
// notion of divergence, literally and figuratively, on its head.
// Wilder features RSI in his 1978 book, New Concepts in Technical Trading Systems.
// This book also includes the Parabolic SAR, Average True Range and the Directional
// Movement Concept (ADX). Despite being developed before the computer age,
// Wilder's indicators have stood the test of time and remain extremely popular.
//  https://school.stockcharts.com/doku.php?id=technical_indicators:relative_strength_index_rsi
//  https://www.investopedia.com/terms/r/rsi.asp
//  https://www.fidelity.com/learning-center/trading-investing/technical-analysis/technical-indicator-guide/RSI

var oneHundred = decimal.NewFromInt(100)

func NewRSI(window int) *RelativeStrengthIndex {
	return &RelativeStrengthIndex{
		window:      window,
		AverageGain: Indicators{NewGain(), NewEMAWithSmoothing(window, One)},
		AverageLoss: Indicators{NewLoss(), NewEMAWithSmoothing(window, One)},
	}
}

var _ Indicator = (*RelativeStrengthIndex)(nil)

type RelativeStrengthIndex struct {
	window      int
	AverageGain Indicator
	AverageLoss Indicator
	count       int
}

func (r *RelativeStrengthIndex) WindowSize() int {
	return r.window
}

func (r *RelativeStrengthIndex) Update(value decimal.Decimal) decimal.Decimal {
	r.count++
	averageGain := r.AverageGain.Update(value)
	averageLoss := r.AverageLoss.Update(value)
	if r.count <= r.window {
		return Zero
	}
	sum := averageGain.Add(averageLoss)
	if sum.IsZero() {
		return Zero
	}
	return averageGain.Div(sum).Mul(oneHundred)
}

func (r *RelativeStrengthIndex) DryUpdate(value decimal.Decimal) decimal.Decimal {
	r.count++
	averageGain := r.AverageGain.DryUpdate(value)
	averageLoss := r.AverageLoss.DryUpdate(value)
	if r.count <= r.window {
		return Zero
	}
	sum := averageGain.Add(averageLoss)
	if sum.IsZero() {
		return Zero
	}
	return averageGain.Div(sum).Mul(oneHundred)
}
