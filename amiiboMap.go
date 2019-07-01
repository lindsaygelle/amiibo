package amiibo

import "github.com/gellel/lexicon"

type AmiiboMap struct {
	lexicon *lexicon.Lexicon
}

func (pointer *AmiiboMap) Add(amiibo *Amiibo) *AmiiboMap {
	pointer.lexicon.Add(amiibo.Hex, amiibo)
	return pointer
}

func (pointer *AmiiboMap) Del(amiibo *Amiibo) bool {
	return pointer.lexicon.Del(amiibo.Hex)
}

func (pointer *AmiiboMap) Each(f func(HEX string, amiibo *Amiibo)) *AmiiboMap {
	pointer.lexicon.Each(func(key string, value interface{}) {
		f(key, value.(*Amiibo))
	})
	return pointer
}

func (pointer *AmiiboMap) Fetch(HEX string) *Amiibo {
	amiibo, _ := pointer.Get(HEX)
	return amiibo
}

func (pointer *AmiiboMap) Get(HEX string) (*Amiibo, bool) {
	value, ok := pointer.lexicon.Get(HEX)
	if ok {
		return value.(*Amiibo), ok
	}
	return nil, ok
}

func (pointer *AmiiboMap) Has(HEX string) bool {
	return pointer.lexicon.Has(HEX)
}
