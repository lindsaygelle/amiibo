package amiibo

import "fmt"

// JPNAmiibo is a formatted JPN Nintendo Amiibo.
type JPNAmiibo struct{}

func newJPNAmiibo(a JPNChartItem, b JPNLineupItem) (v JPNAmiibo, err error) {
	if len(a.GetID()) == 0 {
		err = fmt.Errorf("JPNAmiibo: len(a.GetID())")
	}
	if err != nil {
		return
	}
	if len(b.GetID()) == 0 {
		err = fmt.Errorf("JPNAmiibo: len(b.GetID())")
	}
	if err != nil {
		return
	}
	if a.GetID() != b.GetID() {
		err = fmt.Errorf("JPNAmiibo: a.GetID() != b.GetID()")
	}
	if err != nil {
		return
	}
	return
}
