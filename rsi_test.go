package ta

import (
	"testing"
)

func TestRsi(t *testing.T) {
	compare(t, "result = talib.RSI(testClose, 10)", BatchUpdateFloat64(NewRSI(10), testClose))
}
