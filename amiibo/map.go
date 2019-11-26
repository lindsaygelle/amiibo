package amiibo

// Map is a map of Amiibo.
type Map map[string]*Amiibo

// Add adds an Amiibo to the map by any string as its key.
func (m *Map) Add(key string, a *Amiibo) bool {
	(*m)[key] = a
	return m.Has(key)
}

// Del deletes an Amiibo from the map by its key.
func (m *Map) Del(key string) bool {
	delete(*m, key)
	return m.Has(key) == false
}

// Each iterates through the Amiibo in the map in the same order as a for-in loop.
func (m *Map) Each(fn func(string, *Amiibo)) {
	for k, v := range *m {
		fn(k, v)
	}
}

// Get gets an Amiibo from the map by its key.
func (m *Map) Get(key string) (*Amiibo, bool) {
	var a, ok = (*m)[key]
	return a, ok
}

// Has checks the map for an Amiibo by its key.
func (m *Map) Has(key string) bool {
	var _, ok = m.Get(key)
	return ok
}

// Keys gets the keys of the map in N time.
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

// Val gets the values of the map in N time.
func (m *Map) Val() []*Amiibo {
	var s []*Amiibo
	for _, v := range *m {
		s = append(s, v)
	}
	return s
}
