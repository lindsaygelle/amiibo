package amiibo

import "time"

var (
	_ release = (*Release)(nil)
)

func newRelease() *Release {
	return &Release{}
}

func NewRelease(AU, EU, JP, NA time.Time) *Release {
	return &Release{}
}

type release interface{}

type Release struct{}
