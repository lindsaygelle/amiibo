package amiibo

import (
	"fmt"
	"time"
)

func NewRawAmiiboReleaseDate(s string) *RawAmiiboReleaseDate {
	r := RawAmiiboReleaseDate(s)
	return &r
}

var (
	_ rawAmiiboReleaseDate = (*RawAmiiboReleaseDate)(nil)
)

type rawAmiiboReleaseDate interface {
	String() string
	Time() time.Time
}

// A RawAmiiboReleaseDate string represents the unformatted timestamp found in the releaseDateMask property
// that is held by a RawAmiibo within in the Nintendo XHR HTTP response.
type RawAmiiboReleaseDate string

func (r *RawAmiiboReleaseDate) String() string {
	return fmt.Sprintf("%v", *r)
}

func (r *RawAmiiboReleaseDate) Time() time.Time {
	t, _ := time.Parse("01/02/2006", r.String())
	return t
}
