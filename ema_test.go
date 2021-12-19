package ta

import (
	"testing"
)

func TestEma(t *testing.T) {
	compare(t, "result = talib.EMA(testClose, 5)", BatchUpdateFloat64(NewEMA(5), testClose))
	compare(t, "result = talib.EMA(testClose, 20)", BatchUpdateFloat64(NewEMA(20), testClose))
	compare(t, "result = talib.EMA(testClose, 50)", BatchUpdateFloat64(NewEMA(50), testClose))
	compare(t, "result = talib.EMA(testClose, 100)", BatchUpdateFloat64(NewEMA(100), testClose))
}
