package amiibo

import (
	"fmt"

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

// newRawAmiiboMap returns a new empty raw Amiibo map.
func newRawAmiiboMap() *RawAmiiboMap {
	return &RawAmiiboMap{lexicon: &lexicon.Lexicon{}}
}

// rawAmiiboMap defines the required methods for the raw Amiibo map struct.
type rawAmiiboMap interface {
	Add(rawAmiibo *RawAmiibo) *RawAmiiboMap
	Del(rawAmiibo *RawAmiibo) bool
	Each(f func(key string, rawAmiibo *RawAmiibo)) *RawAmiiboMap
	Empty() bool
	Fetch(key string) *RawAmiibo
	Get(key string) (*RawAmiibo, bool)
	Has(key string) bool
	Intersection(rawAmiiboMap *RawAmiiboMap) *RawAmiiboMap
	Keys() *slice.String
	Len() int
	Map(f func(key string, rawAmiibo *RawAmiibo) *RawAmiibo) *RawAmiiboMap
	Peek(key string) string
	String() string
	Values() *RawAmiiboSlice
}

// A RawAmiiboMap is a map-like struct whose methods are used to perform traversal and mutation operations by key-value pair.
// Each raw Amiibo map contains 0 to N number of raw Amiibo, using the raw Amiibo's ID property as the raw Amiibo maps
// key-value pairing mechanism. The raw Amiibo map contains a private Lexicon, with each method performing a mutation
// operation to this property. This struct is protected to prevent incorrect data assignment as the Lexicon permits
// any data interface to be assigned to the raw Amiibo map.
type RawAmiiboMap struct {
	lexicon *lexicon.Lexicon
}

// Add adds a raw Amiibo to the raw Amiibo map and returns the modified map.
func (pointer *RawAmiiboMap) Add(rawAmiibo *RawAmiibo) *RawAmiiboMap {
	pointer.lexicon.Add(rawAmiibo.HexCode, rawAmiibo)
	return pointer
}

// Del deletes a raw Amiibo from the raw Amiibo map and returns the modified map.
func (pointer *RawAmiiboMap) Del(rawAmiibo *RawAmiibo) bool {
	return pointer.lexicon.Del(rawAmiibo.HexCode)
}

// Each method executes a provided function for each raw Amiibo struct in the raw Amiibo map.
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

// Fetch retrieves the raw Amiibo pointer held by the argument key. Returns nil if raw Amiibo does not exist.
func (pointer *RawAmiiboMap) Fetch(key string) *RawAmiibo {
	rawAmiibo, _ := pointer.Get(key)
	return rawAmiibo
}

// Get returns the raw Amiibo pointer held at the argument key and a boolean indicating if it was successfully retrieved.
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

// Intersection returns a new raw Amiibo map containing the shared raw Amibo between the two raw Amiibo maps.
func (pointer *RawAmiiboMap) Intersection(rawAmiiboMap *RawAmiiboMap) *RawAmiiboMap {
	return &RawAmiiboMap{lexicon: pointer.lexicon.Intersection(rawAmiiboMap.lexicon)}
}

// Keys method returns a slice.String of the raw Amiibo map's own property names, in the same order as we get with a normal loop.
func (pointer *RawAmiiboMap) Keys() *slice.String {
	return pointer.lexicon.Keys()
}

// Len method returns the number of keys in the raw Amiibo map.
func (pointer *RawAmiiboMap) Len() int {
	return pointer.lexicon.Len()
}

// Map executes a provided function once for each raw Amiibo pointer and sets the returned raw Amiibo to the current key.
func (pointer *RawAmiiboMap) Map(f func(key string, rawAmiibo *RawAmiibo) *RawAmiibo) *RawAmiiboMap {
	pointer.lexicon.Map(func(key string, value interface{}) interface{} {
		return f(key, value.(*RawAmiibo))
	})
	return pointer
}

// Peek returns the string value of the raw Amiibo assigned to the argument key.
func (pointer *RawAmiiboMap) Peek(key string) string {
	return pointer.lexicon.Peek(key)
}

func (pointer *RawAmiiboMap) String() string {
	return fmt.Sprintf("%v", pointer.lexicon)
}

// Values method returns a raw Amiibo slice pointer of the raw Amiibo maps own enumerable property values.
func (pointer *RawAmiiboMap) Values() *RawAmiiboSlice {
	return &RawAmiiboSlice{slice: pointer.lexicon.Values()}
}
