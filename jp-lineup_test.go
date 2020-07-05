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
	m := make(map[string]int)
	for _, v := range v.Items {
		if _, ok := m[v.GetID()]; !ok {
			m[v.GetID()] = 0
		}
		m[v.GetID()] = m[v.GetID()] + 1
	}
	a, b := len(m), len(v.Items)
	if a != b {
		t.Fatalf("len: v.Items %d != %d", a, b)
	}
	m = make(map[string]int)
	for _, v := range v.SeriesItems {
		if _, ok := m[v.Name]; !ok {
			m[v.Name] = 0
		}
		m[v.Name] = m[v.Name] + 1
	}
	a, b = len(m), len(v.SeriesItems)
	if a != b {
		t.Fatalf("len: v.SeriesItems %d != %d", a, b)
	}
}
