package amiibo_test

import (
	"testing"

	"github.com/lindsaygelle/amiibo"
)

var engLineupItemFileName = "eng-lineup-item-test.json"

func TestENGLineupItem(t *testing.T) {
	var err error
	var ENGLineupItem = amiibo.ENGLineupItem{}
	_, err = amiibo.WriteENGLineupItem(filefolder, engLineupItemFileName, &ENGLineupItem)
	if err != nil {
		t.Fatal(err)
	}
	_, err = amiibo.ReadENGLineupItem(filefolder, engLineupItemFileName)
	if err != nil {
		t.Fatal(err)
	}
}
