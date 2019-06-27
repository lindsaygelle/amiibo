package amiibo

import "fmt"

var (
	_ rawAmiiboItemSlice = (*RawAmiiboItemSlice)(nil)
)

type rawAmiiboItemSlice interface {
	Each(f func(i int, r *RawAmiiboItem)) *RawAmiiboItemSlice
	Len() int
	String() string
}

type RawAmiiboItemSlice []*RawAmiiboItem

func (r *RawAmiiboItemSlice) Each(f func(i int, r *RawAmiiboItem)) *RawAmiiboItemSlice {
	for i, rawAmiiboItem := range *r {
		f(i, rawAmiiboItem)
	}
	return r
}

func (r *RawAmiiboItemSlice) Len() int {
	return len(*r)
}

func (r *RawAmiiboItemSlice) String() string {
	return fmt.Sprintf("%v", *r)
}
