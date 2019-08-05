package amiibo_test

import (
	"fmt"
	"testing"

	"github.com/gellel/amiibo"
)

func TestRawAmiibo(t *testing.T) {
	b := []byte(rawAmiiboDefault)
	rawAmiibo := amiibo.NewRawAmiibo(&b)
	fmt.Println(rawAmiibo)
}
