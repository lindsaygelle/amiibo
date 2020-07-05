package amiibo

import (
	"testing"
)

func TestGetLineupComingJPN(t *testing.T) {

	_, _, v, err := getLineupComingJPN()
	if err != nil {
		t.Fatal(err)
	}
	if l := len(v.Items); l == 0 {
		t.Fatal("len: v.Items", l)
	}
}
