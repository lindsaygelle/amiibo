package amiibo_test

import (
	"testing"

	"github.com/lindsaygelle/amiibo"
)

var engChartAmiiboFileName = "eng-chart-amiibo-test.json"

func TestENGChartAmiibo(t *testing.T) {
	var err error
	var ENGChartAmiibo = amiibo.ENGChartAmiibo{}
	_, err = amiibo.WriteENGChartAmiibo(filefolder, engChartAmiiboFileName, &ENGChartAmiibo)
	if err != nil {
		t.Fatal(err)
	}
	_, err = amiibo.ReadENGChartAmiibo(filefolder, engChartAmiiboFileName)
	if err != nil {
		t.Fatal(err)
	}
}
