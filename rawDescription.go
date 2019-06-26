package amiibo

import "fmt"

var (
	_ rawAmiiboDescription = (*RawAmiiboDescription)(nil)
)

type rawAmiiboDescription interface {
	String() string
}

type RawAmiiboDescription string

func (r *RawAmiiboDescription) String() string {
	return fmt.Sprintf("%s", *r)
}
