package amiibo

import (
	"github.com/gellel/slice"
)

func getRawAmiiboSlice() *RawAmiiboSlice {
	content, err := net()
	if err != nil {
		panic(err)
	}
	rawPayload, err := unmarshallRawPayload(content)
	if err != nil {
		panic(err)
	}
	rawAmiiboSlice := newRawAmiiboSlice()
	for _, rawMessage := range rawPayload.AmiiboList {
		rawAmiiboSlice.Append(newRawAmiibo(rawMessage))
	}
	return rawAmiiboSlice
}

func newRawAmiiboSlice() *RawAmiiboSlice {
	return &RawAmiiboSlice{slice: &slice.Slice{}}
}

type RawAmiiboSlice struct {
	slice *slice.Slice
}

func (pointer *RawAmiiboSlice) Append(rawAmiibo *RawAmiibo) *RawAmiiboSlice {
	pointer.slice.Append(rawAmiibo)
	return pointer
}
