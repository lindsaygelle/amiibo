package amiibo

import (
	"sync"

	"github.com/gellel/amiibo/mix"
	"github.com/gellel/amiibo/unmix"

	"github.com/gellel/amiibo/compatability"
	"github.com/gellel/amiibo/lineup"
)

const (
	// Version is the semver of the amiibo package.
	Version string = "1.0.0"
)

// Scrape scrapes the Nintendo Amiibo web resources and pushes the content into
// a collection of slices (amiibo and game).
//
// The scrape function processes all of the content found using
// the various tools provided by the amiibo SDK. This process is quite network
// intensive and takes (approx.) two minutes to complete (on a stable wifi connection).
// If the function cannot reach any of the content destinations or is unable to parse
// some of the Nintendo content, the function panics.
// To build a custom and robust implementation of this action,
// the various SDKs APIs and structs can be used.
func Scrape() {
	var (
		c    *compatability.XHR
		cErr error
		m    *mix.Mix
		mErr error
		l    *lineup.XHR
		lErr error
		wg   sync.WaitGroup
	)
	wg.Add(1)
	go func() {
		defer wg.Done()
		c, cErr = compatability.Get()
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		l, lErr = lineup.Get()
	}()
	wg.Wait()
	if cErr != nil {
		panic(cErr)
	}
	if lErr != nil {
		panic(lErr)
	}
	m, mErr = mix.NewMix(c, l)
	if mErr != nil {
		panic(mErr)
	}
	return unmix.Unmix(m)
}
