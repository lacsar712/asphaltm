package app

import (
	"context"

	"github.com/lacsar712/asphaltm/internal/model"
)

func (a *App) BeginBatchScope(ctx context.Context, tower model.TowerID) (context.Context, context.CancelFunc) {
	child, cancel := context.WithCancel(ctx)
	a.batchMu.Lock()
	if prev, ok := a.activeCancel[tower]; ok && prev != nil {
		prev()
	}
	a.activeCancel[tower] = cancel
	a.batchMu.Unlock()
	release := func() {
		a.batchMu.Lock()
		if cur, ok := a.activeCancel[tower]; ok && cur != nil {
			delete(a.activeCancel, tower)
		}
		a.batchMu.Unlock()
		cancel()
	}
	return child, release
}

func (a *App) RunBatch(ctx context.Context, tower model.TowerID, fn func(context.Context) error) error {
	batchCtx, release := a.BeginBatchScope(ctx, tower)
	defer release()
	return fn(batchCtx)
}
