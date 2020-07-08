package amiibo_test

import (
	"testing"

	"github.com/lindsaygelle/amiibo"
)

var jpnChartItemFileName = "jpn-chart-item-test.json"

func TestJPNChartItem(t *testing.T) {
	var err error
	var JPNChartItem = amiibo.JPNChartItem{}
	_, err = amiibo.WriteJPNChartItem(filefolder, jpnChartItemFileName, &JPNChartItem)
	if err != nil {
		t.Fatal(err)
	}
	_, err = amiibo.ReadJPNChartItem(filefolder, jpnChartItemFileName)
	if err != nil {
		t.Fatal(err)
	}
}
