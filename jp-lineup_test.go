package amiibo

import (
	"testing"
)

func TestGetLineupJPN(t *testing.T) {

	_, _, v, err := getJPNLineup()
	if err != nil {
		t.Fatal(err)
	}
	if l := len(v.Items); l == 0 {
		t.Fatal("len: v.Items", l)
	}
	if l := len(v.SeriesItems); l == 0 {
		t.Fatal("len: v.SeriesItems", l)
	}
}
