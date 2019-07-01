package amiibo

import (
	"github.com/gellel/lexicon"
	"github.com/gellel/slice"
)

var (
	_ rawAmiiboMap = (*RawAmiiboMap)(nil)
)

func getRawAmiiboMap(content *[]byte) *RawAmiiboMap {
	rawPayload, err := unmarshallRawPayload(content)
	if err != nil {
		panic(err)
	}
	rawAmiiboMap := newRawAmiiboMap()
	for _, rawMessage := range rawPayload.AmiiboList {
		rawAmiiboMap.Add(newRawAmiibo(rawMessage))
	}
	return rawAmiiboMap
}

func newRawAmiiboMap() *RawAmiiboMap {
	return &RawAmiiboMap{lexicon: &lexicon.Lexicon{}}
}

type rawAmiiboMap interface{}

type RawAmiiboMap struct {
	lexicon *lexicon.Lexicon
}

func (pointer *RawAmiiboMap) Add(rawAmiibo *RawAmiibo) *RawAmiiboMap {
	pointer.lexicon.Add(rawAmiibo.HexCode, rawAmiibo)
	return pointer
}

func (pointer *RawAmiiboMap) Del(rawAmiibo *RawAmiibo) bool {
	return pointer.lexicon.Del(rawAmiibo.HexCode)
}

func (pointer *RawAmiiboMap) Each(f func(key string, rawAmiibo *RawAmiibo)) *RawAmiiboMap {
	pointer.lexicon.Each(func(key string, value interface{}) {
		f(key, value.(*RawAmiibo))
	})
	return pointer
}

func (pointer *RawAmiiboMap) Fetch(key string) *RawAmiibo {
	rawAmiibo, _ := pointer.Get(key)
	return rawAmiibo
}

func (pointer *RawAmiiboMap) Get(key string) (*RawAmiibo, bool) {
	value, ok := pointer.lexicon.Get(key)
	if ok {
		return value.(*RawAmiibo), ok
	}
	return nil, ok
}

func (pointer *RawAmiiboMap) Has(key string) bool {
	return pointer.lexicon.Has(key)
}

func (pointer *RawAmiiboMap) Intersection(amiiboMap *RawAmiiboMap) *RawAmiiboMap {
	return &RawAmiiboMap{lexicon: pointer.lexicon.Intersection(amiiboMap.lexicon)}
}

func (pointer *RawAmiiboMap) Keys() *slice.String {
	return pointer.lexicon.Keys()
}

func (pointer *RawAmiiboMap) Len() int {
	return pointer.lexicon.Len()
}

func (pointer *RawAmiiboMap) Map(f func(key string, rawAmiibo *RawAmiibo) *RawAmiibo) *RawAmiiboMap {
	pointer.lexicon.Map(func(key string, value interface{}) interface{} {
		return f(key, value.(*RawAmiibo))
	})
	return pointer
}

func (pointer *RawAmiiboMap) Peek(key string) string {
	return pointer.lexicon.Peek(key)
}

func (pointer *RawAmiiboMap) Values() *AmiiboSlice {
	return &AmiiboSlice{slice: pointer.lexicon.Values()}
}
