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

type Mix struct {
	Amiibo map[string]*Amiibo
	Games  map[string]*Game
}

func NewMix(c *compatability.XHR, l *lineup.XHR) {
	if c.StatusCode != http.StatusOK {

	}
	if l.StatusCode != http.StatusOK {

	}
}

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
			if _, ok := m[k]; !ok {
				mu.Lock()
				m[k] = &Amiibo{}
				mu.Unlock()
			}
			m[k].Compatability = v
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, v := range i {
			var k = v.Key()
			mu.Lock()
			if _, ok := m[k]; !ok {
				mu.Lock()
				m[k] = &Amiibo{}
				mu.Unlock()
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
			if _, ok := m[k]; !ok {
				mu.Lock()
				m[k] = &Amiibo{}
				mu.Unlock()
			}
			m[k].Lineup = v
		}
	}()
	wg.Wait()
	return m, nil
}

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
			if _, ok := m[k]; !ok {
				mu.Lock()
				m[k] = &Game{}
				mu.Unlock()
			}
			m[k].Game = v
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for _, v := range i {
			var k = v.Key()
			mu.Lock()
			if _, ok := m[k]; !ok {
				mu.Lock()
				m[k] = &Game{}
				mu.Unlock()
			}
			m[k].Item = v
			mu.Unlock()
		}
	}()
	wg.Done()
	return m, nil
}
