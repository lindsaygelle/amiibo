package amiibo

import "github.com/gellel/lexicon"

type RawItemMap struct {
	*lexicon.Lexicon
}

func (pointer *RawItemMap) Add(r *RawAmiiboItem) *RawItemMap {
	pointer.Lexicon.Add(r.Title, r)
	return pointer
}

func (pointer *RawItemMap) Each(f func(key string, rawAmiiboItem *RawAmiiboItem)) *RawItemMap {
	pointer.Lexicon.Each(func(key string, value interface{}) {
		f(key, value.(*RawAmiiboItem))
	})
	return pointer
}

func (pointer *RawItemMap) Fetch(key string) *RawAmiiboItem {
	value, _ := pointer.Get(key)
	return value
}

func (pointer *RawItemMap) Get(key string) (*RawAmiiboItem, bool) {
	value, ok := pointer.Lexicon.Get(key)
	if ok {
		return value.(*RawAmiiboItem), ok
	}
	return nil, ok
}

func (pointer *RawItemMap) Map(f func(key string, r *RawAmiiboItem) *RawAmiiboItem) *RawItemMap {
	pointer.Lexicon.Map(func(key string, value interface{}) interface{} {
		return f(key, value.(*RawAmiiboItem))
	})
	return pointer
}
