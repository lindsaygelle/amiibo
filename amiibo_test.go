package amiibo

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAmiibo(t *testing.T) {

	testAmiiboStruct = newAmiibo(rawAmiiboStructDefault)

	if testAmiiboStruct == nil {
		t.Fatalf("amiibo.NewAmiibo(r *amiibo.RawAmiibo) *amiibo.Amiibo == nil")
	}
	returnType := reflect.TypeOf(testAmiiboStruct).Elem().String()

	expectType := reflect.TypeOf(&Amiibo{}).Elem().String()

	if ok := returnType == expectType; ok != true {
		t.Fatalf(fmt.Sprintf("amiibo.NewAmiibo(r *amiibo.RawAmiibo) %s != %s", returnType, expectType))
	}
}
