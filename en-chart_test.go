package amiibo_test

import (
	"testing"

	"github.com/lindsaygelle/amiibo"
)

func TestGetENChart(t *testing.T) {
	return
	_, _, v, err := amiibo.GetENGChart()
	if err != nil {
		t.Fatal(err)
	}
	if l := len(v.AmiiboList); l == 0 {
		t.Fatal("len: v.AmiiboList")
	}
	if l := len(v.GameList); l == 0 {
		t.Fatalf("len: v.GameList")
	}
	if l := len(v.Items); l == 0 {
		t.Fatalf("len: v.Items")
	}
}
