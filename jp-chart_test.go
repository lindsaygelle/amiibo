package amiibo_test

import (
	"testing"

	"github.com/lindsaygelle/amiibo"
)

func TestGetJPNChart(t *testing.T) {
	_, _, v, err := amiibo.GetJPNChart()
	if err != nil {
		t.Fatal(err)
	}
	if l := len(v.Items); l == 0 {
		t.Fatal("len: v.Items", l)
	}
}
