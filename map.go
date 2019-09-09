package amiibo

type Map map[string]*Amiibo

func (a *Map) Add(amiibo *Amiibo) *Map {
	(*a)[amiibo.ID] = amiibo
	return a
}

func (a *Map) Contains(amiibo *Amiibo) bool {
	return a.Has(amiibo.ID)
}

func (a *Map) Del(ID string) bool {
	delete(*a, ID)
	return (a.Has(ID) == false)
}

func (a *Map) Each(fn func(key string, amiibo *Amiibo)) *Map {
	for key, amiibo := range *a {
		fn(key, amiibo)
	}
	return a
}

func (a *Map) Get(ID string) (*Amiibo, bool) {
	amiibo, ok := (*a)[ID]
	return amiibo, ok
}

func (a *Map) Has(ID string) bool {
	_, ok := a.Get(ID)
	return ok
}
