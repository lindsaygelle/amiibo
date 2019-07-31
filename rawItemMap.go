package amiibo

import (
	"fmt"

	"github.com/gellel/lexicon"
	"github.com/gellel/slice"
)

var (
	_ rawItemMap = (*RawItemMap)(nil)
)

// getRawItemMap returns a populated raw Item map from a parsed Nintendo XHR HTTP response.
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

// newRawItem map returns a new raw Item map pointer.
func newRawItemMap() *RawItemMap {
	return &RawItemMap{lexicon: &lexicon.Lexicon{}}
}

// rawItemMap defines the methods for the RawItemMap struct.
type rawItemMap interface {
	Add(rawItem *RawItem) *RawItemMap
	Del(rawItem *RawItem) bool
	Each(f func(key string, rawItem *RawItem)) *RawItemMap
	Empty() bool
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

// A RawItemMap is a map-like struct whose methods are used to perform traversal and mutation operations by key-value pair.
// Each raw Item map contains 0 to N number of raw Item pointers, using the raw Item's title property as the raw Item maps
// key-value pairing mechanism. The raw Item map contains a private Lexicon, with each method performing a mutation
// operation to this property. This struct is protected to prevent incorrect data assignment as the Lexicon permits
// any data interface to be assigned to the raw Item map.
type RawItemMap struct {
	lexicon *lexicon.Lexicon
}

// Add adds a raw Item to the raw Item map and returns the modified map.
func (pointer *RawItemMap) Add(rawItem *RawItem) *RawItemMap {
	pointer.lexicon.Add(rawItem.Title, rawItem)
	return pointer
}

// Del deletes a raw Item pointer from the raw Item map and returns the modified map.
func (pointer *RawItemMap) Del(rawItem *RawItem) bool {
	return pointer.lexicon.Del(rawItem.Title)
}

// Each method executes a provided function for each raw Item pointer in the raw Item map.
func (pointer *RawItemMap) Each(f func(key string, rawItem *RawItem)) *RawItemMap {
	pointer.lexicon.Each(func(key string, value interface{}) {
		f(key, value.(*RawItem))
	})
	return pointer
}

// Empty returns a boolean indicating whether the raw Item map contains zero values.
func (pointer *RawItemMap) Empty() bool {
	return pointer.lexicon.Empty()
}

// Fetch retrieves the raw Item pointer held by the argument key. Returns nil if raw Item does not exist.
func (pointer *RawItemMap) Fetch(key string) *RawItem {
	rawItem, _ := pointer.Get(key)
	return rawItem
}

// Get returns the raw Item pointer held at the argument key and a boolean indicating if it was successfully retrieved.
func (pointer *RawItemMap) Get(key string) (*RawItem, bool) {
	value, ok := pointer.lexicon.Get(key)
	if ok {
		return value.(*RawItem), ok
	}
	return nil, ok
}

// Has method checks that a given key exists in the raw Item map.
func (pointer *RawItemMap) Has(key string) bool {
	return pointer.lexicon.Has(key)
}

// Intersection returns a new raw Item map containing the shared raw Item pointers between the two raw Item maps.
// Mutations made to the raw Item pointers within the new raw Item map are reflected in the receiver and argument
// raw Item map.
func (pointer *RawItemMap) Intersection(rawItemMap *RawItemMap) *RawItemMap {
	return &RawItemMap{lexicon: pointer.lexicon.Intersection(rawItemMap.lexicon)}
}

// Keys method returns a slice.String of the raw Item maps own property names, in the same order as we get with a normal loop.
func (pointer *RawItemMap) Keys() *slice.String {
	return pointer.lexicon.Keys()
}

// Len method returns the number of raw Item pointers in the raw Item map.
func (pointer *RawItemMap) Len() int {
	return pointer.lexicon.Len()
}

// Map executes a provided function once for each raw Item pointer and sets the returned raw Item pointer to the current key.
func (pointer *RawItemMap) Map(f func(key string, rawItem *RawItem) *RawItem) *RawItemMap {
	pointer.lexicon.Map(func(key string, value interface{}) interface{} {
		return f(key, value.(*RawItem))
	})
	return pointer
}

// Peek returns the string value of the raw Item pointer assigned to the argument key.
func (pointer *RawItemMap) Peek(key string) string {
	return pointer.lexicon.Peek(key)
}

// String returns the string value for the raw Item map.
func (pointer *RawItemMap) String() string {
	return fmt.Sprintf("%v", pointer.lexicon)
}

// Values method returns a raw Item slice pointer of the raw Itme maps own enumerable property values.
func (pointer *RawItemMap) Values() *RawItemSlice {
	return &RawItemSlice{slice: pointer.lexicon.Values()}
}
