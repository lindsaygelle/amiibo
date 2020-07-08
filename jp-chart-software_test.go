package amiibo_test

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/lindsaygelle/amiibo"
)

var jpnChartSoftware amiibo.JPNChartSoftware
var jpnChartSoftwareFileName = "jpn-chart-software.xml"
var jpnChartSoftwareFullpath = filepath.Join(filefolder, jpnChartSoftwareFileName)

func TestGetJPNChartSoftware(t *testing.T) {
	var err error
	if _, err := os.Stat(jpnChartSoftwareFullpath); !os.IsNotExist(err) {
		jpnChartSoftware, err = amiibo.ReadJPNChartSoftware(filefolder, jpnChartSoftwareFileName)
		if err != nil {
			t.Fatal("amiibo.ReadJPNChartSoftware", err)
		}
	} else {
		_, _, jpnChartSoftware, err = amiibo.GetJPNChartSoftware()
		if err != nil {
			t.Fatal(err)
		}
	}
	s, err := amiibo.WriteJPNChartSoftware(filefolder, jpnChartSoftwareFileName, &jpnChartSoftware)
	if err != nil {
		t.Fatal(err)
	}
	if s != jpnChartSoftwareFullpath {
		t.Fatalf("%s != %s", s, jpnChartSoftwareFullpath)
	}
	if l := len(jpnChartSoftware.Items); l == 0 {
		t.Fatal("len: jpnChartSoftware.Items", l)
	}
	_, err = amiibo.ReadJPNChartSoftware(filefolder, jpnChartSoftwareFileName)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.ValueOf(jpnChartSoftware).IsZero() {
		testJPNSoftwareMap(t)
	}
}
