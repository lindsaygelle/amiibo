package game

import "fmt"

const (
	templateErr string = "game (%s) has a collision against (%s) using key %s" // template error collision hash.
)

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

// Fetch gets an game.Game from the map without worrying about if it is found.
func (m *Map) Fetch(key string) *Game {
	var g, _ = m.Get(key)
	return g
}

// Get gets a game.Game from the map by its key.
func (m *Map) Get(key string) (*Game, bool) {
	var g, ok = (*m)[key]
	return g, ok
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
	var (
		err error
		m   = &Map{}
	)
	for _, game := range game {
		var (
			key = game.Get(k)
		)
		if m.Has(key) && err != nil {
			err = fmt.Errorf(templateErr, game.Name, m.Fetch(key).Name, key)
		}
		m.Add(key, game)
	}
	return m, err
}
