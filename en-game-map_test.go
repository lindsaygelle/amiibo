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
	_, err = amiibo.WriteENGGameMap(filefolder, "en-game-map.json", &v)
	if err != nil {
		t.Fatal(err)
	}
	_, err = amiibo.ReadENGGameMap(filefolder, "en-game-map.json")
	if err != nil {
		t.Fatal(err)
	}
}
