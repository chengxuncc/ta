package ta

import (
	"github.com/shopspring/decimal"
)

func NewHistory(window int) *History {
	return &History{
		window: window,
		values: make([]decimal.Decimal, window),
	}
}

var _ Indicator = (*History)(nil)

type History struct {
	window      int
	values      []decimal.Decimal
	oldestIndex int
	newestIndex int
	count       int
}

func (h *History) WindowSize() int {
	return h.window
}

// Update push the latest value and return oldest one
func (h *History) Update(value decimal.Decimal) decimal.Decimal {
	oldest := h.values[h.oldestIndex]
	h.values[h.oldestIndex] = value
	h.newestIndex, h.oldestIndex = h.oldestIndex, (h.oldestIndex+1)%h.window
	h.count++
	return oldest
}

func (h *History) DryUpdate(value decimal.Decimal) decimal.Decimal {
	return h.values[h.oldestIndex]
}
