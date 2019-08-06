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

	fmt.Println(reflect.TypeOf(rawAmiibo))
}
