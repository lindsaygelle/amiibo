package amiibo_test

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/lindsaygelle/amiibo"
)

var jpnChartSoftware amiibo.JPNChartSoftware
var jpnChartSoftwareFilename = "jpn-chart-software.xml"
var jpnChartSoftwareFullpath = filepath.Join(filefolder, jpnChartSoftwareFilename)

func TestGetJPNChartSoftware(t *testing.T) {
	var err error
	if _, err := os.Stat(jpnChartSoftwareFullpath); !os.IsNotExist(err) {
		jpnChartSoftware, err = amiibo.ReadJPNChartSoftware(filefolder, jpnChartSoftwareFilename)
		if err != nil {
			t.Fatal("amiibo.ReadJPNChartSoftware", err)
		}
	} else {
		_, _, jpnChartSoftware, err = amiibo.GetJPNChartSoftware()
		if err != nil {
			t.Fatal(err)
		}
	}
	s, err := amiibo.WriteJPNChartSoftware(filefolder, jpnChartSoftwareFilename, &jpnChartSoftware)
	if err != nil {
		t.Fatal(err)
	}
	if s != jpnChartSoftwareFullpath {
		t.Fatalf("%s != %s", s, jpnChartSoftwareFullpath)
	}
	if l := len(jpnChartSoftware.Items); l == 0 {
		t.Fatal("len: jpnChartSoftware.Items", l)
	}
	if !reflect.ValueOf(jpnChartSoftware).IsZero() {
		testJPNSoftwareMap(t)
	}
}
