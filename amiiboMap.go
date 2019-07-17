package amiibo

import (
	"fmt"

	"github.com/gellel/lexicon"
	"github.com/gellel/slice"
)

var (
	_ amiiboMap = (*AmiiboMap)(nil)
)

func getAmiiboMap(content *[]byte) *AmiiboMap {
	rawPayload, err := unmarshallRawPayload(content)
	if err != nil {
		panic(err)
	}
	amiiboMap := newAmiiboMap()
	for _, rawMessage := range rawPayload.AmiiboList {
		amiiboMap.Add(newAmiibo(newRawAmiibo(rawMessage)))
	}
	return amiiboMap
}

func newAmiiboMap() *AmiiboMap {
	return &AmiiboMap{lexicon: &lexicon.Lexicon{}}
}

type amiiboMap interface {
	Add(amiibo *Amiibo) *AmiiboMap
	Del(amiibo *Amiibo) bool
	Each(f func(key string, amiibo *Amiibo)) *AmiiboMap
	Empty() bool
	Fetch(key string) *Amiibo
	Get(key string) (*Amiibo, bool)
	Has(key string) bool
	Intersection(amiiboMap *AmiiboMap) *AmiiboMap
	Keys() *slice.String
	Len() int
	Map(f func(key string, amiibo *Amiibo) *Amiibo) *AmiiboMap
	Peek(key string) string
	String() string
	Values() *AmiiboSlice
}

// An AmiiboMap is a map-like struct whose methods are used to peform traversal and mutation operations by key-value pair.
// Each Amiibo Map contains 0 to N number of normalized Amiibo, using the Amiibo's ID field as the Amiibo Map's
// key-value pairing mechanism. The Amiibo Map contains a private Lexicon, with each method performing a mutation
// operation to this property. This struct is protected to prevent incorrect data assignment as the Lexicon permits
// any data interface to be assigned to the Amiibo Map.
type AmiiboMap struct {
	lexicon *lexicon.Lexicon
}

// Add adds an Amiibo to the Amiibo map and returns the modified Amiibo map.
func (pointer *AmiiboMap) Add(amiibo *Amiibo) *AmiiboMap {
	pointer.lexicon.Add(amiibo.ID, amiibo)
	return pointer
}

// Del deletes an Amiibo from the Amiibo map and returns the modified Amiibo map.
func (pointer *AmiiboMap) Del(amiibo *Amiibo) bool {
	return pointer.lexicon.Del(amiibo.ID)
}

// Each method executes a provided function for each Amiibo struct in the Amiibo map.
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

// Fetch retrieves the Amiibo pointer held by the argument key. Returns nil if Amiibo does not exist.
func (pointer *AmiiboMap) Fetch(key string) *Amiibo {
	amiibo, _ := pointer.Get(key)
	return amiibo
}

// Get returns the Amiibo pointer held at the argument key and a boolean indicating if it was successfully retrieved.
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

func (pointer *AmiiboMap) Intersection(amiiboMap *AmiiboMap) *AmiiboMap {
	return &AmiiboMap{lexicon: pointer.lexicon.Intersection(amiiboMap.lexicon)}
}

func (pointer *AmiiboMap) Keys() *slice.String {
	return pointer.lexicon.Keys()
}

func (pointer *AmiiboMap) Len() int {
	return pointer.lexicon.Len()
}

func (pointer *AmiiboMap) Map(f func(key string, amiibo *Amiibo) *Amiibo) *AmiiboMap {
	pointer.lexicon.Map(func(key string, value interface{}) interface{} {
		return f(key, value.(*Amiibo))
	})
	return pointer
}

func (pointer *AmiiboMap) Peek(key string) string {
	return pointer.lexicon.Peek(key)
}

func (pointer *AmiiboMap) String() string {
	return fmt.Sprintf("%v", pointer.lexicon)
}

func (pointer *AmiiboMap) Values() *AmiiboSlice {
	return &AmiiboSlice{slice: pointer.lexicon.Values()}
}
