package amiibo

import (
	"fmt"

	"github.com/gellel/lexicon"
	"github.com/gellel/slice"
)

func newSet() *Set {
	return &Set{lexicon: &lexicon.Lexicon{}}
}

func NewSet(amiibo ...*Amiibo) *Set {
	return newSet().Assign(amiibo...)
}

func NewSetFromRaw(r *RawSlice) *Set {
	set := newSet()
	for _, r := range *r {
		set.Add(NewAmiiboFromRaw(r))
	}
	return set
}

var (
	_ set = (*Set)(nil)
)

type set interface {
	Add(amiibo *Amiibo) *Set
	Assign(amiibo ...*Amiibo) *Set
	Del(key string) bool
	Each(f func(key string, slice *Slice)) *Set
	Empty() bool
	Fetch(key string) *Slice
	Get(key string) (*Slice, bool)
	Has(key string) bool
	Keys() *slice.String
	Len() int
	Map(f func(key string, slice *Slice) *Slice) *Set
	Size() int
	String() string
	Values() *Slice
}

type Set struct {
	lexicon *lexicon.Lexicon
}

func (pointer *Set) Add(amiibo *Amiibo) *Set {
	if ok := pointer.lexicon.Has(amiibo.Name); ok != true {
		pointer.lexicon.Add(amiibo.Name, NewSlice())
	}
	pointer.Fetch(amiibo.Name).Append(amiibo)
	return pointer
}

func (pointer *Set) Assign(amiibo ...*Amiibo) *Set {
	for _, amiibo := range amiibo {
		pointer.Add(amiibo)
	}
	return pointer
}

func (pointer *Set) Del(key string) bool {
	return pointer.lexicon.Del(key)
}

func (pointer *Set) Each(f func(key string, slice *Slice)) *Set {
	pointer.lexicon.Each(func(key string, value interface{}) {
		f(key, value.(*Slice))
	})
	return pointer
}

func (pointer *Set) Empty() bool {
	return pointer.lexicon.Empty()
}

func (pointer *Set) Fetch(key string) *Slice {
	slice, _ := pointer.Get(key)
	return slice
}

func (pointer *Set) Get(key string) (*Slice, bool) {
	value, ok := pointer.lexicon.Get(key)
	if ok {
		return value.(*Slice), ok
	}
	return nil, ok
}

func (pointer *Set) Has(key string) bool {
	return pointer.lexicon.Has(key)
}

func (pointer *Set) Keys() *slice.String {
	return pointer.lexicon.Keys()
}

func (pointer *Set) Len() int {
	return pointer.lexicon.Len()
}

func (pointer *Set) Map(f func(key string, slice *Slice) *Slice) *Set {
	pointer.lexicon.Map(func(key string, value interface{}) interface{} {
		return f(key, value.(*Slice))
	})
	return pointer
}

func (pointer *Set) Size() int {
	var n int
	pointer.Each(func(_ string, slice *Slice) {
		n = (n + slice.Len())
	})
	return n
}

func (pointer *Set) String() string {
	return fmt.Sprintf("%v", pointer.lexicon)
}

func (pointer *Set) Values() *Slice {
	slice := NewSlice()
	pointer.Each(func(_ string, s *Slice) {
		slice.Concatenate(s)
	})
	return slice
}
