package amiibo

import (
	"github.com/gellel/lexicon"
	"github.com/gellel/slice"
)

var (
	_ amiiboMap = (*AmiiboMap)(nil)
)

func getAmiiboMap(content *[]byte) *AmiiboMap {
	rawPayload, err := unmarshallRawPayload(content)
	if err != nil {
		panic(err)
	}
	amiiboMap := newAmiiboMap()
	for _, rawMessage := range rawPayload.AmiiboList {
		amiiboMap.Add(newAmiibo(newRawAmiibo(rawMessage)))
	}
	return amiiboMap
}

func newAmiiboMap() *AmiiboMap {
	return &AmiiboMap{lexicon: &lexicon.Lexicon{}}
}

type amiiboMap interface{}

type AmiiboMap struct {
	lexicon *lexicon.Lexicon
}

func (pointer *AmiiboMap) Add(amiibo *Amiibo) *AmiiboMap {
	pointer.lexicon.Add(amiibo.ID, amiibo)
	return pointer
}

func (pointer *AmiiboMap) Del(amiibo *Amiibo) bool {
	return pointer.lexicon.Del(amiibo.ID)
}

func (pointer *AmiiboMap) Each(f func(key string, amiibo *Amiibo)) *AmiiboMap {
	pointer.lexicon.Each(func(key string, value interface{}) {
		f(key, value.(*Amiibo))
	})
	return pointer
}

func (pointer *AmiiboMap) Fetch(key string) *Amiibo {
	amiibo, _ := pointer.Get(key)
	return amiibo
}

func (pointer *AmiiboMap) Get(key string) (*Amiibo, bool) {
	value, ok := pointer.lexicon.Get(key)
	if ok {
		return value.(*Amiibo), ok
	}
	return nil, ok
}

func (pointer *AmiiboMap) Has(key string) bool {
	return pointer.lexicon.Has(key)
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

func (pointer *AmiiboMap) Map(f func(key string, amiibo *Amiibo) *Amiibo) *AmiiboMap {
	pointer.lexicon.Map(func(key string, value interface{}) interface{} {
		return f(key, value.(*Amiibo))
	})
	return pointer
}

func (pointer *AmiiboMap) Peek(key string) string {
	return pointer.lexicon.Peek(key)
}

func (pointer *AmiiboMap) Values() *AmiiboSlice {
	return &AmiiboSlice{slice: pointer.lexicon.Values()}
}
