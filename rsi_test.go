package ta

import (
	"testing"
)

func TestRsi(t *testing.T) {
	rsi := NewRelativeStrengthIndex(10)
	compare(t, "result = talib.RSI(testClose, 10)", BatchUpdate(rsi, Float64sToDecimals(testClose)))
}
