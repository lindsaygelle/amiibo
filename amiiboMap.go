package amiibo

import (
	"fmt"

	"github.com/gellel/lexicon"
	"github.com/gellel/slice"
)

func newAmiiboMap() *AmiiboMap {
	return new(AmiiboMap)
}

var (
	_ amiiboMap = (*AmiiboMap)(nil)
)

type amiiboMap interface {
	Add(amiibo *Amiibo) *AmiiboMap
	Del(key string) bool
	Each(f func(key string, amiibo *Amiibo)) *AmiiboMap
	Empty() bool
	Fetch(key string) *Amiibo
	Get(key string) (*Amiibo, bool)
	Has(key string) bool
	Keys() *slice.String
	Len() int
	Map(f func(key string, amiibo *Amiibo) *Amiibo) *AmiiboMap
	Merge(a ...*AmiiboMap) *AmiiboMap
	String() string
	Values() *AmiiboSlice
}

// AmiiboMap is a map-like object whose methods are used to perform traversal and mutation operations by key-value pair for Amiibo.
type AmiiboMap struct {
	lexicon *lexicon.Lexicon
}

// Add method adds one Amiibo to the Amiibo map using the key reference and returns the modified Amiibo map.
func (pointer *AmiiboMap) Add(amiibo *Amiibo) *AmiiboMap {
	pointer.lexicon.Add(amiibo.Name, amiibo)
	return pointer
}

// Del method removes a entry from the Amiibo map if it exists. Returns a boolean to confirm if it succeeded.
func (pointer *AmiiboMap) Del(key string) bool {
	return pointer.lexicon.Del(key)
}

// Each method executes a provided function once for each Amiibo map element.
func (pointer *AmiiboMap) Each(f func(key string, amiibo *Amiibo)) *AmiiboMap {
	pointer.lexicon.Each(func(key string, value interface{}) {
		f(key, value.(*Amiibo))
	})
	return pointer
}

// Empty returns a boolean indicating whether the Amiibo map contains zero values.
func (pointer *AmiiboMap) Empty() bool {
	return pointer.lexicon.Empty()
}

// Fetch retrieves the string held by the argument key. Returns nil Amiibo if key does not exist.
func (pointer *AmiiboMap) Fetch(key string) *Amiibo {
	value, _ := pointer.Get(key)
	return value
}

// Get returns the Amiibo held at the argument key and a boolean indicating if it was successfully retrieved.
func (pointer *AmiiboMap) Get(key string) (*Amiibo, bool) {
	value, ok := pointer.lexicon.Get(key)
	if ok {
		return value.(*Amiibo), ok
	}
	return nil, ok
}

// Has method checks that a given key exists in the Amiibo map.
func (pointer *AmiiboMap) Has(key string) bool {
	return pointer.lexicon.Has(key)
}

// Keys method returns a slice.String of the Amiibo map's own property names, in the same order as we get with a normal loop.
func (pointer *AmiiboMap) Keys() *slice.String {
	return pointer.lexicon.Keys()
}

// Len method returns the number of keys in the Amiibo map.
func (pointer *AmiiboMap) Len() int {
	return pointer.lexicon.Len()
}

// Map method executes a provided function once for each Amiibo map element and sets the returned value to the current key.
func (pointer *AmiiboMap) Map(f func(key string, amiibo *Amiibo) *Amiibo) *AmiiboMap {
	pointer.lexicon.Map(func(key string, value interface{}) interface{} {
		return f(key, value.(*Amiibo))
	})
	return pointer
}

// Merge merges N number of Amiibo maps.
func (pointer *AmiiboMap) Merge(m ...*AmiiboMap) *AmiiboMap {
	for _, m := range m {
		pointer.lexicon.Merge(m.lexicon)
	}
	return pointer
}

func (pointer *AmiiboMap) String() string {
	return fmt.Sprintf("%v", pointer.lexicon)
}

// Values method returns a slice.String pointer of the Amiibo map's own enumerable property values, in the same order as that provided by a for...in loop.
func (pointer *AmiiboMap) Values() *AmiiboSlice {
	s := newAmiiboSlice()
	pointer.Each(func(_ string, amiibo *Amiibo) {
		s.Append(amiibo)
	})
	return s
}
