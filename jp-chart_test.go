package amiibo_test

import (
	"log"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"

	"github.com/lindsaygelle/amiibo"
)

var jpnChart amiibo.JPNChart

func TestGetJPNChart(t *testing.T) {
	var err error
	_, caller, _, _ := runtime.Caller(0)
	filefolder := filepath.Dir(caller)
	filename := "jpn-chart.xml"
	jpnChart, err = amiibo.ReadJPNChart(filefolder, filename)
	if err != nil {
		log.Println("amiibo.ReadJPNChart", err)
	}
	if err != nil {
		_, _, jpnChart, err = amiibo.GetJPNChart()
		if err != nil {
			t.Fatal(err)
		}
	}
	s, err := amiibo.WriteJPNChart(filefolder, filename, jpnChart)
	if err != nil {
		t.Fatal(err)
	}
	if s != filepath.Join(filefolder, filename) {
		t.Fatalf("%s != %s", s, filepath.Join(filefolder, filename))
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
