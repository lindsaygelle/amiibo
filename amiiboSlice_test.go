package amiibo

import (
	"testing"
)

func TestAmiiboSlice(t *testing.T) {

	t.Parallel()

	amiiboSlice := newAmiiboSlice()

	if testAmiiboStruct == nil {
		t.Fatalf("amiibo.newAmiiboMap() test cannot run testAmiiboStruct is nil")
	}

	if s := amiiboSlice.Append(testAmiiboStruct); s != amiiboSlice {
		t.Fatalf("AmiiboSlice.Append(a *Amiibo) *AmiiboSlice != %v", amiiboSlice)
	}

	if ok := amiiboSlice.Bounds(0); ok != true {
		t.Fatalf("AmiiboSlice.Bounds(i int) bool != true; index of 0 is in bounds but returned false")
	}

	a, ok := amiiboSlice.Get(0)

	if ok != true {
		t.Fatalf("AmiiboSlice.Get(i int) (*Amiibo, bool) bool != true")
	}

	if a != testAmiiboStruct {
		t.Fatalf("AmiiboSlice.Get(i int) (*Amiibo, bool) %v != %v", a, testAmiiboStruct)
	}
}
