package mix_test

import (
	"testing"

	"github.com/gellel/amiibo/compatability"
	"github.com/gellel/amiibo/lineup"
	"github.com/gellel/amiibo/mix"
)

const (
	templateAmiiboErr string = "mix.Amiibo.%s: m.(%s) %s %s"
)

func TestAmiibo(t *testing.T) {
	var (
		c = &compatability.Amiibo{}
		i = &lineup.Item{}
		l = &lineup.Amiibo{}
	)
	var (
		m, err = mix.NewAmiibo(c, i, l)
	)
	if err != nil {
		t.Fatalf(err.Error())
	}
	if m.Compatability == nil {
		t.Fatalf(templateAmiiboErr, "Compatability", "Compatability", "==", "nil")
	}
	if m.Item == nil {
		t.Fatalf(templateAmiiboErr, "Item", "Item", "==", "nil")
	}
	if m.Lineup == nil {
		t.Fatalf(templateAmiiboErr, "Item", "Item", "==", "nil")
	}
}
