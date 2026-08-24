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

// Active reports whether the heat window is still open relative to the
// injected clock, not the wall clock. The process beat (ProcessClock) can be
// frozen during a hold/pause; measuring elapsed time against that clock keeps
// the closure window in step with the process beat instead of the wall clock.
func (w *HeatWindow) Active(anchor time.Time) bool {
	return w.clk.Now().Sub(anchor) < w.duration
}

func (w *HeatWindow) Require(anchor time.Time) error {
	if w.Active(anchor) {
		return nil
	}
	return model.ErrHeatHold
}
