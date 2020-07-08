package amiibo_test

import (
	"testing"

	"github.com/lindsaygelle/amiibo"
)

var jpnAmiiboFileName = "jpn-amiibo-test.json"

func TestJPNAmiibo(t *testing.T) {
	var err error
	var JPNAmiibo = amiibo.JPNAmiibo{}
	_, err = amiibo.WriteJPNAmiibo(filefolder, jpnAmiiboFileName, &JPNAmiibo)
	if err != nil {
		t.Fatal(err)
	}
	_, err = amiibo.ReadJPNAmiibo(filefolder, jpnAmiiboFileName)
	if err != nil {
		t.Fatal(err)
	}
}
