package amiibo

import "fmt"

func NewRawAmiiboURL(URI string) *RawAmiiboURL {
	r := RawAmiiboURL(URI)
	return &r
}

var (
	_ rawAmiiboURL = (*RawAmiiboURL)(nil)
)

type rawAmiiboURL interface{}

type RawAmiiboURL string

func (r *RawAmiiboURL) String() string {
	return fmt.Sprintf("%s%s", "https://www.nintendo.com", *r)
}
