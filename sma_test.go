package ta

import (
	"testing"
)

func TestSma(t *testing.T) {
	compare(t, "result = talib.SMA(testClose, 20)", BatchUpdateFloat64(NewSMA(20), testClose))
}
