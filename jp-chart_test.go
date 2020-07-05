package amiibo_test

import (
	"reflect"
	"testing"

	"github.com/lindsaygelle/amiibo"
)

var jpnChart amiibo.JPNChart

func TestGetJPNChart(t *testing.T) {
	var err error
	_, _, jpnChart, err = amiibo.GetJPNChart()
	if err != nil {
		t.Fatal(err)
	}
	if l := len(jpnChart.Items); l == 0 {
		t.Fatal("len: v.Items", l)
	}
	m := make(map[string]int)
	for _, v := range jpnChart.Items {
		ID := v.GetID()
		if _, ok := m[ID]; !ok {
			m[ID] = 0
		}
		m[ID] = m[ID] + 1
	}
	a, b := len(m), len(jpnChart.Items)
	if a != b {
		t.Fatalf("len: v.Items %d != %d", a, b)
	}
	if !reflect.ValueOf(jpnChart).IsZero() && !reflect.ValueOf(jpnLineup).IsZero() {
		testJPN(t, jpnChart, jpnLineup)
	}
}
