package amiibo

import "github.com/gellel/lexicon"

type amiiboMap interface {
	Add(amiibo *Amiibo) *AmiiboMap
	Each(f func(key string, amiibo *Amiibo)) *AmiiboMap
	Fetch(key string) *Amiibo
	Get(key string) (*Amiibo, bool)
	Map(f func(key string, amiibo *Amiibo) *Amiibo) *AmiiboMap
	Values() *AmiiboSlice
}

type AmiiboMap struct {
	*lexicon.Lexicon
}

func (pointer *AmiiboMap) Add(amiibo *Amiibo) *AmiiboMap {
	pointer.Lexicon.Add(amiibo.Name, amiibo)
	return pointer
}

func (pointer *AmiiboMap) Each(f func(key string, amiibo *Amiibo)) *AmiiboMap {
	pointer.Lexicon.Each(func(key string, value interface{}) {
		f(key, value.(*Amiibo))
	})
	return pointer
}

func (pointer *AmiiboMap) Fetch(key string) *Amiibo {
	value, _ := pointer.Get(key)
	return value
}

func (pointer *AmiiboMap) Get(key string) (*Amiibo, bool) {
	value, ok := pointer.Lexicon.Get(key)
	if ok {
		return value.(*Amiibo), ok
	}
	return nil, ok
}

func (pointer *AmiiboMap) Map(f func(key string, amiibo *Amiibo) *Amiibo) *AmiiboMap {
	pointer.Lexicon.Map(func(key string, value interface{}) interface{} {
		return f(key, value.(*Amiibo))
	})
	return pointer
}

func (pointer *AmiiboMap) Values() *AmiiboSlice {
	slice := NewAmiiboSlice()
	pointer.Each(func(_ string, amiibo *Amiibo) {
		slice.Append(amiibo)
	})
	return slice
}
