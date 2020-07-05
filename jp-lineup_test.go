package amiibo_test

import (
	"testing"

	"github.com/lindsaygelle/amiibo"
)

func TestGetJPNLineup(t *testing.T) {

	_, _, v, err := amiibo.GetJPNLineup()
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
