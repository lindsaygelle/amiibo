package amiibo

import (
	"fmt"

	"github.com/gellel/lexicon"
	"github.com/gellel/slice"
)

func newRawAmiiboMap() *RawAmiiboMap {
	return new(RawAmiiboMap)
}

var (
	_ rawAmiiboMap = (*RawAmiiboMap)(nil)
)

type rawAmiiboMap interface {
	Add(rawAmiibo *RawAmiibo) *RawAmiiboMap
	Del(key string) bool
	Each(f func(key string, rawAmiibo *RawAmiibo)) *RawAmiiboMap
	Empty() bool
	Fetch(key string) *RawAmiibo
	Get(key string) (*RawAmiibo, bool)
	Has(key string) bool
	Keys() *slice.String
	Len() int
	Map(f func(key string, rawAmiibo *RawAmiibo) *RawAmiibo) *RawAmiiboMap
	Merge(a ...*RawAmiiboMap) *RawAmiiboMap
	String() string
	Values() *RawAmiiboSlice
}

// RawAmiiboMap is a map-like object whose methods are used to perform traversal and mutation operations by key-value pair for Raw Amiibo.
type RawAmiiboMap struct {
	lexicon *lexicon.Lexicon
}

// Add method adds one Amiibo to the raw Amiibo map using the key reference and returns the modified raw Amiibo map.
func (pointer *RawAmiiboMap) Add(rawAmiibo *RawAmiibo) *RawAmiiboMap {
	pointer.lexicon.Add(rawAmiibo.AmiiboName.String(), rawAmiibo)
	return pointer
}

// Del method removes a entry from the raw Amiibo map if it exists. Returns a boolean to confirm if it succeeded.
func (pointer *RawAmiiboMap) Del(key string) bool {
	return pointer.lexicon.Del(key)
}

// Each method executes a provided function once for each raw Amiibo map element.
func (pointer *RawAmiiboMap) Each(f func(key string, rawAmiibo *RawAmiibo)) *RawAmiiboMap {
	pointer.lexicon.Each(func(key string, value interface{}) {
		f(key, value.(*RawAmiibo))
	})
	return pointer
}

// Empty returns a boolean indicating whether the raw Amiibo map contains zero values.
func (pointer *RawAmiiboMap) Empty() bool {
	return pointer.lexicon.Empty()
}

// Fetch retrieves the string held by the argument key. Returns nil Amiibo if key does not exist.
func (pointer *RawAmiiboMap) Fetch(key string) *RawAmiibo {
	value, _ := pointer.Get(key)
	return value
}

// Get returns the Amiibo held at the argument key and a boolean indicating if it was successfully retrieved.
func (pointer *RawAmiiboMap) Get(key string) (*RawAmiibo, bool) {
	value, ok := pointer.lexicon.Get(key)
	if ok {
		return value.(*RawAmiibo), ok
	}
	return nil, ok
}

// Has method checks that a given key exists in the raw Amiibo map.
func (pointer *RawAmiiboMap) Has(key string) bool {
	return pointer.lexicon.Has(key)
}

// Keys method returns a slice.String of the raw Amiibo map's own property names, in the same order as we get with a normal loop.
func (pointer *RawAmiiboMap) Keys() *slice.String {
	return pointer.lexicon.Keys()
}

// Len method returns the number of keys in the raw Amiibo map.
func (pointer *RawAmiiboMap) Len() int {
	return pointer.lexicon.Len()
}

// Map method executes a provided function once for each raw Amiibo map element and sets the returned value to the current key.
func (pointer *RawAmiiboMap) Map(f func(key string, rawAmiibo *RawAmiibo) *RawAmiibo) *RawAmiiboMap {
	pointer.lexicon.Map(func(key string, value interface{}) interface{} {
		return f(key, value.(*RawAmiibo))
	})
	return pointer
}

// Merge merges N number of Amiibo maps.
func (pointer *RawAmiiboMap) Merge(m ...*RawAmiiboMap) *RawAmiiboMap {
	for _, m := range m {
		pointer.lexicon.Merge(m.lexicon)
	}
	return pointer
}

func (pointer *RawAmiiboMap) String() string {
	return fmt.Sprintf("%v", pointer.lexicon)
}

// Values method returns a slice.String pointer of the raw Amiibo map's own enumerable property values, in the same order as that provided by a for...in loop.
func (pointer *RawAmiiboMap) Values() *RawAmiiboSlice {
	s := newRawAmiiboSlice()
	pointer.Each(func(_ string, rawAmiibo *RawAmiibo) {
		s.Append(rawAmiibo)
	})
	return s
}
