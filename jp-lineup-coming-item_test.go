package amiibo_test

import (
	"testing"

	"github.com/lindsaygelle/amiibo"
)

var jpnLineupComingItemFileName = "jpn-lineup-coming-item-test.xml"

func TestJPNLineupComingItem(t *testing.T) {
	var err error
	var JPNLineupComingItem = amiibo.JPNLineupComingItem{}
	_, err = amiibo.WriteJPNLineupComingItem(filefolder, jpnLineupComingItemFileName, &JPNLineupComingItem)
	if err != nil {
		t.Fatal(err)
	}
	_, err = amiibo.ReadJPNLineupComingItem(filefolder, jpnLineupComingItemFileName)
	if err != nil {
		t.Fatal(err)
	}
}
