package app

import (
	"context"
	"fmt"
	"time"

	"github.com/lacsar712/asphaltm/internal/model"
)

// CalibrateProbe allows acceptance tests to inject feed calibration faults.
var CalibrateProbe func(ctx context.Context) error

const feedTempLimitC = 55.0

func (a *App) CalibrateFeed(ctx context.Context, tower model.TowerID, holder string) error {
	if err := a.feedLeases.Require(tower, holder, 30*time.Second); err != nil {
		return err
	}
	// Release on every return path, including probe faults that abort a
	// calibration mid-step (e.g. weighing-line loose in step 2). Otherwise the
	// aggregate scale stays leased and the next holder is blocked until a
	// whole-station reset, with no Release record in the aggregate log.
	defer a.feedLeases.ReleaseHolder(tower, holder)
	if CalibrateProbe != nil {
		if err := CalibrateProbe(ctx); err != nil {
			return fmt.Errorf("calibrate: %w", err)
		}
	}
	return nil
}
