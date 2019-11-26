package mix

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gellel/amiibo/compatability"
	"github.com/gellel/amiibo/lineup"
)

const (
	tep string = "*%s is %s"
)

var (
	errCNil = fmt.Errorf(tep, "c", "nil")
	errGNil = fmt.Errorf(tep, "g", "nil")
	errINil = fmt.Errorf(tep, "i", "nil")
	errLNil = fmt.Errorf(tep, "l", "nil")
)

// Mix is a an organised collection of maps for both mix.Amiibo and mix.Game.
// Mix is built from the content of compatability.XHR and lineup.XHR.
type Mix struct {
	Amiibo map[string]*Amiibo
	Games  map[string]*Game
}

// Get performs a HTTP request to Nintendo Amiibo compatability resource and the
// Nintendo Amiibo lineup resource and organises the XHR content into its appropriate maps.
// Returns an error if the Nintendo server returns anything other than http.StatusOK for
// both resource destinations. If the response content cannot be
// handled by either compatability.Get or lineup.Get the corresponding error is returned.
func Get() (*Mix, error) {
	var (
		c   *compatability.XHR
		err error
		l   *lineup.XHR
	)
	c, err = compatability.Get()
	if err != nil {
		return nil, err
	}
	l, err = lineup.Get()
	if err != nil {
		return nil, err
	}
	return NewMix(c, l)
}

// NewMix creates a new instance of mix.Mix, combining the various data points across
// compatability.XHR and lineup.XHR into a ordered data structure.
func NewMix(c *compatability.XHR, l *lineup.XHR) (*Mix, error) {
	if c == nil {
		return nil, fmt.Errorf("*c is nil")
	}
	if l == nil {
		return nil, fmt.Errorf("*l is nil")
	}
	if c.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(c.Status)
	}
	if l.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(l.Status)
	}
	var (
		m = &Mix{}
	)
	m.Amiibo, _ = parseAmiibo(c.Amiibo, l.Items, l.Amiibo)
	m.Games, _ = parseGames(c.Games, c.Items)
	return m, nil
}

// parseAmiibo parses the compatability.Amiibo, lineup.Item and lineup.Amiibo
// sequence into a unified map to be consumed by mix.Mix.
func parseAmiibo(c []*compatability.Amiibo, i []*lineup.Item, l []*lineup.Amiibo) (map[string]*Amiibo, error) {
	if len(c) == 0 && len(i) == 0 && len(l) == 0 {
		return nil, fmt.Errorf("*c, *i and *l are empty")
	}
	var (
		mu sync.Mutex
		wg sync.WaitGroup
	)
	var (
		m = map[string]*Amiibo{}
	)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, v := range c {
			var k = v.Key()
			mu.Lock()
			if _, ok := m[k]; !ok {
				m[k] = &Amiibo{}
			}
			m[k].Compatability = v
			mu.Unlock()
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, v := range i {
			var k = v.Key()
			mu.Lock()
			if _, ok := m[k]; !ok {
				m[k] = &Amiibo{}
			}
			m[k].Item = v
			mu.Unlock()
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, v := range l {
			var k = v.Key()
			mu.Lock()
			if _, ok := m[k]; !ok {
				m[k] = &Amiibo{}
			}
			m[k].Lineup = v
			mu.Unlock()
		}
	}()
	wg.Wait()
	return m, nil
}

// parseGame parses the compatability.Game and compatability.Item sequence into a unified map to be
// consumed by mix.Mix.
func parseGames(g []*compatability.Game, i []*compatability.Item) (map[string]*Game, error) {
	if len(g) == 0 && len(i) == 0 {
		return nil, fmt.Errorf("*g and *i are empty")
	}
	var (
		mu sync.Mutex
		wg sync.WaitGroup
	)
	var (
		m = map[string]*Game{}
	)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, v := range g {
			var k = v.Key()
			mu.Lock()
			if _, ok := m[k]; !ok {
				m[k] = &Game{}
			}
			m[k].Game = v
			mu.Unlock()
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, v := range i {
			var k = v.Key()
			mu.Lock()
			if _, ok := m[k]; !ok {
				m[k] = &Game{}
			}
			m[k].Item = v
			mu.Unlock()
		}
	}()
	wg.Wait()
	return m, nil
}
