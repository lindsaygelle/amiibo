package amiibo

import (
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

func TestRawAmiibo(t *testing.T) {
	r := []byte(rawAmiiboDefault)

	rawAmiibo := NewRawAmiibo(&r)

	if ok := reflect.ValueOf(rawAmiibo).Kind() == reflect.Ptr; ok != true {
		t.Fatalf("amiibo.NewRawAmiibo(b *[]byte) *amiibo.RawAmiibo != uintptr")
	}

	returnType := reflect.TypeOf(rawAmiibo).Elem().String()

	expectType := reflect.TypeOf(&RawAmiibo{}).Elem().String()

	if ok := returnType == expectType; ok != true {
		t.Fatalf(fmt.Sprintf("amiibo.NewRawAmiibo(b *[]byte) %s != %s", returnType, expectType))
	}

	_, file, _, _ := runtime.Caller(0)

	fullpath := filepath.Dir(file)

	if err := WriteRawAmiibo(fullpath, rawAmiibo); err != nil {
		t.Fatalf("amiibo.WriteRawAmiibo(f string, r *RawAmiibo) error; err != nil; %v", err)
	}

	b, err := OpenRawAmiibo(fullpath, rawAmiibo.HexCode)

	if err != nil {
		t.Fatalf("amiibo.OpenRawAmiibo(f string, ID string) (*[]byte, error); err != nil; %v", err)
	}

	returnType = reflect.TypeOf(b).Elem().String()

	expectType = reflect.TypeOf(&[]byte{}).Elem().String()

	if ok := returnType == expectType; ok != true {
		t.Fatalf("amiibo.OpenRawAmiibo(f string, ID string) (*[]byte, error); %s != %s", returnType, expectType)
	}

	rawAmiibo, err = GetRawAmiibo(fullpath, rawAmiibo.HexCode)

	if err != nil {
		t.Fatalf("amiibo.GetRawAmiibo(f string, ID string) (*RawAmiibo, error); err != nil; %v", err)
	}

	returnType = reflect.TypeOf(rawAmiibo).Elem().String()

	expectType = reflect.TypeOf(&RawAmiibo{}).Elem().String()

	if ok := returnType == expectType; ok != true {
		t.Fatalf("amiibo.GetRawAmiibo(f string, ID string) (*RawAmiibo, error); %s != %s", returnType, expectType)
	}

	if err := DeleteRawAmiibo(fullpath, rawAmiibo); err != nil {
		t.Fatalf("amiibo.DeleteRawAmiibo(f string, r *RawAmiibo) error; err != err; %v", err)
	}
}
