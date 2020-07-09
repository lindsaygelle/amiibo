package amiibo_test

import (
	"testing"

	"github.com/lindsaygelle/amiibo"
)

var jpnSoftwareFileName = "jpn-software-test.json"

func TestJPNSoftware(t *testing.T) {
	var err error
	var JPNSoftware = amiibo.JPNSoftware{}
	_, err = amiibo.WriteJPNSoftware(filefolder, jpnSoftwareFileName, &JPNSoftware)
	if err != nil {
		t.Fatal(err)
	}
	_, err = amiibo.ReadJPNSoftware(filefolder, jpnSoftwareFileName)
	if err != nil {
		t.Fatal(err)
	}
}
