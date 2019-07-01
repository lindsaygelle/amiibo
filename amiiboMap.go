package amiibo

import (
	"github.com/gellel/lexicon"
	"github.com/gellel/slice"
)

var (
	_ amiiboMap = (*AmiiboMap)(nil)
)

func newAmiiboMap() *AmiiboMap {
	return &AmiiboMap{lexicon: &lexicon.Lexicon{}}
}

type amiiboMap interface{}

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

func (pointer *AmiiboMap) Intersection(amiiboMap *AmiiboMap) *AmiiboMap {
	return &AmiiboMap{lexicon: pointer.lexicon.Intersection(amiiboMap.lexicon)}
}

func (pointer *AmiiboMap) Keys() *slice.String {
	return pointer.lexicon.Keys()
}

func (pointer *AmiiboMap) Len() int {
	return pointer.lexicon.Len()
}

func (pointer *AmiiboMap) Map(f func(HEX string, amiibo *Amiibo) *Amiibo) *AmiiboMap {
	pointer.lexicon.Map(func(key string, value interface{}) interface{} {
		return f(key, value.(*Amiibo))
	})
	return pointer
}

func (pointer *AmiiboMap) Peek(HEX string) string {
	return pointer.lexicon.Peek(HEX)
}

func (pointer *AmiiboMap) Values() *AmiiboSlice {
	return &AmiiboSlice{slice: pointer.lexicon.Values()}
}
