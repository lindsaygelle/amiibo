package amiibo_test

import (
	"testing"

	"github.com/gellel/amiibo"
)

func TestGetRawAmiibo(t *testing.T) {

	raw, err := amiibo.GetRawAmiibo()
	if err != nil {
		t.Fatalf("amiibo.GetRawAmiibo() did not successfully fetch raw amiibo " + err.Error())
	}
	if ok := len(raw.Amiibo) > 0; ok != true {
		t.Fatalf("amiibo.GetRawAmiibo() response length if empty")
	}
}
