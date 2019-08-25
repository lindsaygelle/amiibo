package amiibo

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRawItem(t *testing.T) {

	b := []byte(rawItemDefault)

	rawItem := NewRawItem(&b)

	if ok := reflect.ValueOf(rawItem).Kind() == reflect.Ptr; ok != true {
		t.Fatalf("amiibo.NewRawItem(b *[]byte) *amiibo.RawItem != uintptr")
	}

	returnType := reflect.TypeOf(rawItem).Elem().String()

	expectType := reflect.TypeOf(&RawItem{}).Elem().String()

	if ok := returnType == expectType; ok != true {
		t.Fatalf(fmt.Sprintf("amiibo.NewRawItem(b *[]byte) %s != %s", returnType, expectType))
	}
}
