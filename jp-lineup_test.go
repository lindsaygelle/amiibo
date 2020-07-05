package amiibo_test

import (
	"reflect"
	"testing"

	"github.com/lindsaygelle/amiibo"
)

var jpnLineup amiibo.JPNLineup

func TestGetJPNLineup(t *testing.T) {
	var err error
	_, _, jpnLineup, err = amiibo.GetJPNLineup()
	if err != nil {
		t.Fatal(err)
	}
	if l := len(jpnLineup.Items); l == 0 {
		t.Fatal("len: jpnLineup.Items", l)
	}
	if l := len(jpnLineup.SeriesItems); l == 0 {
		t.Fatal("len: jpnLineup.SeriesItems", l)
	}
	m := make(map[string]int)
	for _, v := range jpnLineup.Items {
		ID := v.GetID()
		if _, ok := m[ID]; !ok {
			m[ID] = 0
		}
		m[ID] = m[ID] + 1
	}
	a, b := len(m), len(jpnLineup.Items)
	if a != b {
		t.Fatalf("len: jpnLineup.Items %d != %d", a, b)
	}
	m = make(map[string]int)
	for _, v := range jpnLineup.SeriesItems {
		ID := v.Name
		if _, ok := m[ID]; !ok {
			m[ID] = 0
		}
		m[ID] = m[ID] + 1
	}
	a, b = len(m), len(jpnLineup.SeriesItems)
	if a != b {
		t.Fatalf("len: jpnLineup.SeriesItems %d != %d", a, b)
	}
	if !reflect.ValueOf(jpnChart).IsZero() && !reflect.ValueOf(jpnLineup).IsZero() {
		testJPN(t, jpnChart, jpnLineup)
	}
}
