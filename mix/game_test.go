package mix_test

import (
	"testing"

	"github.com/gellel/amiibo/compatability"
	"github.com/gellel/amiibo/mix"
)

const (
	templateGameErr string = "mix.Game.%s: m.(%s) %s %s"
)

func TestGame(t *testing.T) {
	var (
		g = &compatability.Game{}
		i = &compatability.Item{}
	)
	var (
		m, err = mix.NewGame(g, i)
	)
	if err != nil {
		t.Fatalf(err.Error())
	}
	if m.Game == nil {
		t.Fatalf(templateGameErr, "Game", "Game", "==", "nil")
	}
	if m.Item == nil {
		t.Fatalf(templateGameErr, "Item", "Item", "==", "nil")
	}
}
