package amiibo_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gellel/amiibo"
)

func TestSlice(t *testing.T) {

	if ok := reflect.ValueOf(amiibo.NewSlice()).Kind() == reflect.Ptr; ok != true {
		t.Fatalf("reflect.ValueOf(amiibo.NewSlice()) != reflect.Ptr")
	}
}

func TestSliceAppend(t *testing.T) {

	slice := amiibo.NewSlice()

	previousLength := slice.Len()

	slice.Append(&amiibo.Amiibo{Name: "test"})

	currentLength := slice.Len()

	if ok := previousLength < currentLength; ok != true {
		t.Fatalf("slice.Append(amiibo *Amiibo) did not append an amiibo to the slice")
	}
}

func TestSliceAssign(t *testing.T) {

	slice := amiibo.NewSlice()

	previousLength := slice.Len()

	slice.Assign(&amiibo.Amiibo{Name: "a"}, &amiibo.Amiibo{Name: "b"})

	currentLength := slice.Len()

	if ok := previousLength < currentLength; ok != true {
		t.Fatalf("slice.Assign(amiibo ...*Amiibo) did not append two amiibo's to the slice")
	}
}

func TestSliceConcatenate(t *testing.T) {

	a := amiibo.NewSlice(&amiibo.Amiibo{Name: "a"})
	b := amiibo.NewSlice(&amiibo.Amiibo{Name: "b"})
	c := amiibo.NewSlice()

	previousLength := c.Len()

	c.Concatenate(a).Concatenate(b)

	currentLength := c.Len()

	if ok := previousLength < currentLength; ok != true {
		t.Fatalf("slice.Concatenate(slice *amiibo.Slice) did not append an amiibo to the slice")
	}
	if ok := currentLength == (a.Len() + b.Len()); ok != true {
		t.Fatalf("slice.Concatenate(slcie *amiibo.Slice) did not increment the length of the receiver slice to the sum of A + B")
	}
}

func TestSliceEach(t *testing.T) {

	amiibo.NewSlice(&amiibo.Amiibo{Name: "test"}).Each(func(i int, amiibo *amiibo.Amiibo) {})
}

func TestSliceMap(t *testing.T) {

	slice := amiibo.NewSlice(&amiibo.Amiibo{Name: "a"}).Map(func(i int, amiibo *amiibo.Amiibo) *amiibo.Amiibo {
		amiibo.Name = fmt.Sprintf("%v", i)
		return amiibo
	})
	if amiibo := slice.Fetch(0); amiibo != nil {
		if amiibo.Name != "0" {
			t.Fatalf("slice.Map(f func(i int, amiibo *amiibo.Amiibo) did not mutate the amiibo at the reference integer")
		}
	}
}

func TestSliceSplice(t *testing.T) {

	a := amiibo.NewSlice(&amiibo.Amiibo{Name: "a"}, &amiibo.Amiibo{Name: "b"})

	b := a.Splice(0, 1)

	a.Each(func(i int, amiibo *amiibo.Amiibo) {
		fmt.Println(amiibo.Name)
	})

	b.Each(func(i int, amiibo *amiibo.Amiibo) {
		fmt.Println(amiibo.Name)
	})

	if ok := (a.Len() + b.Len()) == 2; ok != true {
		t.Fatalf("slice.Splice(start, end int) did not truncate the length of the slice")
	}

	if ok := a.Poll().Name == "b" && b.Poll().Name == "a"; ok != true {
		t.Fatalf("slice.Splice(start, end int) did not consume the first N elements from base slice")
	}
}
