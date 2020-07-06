package amiibo_test

import (
	"path/filepath"
	"reflect"
	"runtime"
	"testing"

	"github.com/lindsaygelle/amiibo"
)

var jpnLineup amiibo.JPNLineup

func TestGetJPNLineup(t *testing.T) {
	var err error
	_, caller, _, _ := runtime.Caller(0)
	filefolder := filepath.Dir(caller)
	filename := "jpn-lineup.xml"
	jpnLineup, err = amiibo.ReadJPNLineup(filefolder, filename)
	if err != nil {
		t.Log("amiibo.ReadJPNLineup", err)
	}
	if err != nil {
		_, _, jpnLineup, err = amiibo.GetJPNLineup()
		if err != nil {
			t.Fatal(err)
		}
	}
	s, err := amiibo.WriteJPNLineup(filefolder, filename, jpnLineup)
	if err != nil {
		t.Fatal(err)
	}
	if s != filepath.Join(filefolder, filename) {
		t.Fatalf("%s != %s", s, filepath.Join(filefolder, filename))
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
