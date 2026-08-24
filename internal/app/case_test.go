package app

import (
	"context"
	"testing"

	"github.com/lacsar712/asphaltm/internal/config"
	"github.com/lacsar712/asphaltm/internal/model"
)

func TestCase(t *testing.T) {
	a, err := New(config.Default())
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()
	scopeA, _ := a.BeginBatchScope(ctx, model.TowerID("tower-a"))
	scopeB, releaseB := a.BeginBatchScope(ctx, model.TowerID("tower-b"))
	defer releaseB()
	if scopeA.Err() != nil {
		t.Fatal("tower A batch scope cancelled when tower B began")
	}
	if scopeB.Err() != nil {
		t.Fatal("tower B batch scope cancelled at start")
	}
}
