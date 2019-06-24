package amiibo

import "time"

var (
	_ release = (*Release)(nil)
)

func newRelease() *Release {
	return &Release{}
}

func NewRelease(AU, EU, JP, NA time.Time) *Release {
	return &Release{
		AU: AU,
		EU: EU,
		JP: JP,
		NA: NA}
}

type release interface{}

type Release struct {
	AU time.Time
	EU time.Time
	JP time.Time
	NA time.Time
}
