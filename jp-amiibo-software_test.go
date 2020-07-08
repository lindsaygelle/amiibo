package amiibo_test

import (
	"testing"

	"github.com/lindsaygelle/amiibo"
)

var jpnAmiiboSoftwareFileName = "jpn-amiibo-software-test.json"

func TestJPNAmiiboSoftware(t *testing.T) {
	var err error
	var JPNAmiiboSoftware = amiibo.JPNAmiiboSoftware{}
	_, err = amiibo.WriteJPNAmiiboSoftware(filefolder, jpnAmiiboSoftwareFileName, &JPNAmiiboSoftware)
	if err != nil {
		t.Fatal(err)
	}
	_, err = amiibo.ReadJPNAmiiboSoftware(filefolder, jpnAmiiboSoftwareFileName)
	if err != nil {
		t.Fatal(err)
	}
}
