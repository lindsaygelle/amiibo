package amiibo

import (
	"crypto/md5"
	"fmt"
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

	if n := amiiboSlice.Push(&Amiibo{ID: fmt.Sprintf("%x", md5.Sum([]byte("Test"))), Name: "Test"}); n != 2 {
		t.Fatalf("AmiiboSlice.Push(a *Amiibo) int did not return the expected length; %v != 2", n)
	}

	amiiboSlice.Assign(testAmiiboStruct, testAmiiboStruct, &Amiibo{ID: fmt.Sprintf("%x", md5.Sum([]byte("Test"))), Name: "Test"})

	N := amiiboSlice.Len()

	amiiboSlice = amiiboSlice.Set()

	if ok := amiiboSlice.Len() != N; ok != true {
		t.Fatalf("AmiiboSlice.Set() *AmiiboSlice did not correctly trim the length of the AmiiboSlice")
	}

	if ok := amiiboSlice.Len() == 2; ok != true {
		t.Fatalf("AmiiboSlice.Set() *AmiiboSlice did not build the expected AmiiboSlice")
	}

	a, b := amiiboSlice.Fetch(0), amiiboSlice.Fetch(1)

	amiiboSlice.Swap(0, 1)

	c, d := amiiboSlice.Fetch(0), amiiboSlice.Fetch(1)

	if ok := a == d && b == c; ok != true {
		t.Fatalf("AmiiboSlice.Swap(i, j int) bool did not swap i to j and j to i")
	}

	n := amiiboSlice.Len()
	amiiboSlice.EachReverse(func(i int, _ *Amiibo) {
		if i > n {
			t.Fatalf("AmiiboSlice.EachReverse(i int, a *Amiibo) *AmiiboSlice does not iterate in reverse order; %v > %v", i, n)
		}
	})
}
