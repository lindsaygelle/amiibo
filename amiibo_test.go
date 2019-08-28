package amiibo

import (
	"crypto/md5"
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

func TestAmiibo(t *testing.T) {

	testAmiiboStruct = NewAmiibo(rawAmiiboStructDefault)

	if testAmiiboStruct == nil {
		t.Fatalf("amiibo.NewAmiibo(r *amiibo.RawAmiibo) *amiibo.Amiibo == nil")
	}

	t.Parallel()

	returnType := reflect.TypeOf(testAmiiboStruct).Elem().String()

	expectType := reflect.TypeOf(&Amiibo{}).Elem().String()

	if ok := returnType == expectType; ok != true {
		t.Fatalf("amiibo.NewAmiibo(r *amiibo.RawAmiibo) %s != %s", returnType, expectType)
	}

	if hashMD5 := fmt.Sprintf("%x", md5.Sum([]byte(rawAmiiboStructDefault.AmiiboName))); hashMD5 != testAmiiboStruct.ID {
		t.Fatalf("Amiibo.ID != md5.Sum([]byte(RawAmiibo.AmiiboName); %s != %s", testAmiiboStruct.ID, hashMD5)
	}

	_, file, _, _ := runtime.Caller(0)

	fullpath := filepath.Dir(file)

	if err := WriteAmiibo(fullpath, testAmiiboStruct); err != nil {
		t.Fatalf("amiibo.WriteAmiibo(f string, a *Amiibo) err returned err; %v", err)
	}

	b, err := OpenAmiibo(fullpath, testAmiiboStruct.ID)

	if err != nil {
		t.Fatalf("amiibo.OpenAmiibo(f string, ID string) (*[]byte, error); err != nil; %v", err)
	}

	returnType = reflect.TypeOf(b).Elem().String()

	expectType = reflect.TypeOf(&[]byte{}).Elem().String()

	if ok := returnType == expectType; ok != true {
		t.Fatalf("amiibo.OpenAmiibo(f string, ID string) (*[]byte, error) %s != %s", returnType, expectType)
	}

	if err := DeleteAmiibo(fullpath, testAmiiboStruct); err != nil {
		t.Fatalf("amiibo.DeleteAmiibo(f string, a *Amiibo) err returned err; %v", err)
	}

}
