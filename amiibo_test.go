package amiibo

import (
	"crypto/md5"
	"fmt"
	"reflect"
	"testing"
)

func TestAmiibo(t *testing.T) {

	testAmiiboStruct = newAmiibo(rawAmiiboStructDefault)

	if testAmiiboStruct == nil {
		t.Fatalf("amiibo.NewAmiibo(r *amiibo.RawAmiibo) *amiibo.Amiibo == nil")
	}

	t.Parallel()

	returnType := reflect.TypeOf(testAmiiboStruct).Elem().String()

	expectType := reflect.TypeOf(&Amiibo{}).Elem().String()

	if ok := returnType == expectType; ok != true {
		t.Fatalf(fmt.Sprintf("amiibo.NewAmiibo(r *amiibo.RawAmiibo) %s != %s", returnType, expectType))
	}

	if hashMD5 := fmt.Sprintf("%x", md5.Sum([]byte(rawAmiiboStructDefault.AmiiboName))); hashMD5 != testAmiiboStruct.ID {
		t.Fatalf(fmt.Sprintf("Amiibo.ID != md5.Sum([]byte(RawAmiibo.AmiiboName); %s != %s", testAmiiboStruct.ID, hashMD5))
	}
}
