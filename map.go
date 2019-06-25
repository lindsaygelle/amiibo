package amiibo

import (
	"fmt"

	"github.com/gellel/lexicon"
	"github.com/gellel/slice"
)

var (
	_ m = (*Map)(nil)
)

func NewMap(amiibo ...*Amiibo) *Map {
	return (&Map{lexicon: lexicon.New()}).Mesh(amiibo...)
}

type m interface {
	Add(amiibo *Amiibo) *Map
	Del(key string) bool
	Contains(amiibo *Amiibo) bool
	Each(f func(amiibo *Amiibo)) *Map
	Empty() bool
	Fetch(key string) *Amiibo
	Get(key string) (*Amiibo, bool)
	Has(key string) bool
	Keys() *slice.String
	Len() int
	Map(f func(amiibo *Amiibo) *Amiibo) *Map
	Merge(m ...*Map) *Map
	Mesh(amiibo ...*Amiibo) *Map
	Values() *Slice
}

type Map struct {
	lexicon *lexicon.Lexicon
}

func (pointer *Map) Add(amiibo *Amiibo) *Map {
	pointer.lexicon.Add(amiibo.ID, amiibo)
	return pointer
}

func (pointer *Map) Del(key string) bool {
	return pointer.lexicon.Del(key)
}

func (pointer *Map) Contains(amiibo *Amiibo) bool {
	return pointer.lexicon.Has(amiibo.Head + amiibo.Tail)
}

func (pointer *Map) Each(f func(amiibo *Amiibo)) *Map {
	pointer.lexicon.Each(func(_ string, i interface{}) {
		f(i.(*Amiibo))
	})
	return pointer
}

func (pointer *Map) Empty() bool {
	return pointer.lexicon.Empty()
}

func (pointer *Map) Fetch(key string) *Amiibo {
	amiibo, _ := pointer.Get(key)
	return amiibo
}

func (pointer *Map) Get(key string) (*Amiibo, bool) {
	value, ok := pointer.lexicon.Get(key)
	if ok {
		return value.(*Amiibo), ok
	}
	return nil, ok
}

func (pointer *Map) Has(key string) bool {
	return pointer.lexicon.Has(key)
}

func (pointer *Map) Keys() *slice.String {
	return pointer.lexicon.Keys()
}

func (pointer *Map) Len() int {
	return pointer.lexicon.Len()
}

func (pointer *Map) Map(f func(amiibo *Amiibo) *Amiibo) *Map {
	pointer.lexicon.Map(func(_ string, i interface{}) interface{} {
		return f(i.(*Amiibo))
	})
	return pointer
}

func (pointer *Map) Merge(m ...*Map) *Map {
	for _, m := range m {
		pointer.lexicon.Merge(m.lexicon)
	}
	return pointer
}

func (pointer *Map) Mesh(amiibo ...*Amiibo) *Map {
	for _, amiibo := range amiibo {
		pointer.Add(amiibo)
	}
	return pointer
}

func (pointer *Map) String() string {
	return fmt.Sprintf("%v", pointer.lexicon)
}

func (pointer *Map) Values() *Slice {
	slice := NewSlice()
	pointer.Each(func(amiibo *Amiibo) {
		slice.Append(amiibo)
	})
	return slice
}
