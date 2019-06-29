package amiibo

import "github.com/gellel/lexicon"

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
