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
	m := make(map[string]int)
	for _, v := range v.Items {
		if _, ok := m[v.Code]; !ok {
			m[v.Code] = 0
		}
		m[v.Code] = m[v.Code] + 1
	}
	a, b := len(m), len(v.Items)
	if a != b {
		t.Fatalf("len: v.Items %d != %d", a, b)
	}
}
