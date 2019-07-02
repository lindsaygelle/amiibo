package amiibo

import "github.com/gellel/lexicon"

func getItemMap() {}

func newItemMap() *ItemMap {
	return &ItemMap{lexicon: &lexicon.Lexicon{}}
}

type ItemMap struct {
	lexicon *lexicon.Lexicon
}
