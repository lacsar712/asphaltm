package app

import (
	"context"
	"errors"
	"testing"

	"github.com/lacsar712/asphaltm/internal/config"
	"github.com/lacsar712/asphaltm/internal/model"
)

func TestCase(t *testing.T) {
	a, err := New(config.Default())
	if err != nil {
		t.Fatal(err)
	}
	err = a.HandleBinderTrip(context.Background(), model.TowerID(a.cfg.TowerID), 95.0)
	if err == nil {
		t.Fatal("expected heat overtemperature error")
	}
	if !errors.Is(err, model.ErrBinderTrip) {
		t.Fatalf("expected ErrBinderTrip, got %v", err)
	}
}
