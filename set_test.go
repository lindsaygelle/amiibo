package amiibo_test

import (
	"reflect"
	"testing"

	"github.com/gellel/amiibo"
)

func TestSet(t *testing.T) {

	if ok := reflect.ValueOf(amiibo.NewSet()).Kind() == reflect.Ptr; ok != true {
		t.Fatalf("reflect.ValueOf(amiibo.NewSet()) != reflect.Ptr")
	}
}

func TestSetAdd(t *testing.T) {

	set := amiibo.NewSet()

	set.Add(&amiibo.Amiibo{Name: "a"})

	slice, ok := set.Get("a")

	if ok != true {
		t.Fatalf("set.Add(amiibo *amiibo.Amiibo) did not add a new key reference")
	}
	if slice == nil {
		t.Fatalf("set.Add(amiibo *amiibo.Amiibo) did not instantiate a new amiibo.Slice pointer")
	}

	previousLength := slice.Len()

	set.Add(&amiibo.Amiibo{Name: "a"})

	currentLength := slice.Len()

	if ok = previousLength < currentLength; ok != true {
		t.Fatalf("set.Add(amiibo *amiibo.Amiibo) did not push a same-key value onto the reference amiibo.Slice pointer")
	}
}

func TestSetSize(t *testing.T) {

	set := amiibo.NewSet().Add(&amiibo.Amiibo{Name: "a"}).Add(&amiibo.Amiibo{Name: "a"}).Add(&amiibo.Amiibo{Name: "b"})

	if ok := set.Size() == 3; ok != true {
		t.Fatalf("set.Size() did not calculate the accurate size of the set")
	}
}
