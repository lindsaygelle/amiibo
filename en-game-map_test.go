package amiibo_test

import (
	"testing"

	"github.com/lindsaygelle/amiibo"
)

func testENGGameMap(t *testing.T) {
	var v, err = amiibo.NewENGGameMap(&engChart)
	if err != nil {
		t.Fatal(err)
	}
	if len(v) == 0 {
		t.Fatal("len(engChart) == 0")
	}
	if l := len(v); l != ((len(engChart.GameList) + len(engChart.Items)) / 2) {
		t.Logf("engGameMap %d engChart.GameList %d engChart.Items %d", l, len(engChart.GameList), len(engChart.Items))
	}
}
