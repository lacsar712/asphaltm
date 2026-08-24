package app

import (
	"context"

	"github.com/lacsar712/asphaltm/internal/clock"
)

// SegmentPlan describes inter-segment vent staging for drying batches.

type SegmentPlan struct {
	VentSteps int
}

func (a *App) ExecutePlan(ctx context.Context, plan SegmentPlan) error {
	if a.scheduler == nil {
		return nil
	}
	// Propagate the caller's ctx (which carries the mix-screen withdrawal /
	// cancel receipt) into the scheduling layer. Previously this handed
	// context.Background() to InstallVentPlanCtx, so a cancellation raised at the
	// mixing screen never reached the plan layer and the scheduler kept
	// appending heating/vent steps from the stale plan table.
	return a.scheduler.InstallVentPlanCtx(ctx, clock.VentPlan{VentSteps: plan.VentSteps}, "segment-plan")
}

func (a *App) SegmentVentStepsDone() int {
	if a.scheduler == nil {
		return 0
	}
	return a.scheduler.VentStepsDone()
}
