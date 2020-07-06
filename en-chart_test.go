package amiibo_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/lindsaygelle/amiibo"
)

var engChart amiibo.ENGChart
var engChartFileName = "eng-chart.json"
var engChartFullpath = filepath.Join(filefolder, engChartFileName)

func TestGetENGChart(t *testing.T) {
	var err error
	if _, err := os.Stat(engChartFullpath); !os.IsNotExist(err) {
		engChart, err = amiibo.ReadENGChart(filefolder, engChartFileName)
		if err != nil {
			t.Fatal("amiibo.ReadENGChart", err)
		}
	} else {
		_, _, engChart, err = amiibo.GetENGChart()
		if err != nil {
			t.Fatal(err)
		}
	}
	s, err := amiibo.WriteENGChart(filefolder, engChartFileName, engChart)
	if err != nil {
		t.Fatal(err)
	}
	if s != engChartFullpath {
		t.Fatalf("%s != %s", s, engChartFullpath)
	}
	if l := len(engChart.AmiiboList); l == 0 {
		t.Fatal("len: enLineup.AmiiboList", l)
	}
	if l := len(engChart.GameList); l == 0 {
		t.Fatal("len: enLineup.GameList", l)
	}
	if l := len(engChart.Items); l == 0 {
		t.Fatal("len: enLineup.Items", l)
	}
}
