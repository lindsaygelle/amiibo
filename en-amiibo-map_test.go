package amiibo_test

import (
	"fmt"
	"testing"

	"github.com/lindsaygelle/amiibo"
)

func testENGAmiiboMap(t *testing.T) {
	var v, err = amiibo.NewENGAmiiboMap(engChart, engLineup)
	if err != nil {
		t.Fatal(err)
	}
	if l := len(v); l == ((len(engChart.AmiiboList) + len(engLineup.AmiiboList) + len(engLineup.Items)) / 3) {
		t.Fatalf("engAmiiboMap %d engChart.AmiiboList %d engLineup.AmiiboList %d engLineup.Items %d", l, len(engChart.AmiiboList), len(engLineup.AmiiboList), len(engLineup.Items))
	}
	for _, v := range v {
		fmt.Println(v.Description)
	}
}
