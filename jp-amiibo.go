package amiibo

import "fmt"

// JPNAmiibo is a formatted JPN Nintendo Amiibo.
type JPNAmiibo struct{}

func newJPNAmiibo(a JPNChartItem, b JPNLineupItem) (v JPNAmiibo, err error) {
	if a.GetID() != b.GetID() {
		err = fmt.Errorf("JPNAmiibo: ID")
	}
	if err != nil {
		return
	}
	return
}
