package amiibo_test

import (
	"testing"

	"github.com/lindsaygelle/amiibo"
)

var engGameMapFileName = "eng-game-map.json"

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
	_, err = amiibo.WriteENGGameMap(filefolder, engGameMapFileName, &v)
	if err != nil {
		t.Fatal(err)
	}
	_, err = amiibo.ReadENGGameMap(filefolder, engGameMapFileName)
	if err != nil {
		t.Fatal(err)
	}
}
