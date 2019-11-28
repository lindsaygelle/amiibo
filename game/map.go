package game

import "fmt"

// Map is a map of game.Game.
type Map map[string]*Game

// Add adds an game.Game to the map by any string as its key.
func (m *Map) Add(key string, a *Game) bool {
	(*m)[key] = a
	return m.Has(key)
}

// Del deletes a game.Game from the map by its key.
func (m *Map) Del(key string) bool {
	delete(*m, key)
	return m.Has(key) == false
}

// Each iterates through the stored game.Game in the map in the same order as in a for-each loop.
func (m *Map) Each(fn func(string, *Game)) {
	for k, v := range *m {
		fn(k, v)
	}
}

// Get gets a game.Game from the map by its key.
func (m *Map) Get(key string) (*Game, bool) {
	var a, ok = (*m)[key]
	return a, ok
}

// Has checks the map for game.Game by its key.
func (m *Map) Has(key string) bool {
	var _, ok = m.Get(key)
	return ok
}

// Keys gets the keys of the map in O(N) time.
func (m *Map) Keys() []string {
	var s []string
	for k := range *m {
		s = append(s, k)
	}
	return s
}

// Len gets the length of the map.
func (m *Map) Len() int {
	return len(*m)
}

// Val gets the values of the map in O(N) time.
func (m *Map) Val() []*Game {
	var s []*Game
	for _, v := range *m {
		s = append(s, v)
	}
	return s
}

// NewMap creates a new instance of game.Map in O(N) time using the argument
// key as the hashing mechanism. Does not reconcile hash collisions but will
// return a non nil error if an error occurs. Will always return an game map pointer
// even if there are no Game provided to the function.
func NewMap(k string, game ...*Game) (*Map, error) {
	const (
		template string = "map has collision using key '%s'"
	)
	var (
		err error
		m   = &Map{}
	)
	for _, game := range game {
		var (
			key = game.Get(k)
		)
		if m.Has(key) && err != nil {
			err = fmt.Errorf(template, k)
		}
		m.Add(key, game)
	}
	return m, err
}
