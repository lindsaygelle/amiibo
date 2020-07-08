package amiibo_test

import (
	"testing"

	"github.com/lindsaygelle/amiibo"
)

var engLineupAmiiboFileName = "eng-lineup-amiibo-test.json"

func TestENGLineupAmiibo(t *testing.T) {
	var err error
	var ENGLineupAmiibo = amiibo.ENGLineupAmiibo{}
	_, err = amiibo.WriteENGLineupAmiibo(filefolder, engLineupAmiiboFileName, &ENGLineupAmiibo)
	if err != nil {
		t.Fatal(err)
	}
	_, err = amiibo.ReadENGLineupAmiibo(filefolder, engLineupAmiiboFileName)
	if err != nil {
		t.Fatal(err)
	}
}
