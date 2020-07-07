package amiibo_test

import (
	"testing"

	"github.com/lindsaygelle/amiibo"
)

func testJPN(t *testing.T, jpnChart amiibo.JPNChart, jpnLineup amiibo.JPNLineup) {
	m := make(map[string]int)
	for _, v := range jpnChart.Items {
		ID := v.GetID()
		if _, ok := m[ID]; !ok {
			m[ID] = 0
		}
		m[ID] = m[ID] + 1
	}
	for _, v := range jpnLineup.Items {
		ID := v.GetID()
		if _, ok := m[ID]; !ok {
			m[ID] = 0
		}
		m[ID] = m[ID] + 1
	}
	a, b := len(jpnChart.Items), len(jpnLineup.Items)
	if a != b {
		t.Logf("len: jpnChart.Items %d != jpnLineup.Items %d", a, b)
	}
}
