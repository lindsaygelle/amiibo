package amiibo

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRawAmiibo(t *testing.T) {
	b := []byte(rawAmiiboDefault)

	rawAmiibo := NewRawAmiibo(&b)

	if ok := reflect.ValueOf(rawAmiibo).Kind() == reflect.Ptr; ok != true {
		t.Fatalf("amiibo.NewRawAmiibo(b *[]byte) *amiibo.RawAmiibo != uintptr")
	}

	returnType := reflect.TypeOf(rawAmiibo).Elem().String()

	expectType := reflect.TypeOf(&RawAmiibo{}).Elem().String()

	if ok := returnType == expectType; ok != true {
		t.Fatalf(fmt.Sprintf("amiibo.NewRawAmiibo(b *[]byte) %s != %s", returnType, expectType))
	}
}
