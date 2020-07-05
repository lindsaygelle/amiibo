package amiibo

import (
	"testing"
)

func TestGetLineup(t *testing.T) {

	_, _, v, err := getLineup()
	if err != nil {
		t.Fatal(err)
	}
	if l := len(v.AmiiboList); l == 0 {
		t.Fatal("len: v.AmiiboList", l)
	}
	if l := len(v.Items); l == 0 {
		t.Fatal("len: v.Items", l)
	}
}
