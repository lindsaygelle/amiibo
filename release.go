package amiibo

import (
	"fmt"
	"time"
)

var (
	_ release = (*Release)(nil)
)

func newRelease() *Release {
	return &Release{}
}

func NewRelease(AU, EU, JP, NA string) *Release {
	const (
		t string = "2006-01-02T15:04:05.000Z"
		f string = "%sT00:00.000Z"
	)
	var (
		au, _ = time.Parse(t, fmt.Sprintf(f, AU))
		eu, _ = time.Parse(t, fmt.Sprintf(f, EU))
		jp, _ = time.Parse(t, fmt.Sprintf(f, JP))
		na, _ = time.Parse(t, fmt.Sprintf(f, NA))
	)
	return &Release{
		AU: au,
		EU: eu,
		JP: jp,
		NA: na}
}

type release interface{}

type Release struct {
	AU time.Time `json:"au"`
	EU time.Time `json:"eu"`
	JP time.Time `json:"jp"`
	NA time.Time `json:"na"`
}
