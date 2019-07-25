package amiibo

import (
	"fmt"

	"github.com/gellel/lexicon"
	"github.com/gellel/slice"
)

var (
	_ itemMap = (*ItemMap)(nil)
)

// NewItemMap returns a new Item map pointer. A Item map pointer can be built
// from a cached XHR payload or directly from the Nintendo Amiibo source. To create from source
// parse in the optional byte code pointer, otherwise leave empty and it will be collected from
// the Nintendo XHR HTTP response.
func NewItemMap(b ...byte) {}

// getItemMap returns a populated Item map from a parsed Nintendo XHR HTTP response.
func getItemMap(content *[]byte) *ItemMap {
	rawPayload, err := unmarshallRawPayload(content)
	if err != nil {
		panic(err)
	}
	itemMap := newItemMap()
	for _, rawMessage := range rawPayload.Items {
		r := newRawItem(rawMessage)
		i := newItem(r)
		itemMap.Add(i)
	}
	return itemMap
}

// newItemMap returns a new empty Item map pointer.
func newItemMap() *ItemMap {
	return &ItemMap{lexicon: &lexicon.Lexicon{}}
}

// itemMap defines the required methods for the Item map struct.
type itemMap interface {
	Add(item *Item) *ItemMap
	Del(item *Item) bool
	Each(f func(key string, item *Item)) *ItemMap
	Empty() bool
	Fetch(key string) *Item
	Get(key string) (*Item, bool)
	Has(key string) bool
	Intersection(itemMap *ItemMap) *ItemMap
	Keys() *slice.String
	Len() int
	Map(f func(key string, item *Item) *Item) *ItemMap
	Peek(key string) string
	String() string
	Values() *ItemSlice
}

// An ItemMap is a map-like struct whose methods are used to peform traversal and mutation operations by key-value pair.
// Each Item map contains 0 to N number of normalized Item pointers, using the Item's ID field as the Item map's
// key-value pairing mechanism. The Item map contains a private Lexicon, with each method performing a mutation
// operation to this property. This struct is protected to prevent incorrect data assignment as the Lexicon permits
// any data interface to be assigned to the Item map.
type ItemMap struct {
	lexicon *lexicon.Lexicon
}

// Add adds a Item pointer to the Item map and returns the modified map.
func (pointer *ItemMap) Add(item *Item) *ItemMap {
	pointer.lexicon.Add(item.ID, item)
	return pointer
}

// Del deletes a Item pointer from the Item map and returns the modified map.
func (pointer *ItemMap) Del(item *Item) bool {
	return pointer.lexicon.Del(item.ID)
}

// Each method executes a provided function for each Item struct in the Item map.
func (pointer *ItemMap) Each(f func(key string, item *Item)) *ItemMap {
	pointer.lexicon.Each(func(key string, value interface{}) {
		f(key, value.(*Item))
	})
	return pointer
}

// Empty returns a boolean indicating whether the raw Amiibo map contains zero values.
func (pointer *ItemMap) Empty() bool {
	return pointer.lexicon.Empty()
}

// Fetch retrieves the Item pointer held by the argument key. Returns nil if Item does not exist.
func (pointer *ItemMap) Fetch(key string) *Item {
	item, _ := pointer.Get(key)
	return item
}

// Get returns the Item pointer held at the argument key and a boolean indicating if it was successfully retrieved.
func (pointer *ItemMap) Get(key string) (*Item, bool) {
	value, ok := pointer.lexicon.Get(key)
	if ok {
		return value.(*Item), ok
	}
	return nil, ok
}

// Has method checks that a given key exists in the Item map.
func (pointer *ItemMap) Has(key string) bool {
	return pointer.lexicon.Has(key)
}

// Intersection returns a new Item map containing the shared Item pointers between the two Item maps.
func (pointer *ItemMap) Intersection(itemMap *ItemMap) *ItemMap {
	return &ItemMap{lexicon: pointer.lexicon.Intersection(itemMap.lexicon)}
}

// Keys method returns a slice.String of the Item map's own property names, in the same order as we get with a normal loop.
func (pointer *ItemMap) Keys() *slice.String {
	return pointer.lexicon.Keys()
}

// Len method returns the number of Item pointer in the Item map.
func (pointer *ItemMap) Len() int {
	return pointer.lexicon.Len()
}

// Map executes a provided function once for each Item pointer and sets the returned Item pointer to the current key.
func (pointer *ItemMap) Map(f func(key string, item *Item) *Item) *ItemMap {
	pointer.lexicon.Map(func(key string, value interface{}) interface{} {
		return f(key, value.(*Item))
	})
	return pointer
}

// Peek returns the string value of the Item pointer assigned to the argument key.
func (pointer *ItemMap) Peek(key string) string {
	return pointer.lexicon.Peek(key)
}

func (pointer *ItemMap) String() string {
	return fmt.Sprintf("%v", pointer.lexicon)
}

func (pointer *ItemMap) Values() *ItemSlice {
	slice := newItemSlice()
	pointer.Each(func(key string, item *Item) {
		slice.Append(item)
	})
	return slice
}
