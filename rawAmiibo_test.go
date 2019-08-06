package amiibo_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gellel/amiibo"
)

func TestRawAmiibo(t *testing.T) {
	b := []byte(rawAmiiboDefault)
	rawAmiibo := amiibo.NewRawAmiibo(&b)

	if ok := reflect.ValueOf(rawAmiibo).Kind() == reflect.Ptr; ok != true {
		t.Fatalf("amiibo.NewRawAmiibo(b *[]byte) *amiibo.RawAmiibo != uintptr")
	}

	returnType := reflect.TypeOf(rawAmiibo).Elem().String()

	expectType := reflect.TypeOf(&amiibo.RawAmiibo{}).Elem().String()

	if ok := returnType == expectType; ok != true {
		t.Fatalf(fmt.Sprintf("amiibo.NewRawAmiibo(b *[]byte) %s != %s", returnType, expectType))
	}
}
