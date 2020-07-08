package amiibo_test

import (
	"testing"

	"github.com/lindsaygelle/amiibo"
)

func testJPNSoftwareMap(t *testing.T) {
	var v, err = amiibo.NewJPNSoftwareMap(&jpnChartSoftware)
	if err != nil {
		t.Fatal(err)
	}
	if len(v) == 0 {
		t.Fatal("len(JPNSoftwareMap) == 0")
	}
	if l := len(v); l != len(jpnChartSoftware.Items) {
		t.Logf("JPNSoftwareMap %d jpnChartSoftware.Items %d", l, len(jpnChartSoftware.Items))
	}
	_, err = amiibo.WriteJPNSoftwareMap(filefolder, "jp-software-map.json", &v)
	if err != nil {
		t.Fatal(err)
	}
}
