package amiibo_test

import (
	"testing"

	"github.com/lindsaygelle/amiibo"
)

var engChartItemFileName = "eng-chart-item-test.json"

func TestENGChartItem(t *testing.T) {
	var err error
	var ENGChartItem = amiibo.ENGChartItem{}
	_, err = amiibo.WriteENGChartItem(filefolder, engChartItemFileName, &ENGChartItem)
	if err != nil {
		t.Fatal(err)
	}
	_, err = amiibo.ReadENGChartItem(filefolder, engChartItemFileName)
	if err != nil {
		t.Fatal(err)
	}
}
