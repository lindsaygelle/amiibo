package amiibo

import (
	"github.com/gellel/lexicon"
	"github.com/gellel/slice"
)

func getRawItemMap(content *[]byte) *RawItemMap {
	rawPayload, err := unmarshallRawPayload(content)
	if err != nil {
		panic(err)
	}
	rawItemMap := newRawItemMap()
	for _, rawMessage := range rawPayload.Items {
		rawItemMap.Add(newRawItem(rawMessage))
	}
	return rawItemMap
}

func newRawItemMap() *RawItemMap {
	return &RawItemMap{lexicon: &lexicon.Lexicon{}}
}

type rawItemMap interface {
	Add(rawItem *RawItem) *RawItemMap
	Del(rawItem *RawItem) bool
	Each(f func(key string, rawItem *RawItem)) *RawItemMap
	Fetch(key string) *RawItem
	Get(key string) (*RawItem, bool)
	Has(key string) bool
	Intersection(rawItemMap *RawItemMap) *RawItemMap
	Keys() *slice.String
	Len() int
	Map(f func(key string, rawItem *RawItem) *RawItem) *RawItemMap
	Peek(key string) string
	String() string
	Values() *RawItemSlice
}

type RawItemMap struct {
	lexicon *lexicon.Lexicon
}

func (pointer *RawItemMap) Add(rawItem *RawItem) *RawItemMap {
	pointer.lexicon.Add(rawItem.Title, rawItem)
	return pointer
}

func (pointer *RawItemMap) Del(rawItem *RawItem) bool {
	return pointer.lexicon.Del(rawItem.Title)
}

func (pointer *RawItemMap) Each(f func(key string, rawItem *RawItem)) *RawItemMap {
	pointer.lexicon.Each(func(key string, value interface{}) {
		f(key, value.(*RawItem))
	})
	return pointer
}

func (pointer *RawItemMap) Fetch(key string) *RawItem {
	rawItem, _ := pointer.Get(key)
	return rawItem
}

func (pointer *RawItemMap) Get(key string) (*RawItem, bool) {
	value, ok := pointer.lexicon.Get(key)
	if ok {
		return value.(*RawItem), ok
	}
	return nil, ok
}

func (pointer *RawItemMap) Has(key string) bool {
	return pointer.lexicon.Has(key)
}

func (pointer *RawItemMap) Intersection(rawItemMap *RawItemMap) *RawItemMap {
	return &RawItemMap{lexicon: pointer.lexicon.Intersection(rawItemMap.lexicon)}
}

func (pointer *RawItemMap) Keys() *slice.String {
	return pointer.lexicon.Keys()
}

func (pointer *RawItemMap) Len() int {
	return pointer.lexicon.Len()
}

func (pointer *RawItemMap) Map(f func(key string, rawItem *RawItem) *RawItem) *RawItemMap {
	pointer.lexicon.Map(func(key string, value interface{}) interface{} {
		return f(key, value.(*RawItem))
	})
	return pointer
}

func (pointer *RawItemMap) Peek(key string) string {
	return pointer.lexicon.Peek(key)
}

func (pointer *RawItemMap) Values() *RawItemSlice {
	slice := newRawItemSlice()
	pointer.Each(func(key string, rawItem *RawItem) {
		slice.Append(rawItem)
	})
	return slice
}
