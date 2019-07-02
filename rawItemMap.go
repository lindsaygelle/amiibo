package amiibo

import "github.com/gellel/lexicon"

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

func newRawItemMap() *RawItemMap {
	return &RawItemMap{lexicon: &lexicon.Lexicon{}}
}

type RawItemMap struct {
	lexicon *lexicon.Lexicon
}

func (pointer *RawItemMap) Add(rawItem *RawItem) *RawItemMap {
	pointer.lexicon.Add(rawItem.Title, rawItem)
	return pointer
}
