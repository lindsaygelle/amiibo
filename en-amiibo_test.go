package amiibo_test

import (
	"testing"

	"github.com/lindsaygelle/amiibo"
)

func TestENGAmiibo(t *testing.T) {
	var err error
	var ENGAmiibo = amiibo.ENGAmiibo{}
	_, err = amiibo.WriteENGAmiibo(filefolder, "en-amiibo-test.json", &ENGAmiibo)
	if err != nil {
		t.Fatal(err)
	}
	_, err = amiibo.ReadENGAmiibo(filefolder, "en-amiibo-test.json")
	if err != nil {
		t.Fatal(err)
	}
}
