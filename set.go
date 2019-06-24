package amiibo

import "github.com/gellel/lexicon"

type set interface {
	Add(amiibo *Amiibo) *Set
	Del(key string) bool
	Get(key string) *Slice
	Has(key string) bool
}

type Set struct {
	lexicon *lexicon.Lexicon
}
