package amiibo

import (
	"fmt"

	"github.com/gellel/lexicon"
	"github.com/gellel/slice"
)

func newItemMap() *ItemMap {
	return new(ItemMap)
}

var (
	_ itemMap = (*ItemMap)(nil)
)

type itemMap interface {
	Add(item *Item) *ItemMap
	Del(key string) bool
	Each(f func(key string, item *Item)) *ItemMap
	Empty() bool
	Fetch(key string) *Item
	Get(key string) (*Item, bool)
	Has(key string) bool
	Keys() *slice.String
	Len() int
	Map(f func(key string, item *Item) *Item) *ItemMap
	Merge(a ...*ItemMap) *ItemMap
	String() string
	Values() *ItemSlice
}

// ItemMap is a map-like object whose methods are used to perform traversal and mutation operations by key-value pair for Amiibo.
type ItemMap struct {
	lexicon *lexicon.Lexicon
}

// Add method adds one Amiibo to the Amiibo Item map using the key reference and returns the modified Amiibo Item map.
func (pointer *ItemMap) Add(item *Item) *ItemMap {
	pointer.lexicon.Add(item.Name, item)
	return pointer
}

// Del method removes a entry from the Amiibo Item map if it exists. Returns a boolean to confirm if it succeeded.
func (pointer *ItemMap) Del(key string) bool {
	return pointer.lexicon.Del(key)
}

// Each method executes a provided function once for each Amiibo Item map element.
func (pointer *ItemMap) Each(f func(key string, item *Item)) *ItemMap {
	pointer.lexicon.Each(func(key string, value interface{}) {
		f(key, value.(*Item))
	})
	return pointer
}

// Empty returns a boolean indicating whether the Amiibo Item map contains zero values.
func (pointer *ItemMap) Empty() bool {
	return pointer.lexicon.Empty()
}

// Fetch retrieves the string held by the argument key. Returns nil Amiibo if key does not exist.
func (pointer *ItemMap) Fetch(key string) *Item {
	value, _ := pointer.Get(key)
	return value
}

// Get returns the Amiibo held at the argument key and a boolean indicating if it was successfully retrieved.
func (pointer *ItemMap) Get(key string) (*Item, bool) {
	value, ok := pointer.lexicon.Get(key)
	if ok {
		return value.(*Item), ok
	}
	return nil, ok
}

// Has method checks that a given key exists in the Amiibo Item map.
func (pointer *ItemMap) Has(key string) bool {
	return pointer.lexicon.Has(key)
}

// Keys method returns a slice.String of the Amiibo Item map's own property names, in the same order as we get with a normal loop.
func (pointer *ItemMap) Keys() *slice.String {
	return pointer.lexicon.Keys()
}

// Len method returns the number of keys in the Amiibo Item map.
func (pointer *ItemMap) Len() int {
	return pointer.lexicon.Len()
}

// Map method executes a provided function once for each Amiibo Item map element and sets the returned value to the current key.
func (pointer *ItemMap) Map(f func(key string, item *Item) *Item) *ItemMap {
	pointer.lexicon.Map(func(key string, value interface{}) interface{} {
		return f(key, value.(*Item))
	})
	return pointer
}

// Merge merges N number of Amiibo maps.
func (pointer *ItemMap) Merge(m ...*ItemMap) *ItemMap {
	for _, m := range m {
		pointer.lexicon.Merge(m.lexicon)
	}
	return pointer
}

func (pointer *ItemMap) String() string {
	return fmt.Sprintf("%v", pointer.lexicon)
}

// Values method returns a slice.String pointer of the Amiibo Item map's own enumerable property values, in the same order as that provided by a for...in loop.
func (pointer *ItemMap) Values() *ItemSlice {
	s := newItemSlice()
	pointer.Each(func(_ string, item *Item) {
		s.Append(item)
	})
	return s
}
