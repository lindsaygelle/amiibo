package amiibo_test

import (
	"testing"

	"github.com/lindsaygelle/amiibo"
)

func testJPNAmiiboMap(t *testing.T) {
	var v, err = amiibo.NewJPNAmiiboMap(&jpnChart, &jpnLineup)
	if err != nil {
		t.Fatal(err)
	}
	if len(v) == 0 {
		t.Fatal("len(jpnAmiiboMap) == 0")
	}
	if l := len(v); l != ((len(jpnChart.Items) + len(jpnLineup.Items)) / 2) {
		t.Logf("jpnAmiiboMap %d jpnChart.Items %d jpnLineup.Items %d", l, len(jpnChart.Items), len(jpnLineup.Items))
	}
	_, err = amiibo.WriteJPNAmiiboMap(filefolder, "jp-amiibo-map.json", &v)
	if err != nil {
		t.Fatal(err)
	}
}
