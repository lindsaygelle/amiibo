package amiibo_test

import (
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/lindsaygelle/amiibo"
)

var jpnChart amiibo.JPNChart
var jpnChartFileName = "jpn-chart.xml"
var jpnChartFullpath = filepath.Join(filefolder, jpnChartFileName)

func TestGetJPNChart(t *testing.T) {
	var err error
	if _, err := os.Stat(jpnChartFullpath); !os.IsNotExist(err) {
		jpnChart, err = amiibo.ReadJPNChart(filefolder, jpnChartFileName)
		if err != nil {
			log.Println("amiibo.ReadJPNChart", err)
		}
	} else {
		_, _, jpnChart, err = amiibo.GetJPNChart()
		if err != nil {
			t.Fatal(err)
		}
	}
	s, err := amiibo.WriteJPNChart(filefolder, jpnChartFileName, jpnChart)
	if err != nil {
		t.Fatal(err)
	}
	if s != jpnChartFullpath {
		t.Fatalf("%s != %s", s, jpnChartFullpath)
	}
	if l := len(jpnChart.Items); l == 0 {
		t.Fatal("len: jpnChart.Items", l)
	}
}
