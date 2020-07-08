package amiibo_test

import (
	"testing"

	"github.com/lindsaygelle/amiibo"
)

var jpnChartSoftwareItemFileName = "jpn-chart-software-item-test.xml"

func TestJPNChartSoftwareItem(t *testing.T) {
	var err error
	var JPNChartSoftwareItem = amiibo.JPNChartSoftwareItem{}
	_, err = amiibo.WriteJPNChartSoftwareItem(filefolder, jpnChartSoftwareItemFileName, &JPNChartSoftwareItem)
	if err != nil {
		t.Fatal(err)
	}
	_, err = amiibo.ReadJPNChartSoftwareItem(filefolder, jpnChartSoftwareItemFileName)
	if err != nil {
		t.Fatal(err)
	}
}
