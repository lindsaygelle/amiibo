package amiibo

import "fmt"

func NewRawAmiiboPresentedBy(s string) *RawAmiiboPresentedBy {
	r := RawAmiiboPresentedBy(s)
	return &r
}

var (
	_ rawAmiiboPresentedBy = (*RawAmiiboPresentedBy)(nil)
)

type rawAmiiboPresentedBy interface{}

type RawAmiiboPresentedBy string

func (r *RawAmiiboPresentedBy) String() string {
	return fmt.Sprintf("%s", *r)
}
