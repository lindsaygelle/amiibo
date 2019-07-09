package amiibo

import (
	"fmt"

	"github.com/gellel/lexicon"
	"github.com/gellel/slice"
)

var (
	_ itemMap = (*ItemMap)(nil)
)

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

func newItemMap() *ItemMap {
	return &ItemMap{lexicon: &lexicon.Lexicon{}}
}

type itemMap interface {
	Add(item *Item) *ItemMap
	Del(item *Item) bool
	Each(f func(key string, item *Item)) *ItemMap
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

type ItemMap struct {
	lexicon *lexicon.Lexicon
}

func (pointer *ItemMap) Add(item *Item) *ItemMap {
	pointer.lexicon.Add(item.ID, item)
	return pointer
}

func (pointer *ItemMap) Del(item *Item) bool {
	return pointer.lexicon.Del(item.ID)
}

func (pointer *ItemMap) Each(f func(key string, item *Item)) *ItemMap {
	pointer.lexicon.Each(func(key string, value interface{}) {
		f(key, value.(*Item))
	})
	return pointer
}

func (pointer *ItemMap) Fetch(key string) *Item {
	item, _ := pointer.Get(key)
	return item
}

func (pointer *ItemMap) Get(key string) (*Item, bool) {
	value, ok := pointer.lexicon.Get(key)
	if ok {
		return value.(*Item), ok
	}
	return nil, ok
}

func (pointer *ItemMap) Has(key string) bool {
	return pointer.lexicon.Has(key)
}

func (pointer *ItemMap) Intersection(itemMap *ItemMap) *ItemMap {
	return &ItemMap{lexicon: pointer.lexicon.Intersection(itemMap.lexicon)}
}

func (pointer *ItemMap) Keys() *slice.String {
	return pointer.lexicon.Keys()
}

func (pointer *ItemMap) Len() int {
	return pointer.lexicon.Len()
}

func (pointer *ItemMap) Map(f func(key string, item *Item) *Item) *ItemMap {
	pointer.lexicon.Map(func(key string, value interface{}) interface{} {
		return f(key, value.(*Item))
	})
	return pointer
}

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
