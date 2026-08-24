package clock

import (
	"time"

	"github.com/lacsar712/asphaltm/internal/model"
)

type HeatWindow struct {
	clk      Clock
	duration time.Duration
}

func NewHeatWindow(clk Clock, duration time.Duration) *HeatWindow {
	if duration <= 0 {
		duration = 2 * time.Minute
	}
	return &HeatWindow{clk: clk, duration: duration}
}

func (w *HeatWindow) Active(anchor time.Time) bool {
	return WindowElapsed(w.clk, anchor, w.duration)
}

func (w *HeatWindow) Require(anchor time.Time) error {
	if w.Active(anchor) {
		return nil
	}
	return model.ErrHeatHold
}
