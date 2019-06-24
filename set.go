package amiibo

import (
	"github.com/gellel/lexicon"
	"github.com/gellel/slice"
)

type set interface {
	Add(amiibo *Amiibo) *Set
	Del(key string) bool
	Each(f func(key string, slice *Slice)) *Set
	Get(key string) *Slice
	Has(key string) bool
	Keys() *slice.String
	Len() int
	Map(f func(key string, slice *Slice) *Slice) *Set
	Values() *Slice
}

type Set struct {
	lexicon *lexicon.Lexicon
}
