package amiibo

import "fmt"

const (
	templateErr string = "amiibo (%s) has a collision against (%s) using key %s" // template error collision hash.
)

// Map is a map of amiibo.Amiibo.
type Map map[string]*Amiibo

// Add adds an amiibo.Amiibo to the map by any string as its key.
func (m *Map) Add(key string, a *Amiibo) bool {
	(*m)[key] = a
	return m.Has(key)
}

// Del deletes an ammibo.Amiibo from the map by its key.
func (m *Map) Del(key string) bool {
	delete(*m, key)
	return m.Has(key) == false
}

// Each iterates through the stored amiibo.Amiibo in the map in the same order as in a for-each loop.
func (m *Map) Each(fn func(string, *Amiibo)) {
	for k, v := range *m {
		fn(k, v)
	}
}

// Fetch gets an amiib.Amiibo from the map without worrying about if it is found.
func (m *Map) Fetch(key string) *Amiibo {
	var a, _ = m.Get(key)
	return a
}

// Get gets an amiibo.Amiibo from the map by its key.
func (m *Map) Get(key string) (*Amiibo, bool) {
	var a, ok = (*m)[key]
	return a, ok
}

// Has checks the map for an amiibo.Amiibo by its key.
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
func (m *Map) Val() []*Amiibo {
	var s []*Amiibo
	for _, v := range *m {
		s = append(s, v)
	}
	return s
}

// NewMap creates a new instance of amiibo.Map in O(N) time using the argument
// key as the hashing mechanism. Does not reconcile hash collisions but will
// return a non nil error if an error occurs. Will always return an Amiibo map pointer
// even if there are no Amiibo provided to the function.
func NewMap(k string, amiibo ...*Amiibo) (*Map, error) {
	var (
		err error
		m   = &Map{}
	)
	for _, amiibo := range amiibo {
		var (
			key = amiibo.Get(k)
		)
		if m.Has(key) && err != nil {
			err = fmt.Errorf(templateErr, amiibo.Name, m.Fetch(key).Name, key)
		}
		m.Add(key, amiibo)
	}
	return m, err
}
