package amiibo

import "time"

var (
	_ rawAmiiboReleaseDate = (*RawAmiiboReleaseDate)(nil)
)

type rawAmiiboReleaseDate interface{}

type RawAmiiboReleaseDate string

func (r *RawAmiiboReleaseDate) Time() time.Time {
	return time.Now()
}
