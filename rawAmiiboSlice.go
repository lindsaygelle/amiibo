package amiibo

import "fmt"

var (
	_ rawAmiiboSlice = (*RawAmiiboSlice)(nil)
)

type rawAmiiboSlice interface {
	Each(f func(i int, r *RawAmiibo)) *RawAmiiboSlice
	Len() int
	String() string
}

// A RawAmiiboSlice type represents a collection RawAmiibo pointers.
type RawAmiiboSlice []*RawAmiibo

func (r *RawAmiiboSlice) Each(f func(i int, r *RawAmiibo)) *RawAmiiboSlice {
	for i, r := range *r {
		f(i, r)
	}
	return r
}

func (r *RawAmiiboSlice) Len() int {
	return len(*r)
}

func (r *RawAmiiboSlice) String() string {
	return fmt.Sprintf("%v", *r)
}
