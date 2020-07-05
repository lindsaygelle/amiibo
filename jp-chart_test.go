package amiibo

import "testing"

func TestChartJPN(t *testing.T) {
	_, _, v, err := GetJPNChart()
	if err != nil {
		t.Fatal(err)
	}
	if l := len(v.Items); l == 0 {
		t.Fatal("len: v.Items", l)
	}
}
