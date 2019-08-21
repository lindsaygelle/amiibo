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

	a = amiiboSlice.Poll()

	if amiiboSlice.Len() != 0 {
		t.Fatalf("AmiiboSlice.Poll() *Amiibo did not reduce the length of the AmiiboSlice")
	}

	if ok := amiiboSlice.Assign(a).Len() == 1; ok != true {
		t.Fatalf("AmiiboSlice.Assign(a ...*Amiibo) *AmiiboSlice did not increase the length of the AmiiboSlice")
	}

	if n := amiiboSlice.Push(&Amiibo{Name: "Test"}); n != 2 {
		t.Fatalf("AmiiboSlice.Push(a *Amiibo) int did not return the expected length; %v != 2", n)
	}

	amiiboSlice.Assign(testAmiiboStruct, testAmiiboStruct, &Amiibo{Name: "Test"})

	previousN := amiiboSlice.Len()

	amiiboSlice = amiiboSlice.Set()

	if currentN := amiiboSlice.Len(); currentN == previousN {
		t.Fatalf("AmiiboSlice.Set() *AmiiboSlice did not trim the length of the AmiiboSlice")
	}
}
