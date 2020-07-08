package amiibo_test

import (
	"testing"

	"github.com/lindsaygelle/amiibo"
)

var jpnLineupItemFileName = "jpn-lineup-item-test.xml"

func TestJPNLineupItem(t *testing.T) {
	var err error
	var JPNLineupItem = amiibo.JPNLineupItem{}
	_, err = amiibo.WriteJPNLineupItem(filefolder, jpnLineupItemFileName, &JPNLineupItem)
	if err != nil {
		t.Fatal(err)
	}
	_, err = amiibo.ReadJPNLineupItem(filefolder, jpnLineupItemFileName)
	if err != nil {
		t.Fatal(err)
	}
}
