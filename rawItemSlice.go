package amiibo

import "fmt"

var (
	_ rawAmiiboItemSlice = (*RawAmiiboItemSlice)(nil)
)

type rawAmiiboItemSlice interface {
	String() string
}

type RawAmiiboItemSlice []*RawAmiiboItem

func (r *RawAmiiboItemSlice) String() string {
	return fmt.Sprintf("%v", *r)
}
