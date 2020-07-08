package amiibo_test

import (
	"testing"

	"github.com/lindsaygelle/amiibo"
)

var jpnLineupSeriesItemFileName = "jpn-lineup-series-item-test.xml"

func TestJPNLineupSeriesItem(t *testing.T) {
	var err error
	var JPNLineupSeriesItem = amiibo.JPNLineupSeriesItem{}
	_, err = amiibo.WriteJPNLineupSeriesItem(filefolder, jpnLineupSeriesItemFileName, &JPNLineupSeriesItem)
	if err != nil {
		t.Fatal(err)
	}
	_, err = amiibo.ReadJPNLineupSeriesItem(filefolder, jpnLineupSeriesItemFileName)
	if err != nil {
		t.Fatal(err)
	}
}
