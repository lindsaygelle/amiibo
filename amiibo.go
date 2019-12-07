package amiibo

import (
	"sync"

	"github.com/gellel/amiibo/amiibo"
	"github.com/gellel/amiibo/compatability"
	"github.com/gellel/amiibo/game"
	"github.com/gellel/amiibo/lineup"
	"github.com/gellel/amiibo/mix"
	"github.com/gellel/amiibo/unmix"
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
// intensive and takes (approx.) a few minutes to complete (on a stable wifi connection).
// If the function cannot reach any of the content destinations or is unable to parse
// some of the Nintendo content, the function may return the corresponding error
// handler that describes why the routine was unable to parse the content.
// To build a custom and more robust implementation of this action (with greater verbosity
// of outcomes per parsing stage) the package level APIs and structs can be used.
// This function is intended to be a basic out of the box utility.
func Scrape() ([]*amiibo.Amiibo, []*game.Game, error) {
	var (
		a    []*amiibo.Amiibo
		c    *compatability.XHR
		cErr error
		g    []*game.Game
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
		return nil, nil, cErr
	}
	if lErr != nil {
		return nil, nil, lErr
	}
	m, mErr = mix.NewMix(c, l)
	if mErr != nil {
		return nil, nil, mErr
	}
	a, g = unmix.Unmix(m)
	return a, g, nil
}
