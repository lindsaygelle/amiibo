package amiibo_test

import (
	"testing"

	"github.com/lindsaygelle/amiibo"
)

func testJPNAmiiboMap(t *testing.T) {
	var v, err = amiibo.NewJPNAmiiboMap(jpnChart, jpnLineup)
	if err != nil {
		t.Fatal(err)
	}
	if l := len(v); l != ((len(jpnChart.Items) + len(jpnLineup.Items)) / 2) {
		t.Fatalf("jpnAmiiboMap %d jpnChart.Items %d jpnLineup.Items %d", l, len(jpnChart.Items), len(jpnLineup.Items))
	}
	// for _, v := range v {
	//	fmt.Println(len(v.Software))
	// }
	_, err = amiibo.WriteJPNAmiiboMap(filefolder, "jp-amiibo-map.json", v)
}
