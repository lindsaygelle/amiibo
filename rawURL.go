package amiibo

import "fmt"

// NewRawAmiiboURL returns a new RawAmiiboURL string pointer.
func NewRawAmiiboURL(URI string) *RawAmiiboURL {
	r := RawAmiiboURL(URI)
	return &r
}

var (
	_ rawAmiiboURL = (*RawAmiiboURL)(nil)
)

type rawAmiiboURL interface {
	String() string
}

// A RawAmiiboURL string represents a Nintendo URI fragment found in a RawAmiibo or
// RawAmiiboItem within in the Nintendo XHR HTTP response.
type RawAmiiboURL string

func (r *RawAmiiboURL) String() string {
	return fmt.Sprintf("%s%s", "https://www.nintendo.com", *r)
}
