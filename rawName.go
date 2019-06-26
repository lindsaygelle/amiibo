package amiibo

import "fmt"

var (
	_ rawAmiiboName = (*RawAmiiboName)(nil)
)

type rawAmiiboName interface{}

type RawAmiiboName string

func (r *RawAmiiboName) String() string {
	return fmt.Sprintf("%s", *r)
}
