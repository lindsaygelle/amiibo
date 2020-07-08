package amiibo_test

import (
	"testing"

	"github.com/lindsaygelle/amiibo"
)

var jpnAmiiboMapFileName = "jpn-amiibo-map.json"

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
	_, err = amiibo.WriteJPNAmiiboMap(filefolder, jpnAmiiboMapFileName, &v)
	if err != nil {
		t.Fatal(err)
	}
	_, err = amiibo.ReadJPNAmiiboMap(filefolder, jpnAmiiboMapFileName)
	if err != nil {
		t.Fatal(err)
	}
}
