package amiibo

import (
	"fmt"
	"time"
)

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

// Time parses the raw Amiibo release date mask into a time.Time struct.
func (r *RawAmiiboReleaseDate) Time() time.Time {
	t, _ := time.Parse("01/02/2006", r.String())
	return t
}
