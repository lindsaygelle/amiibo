package amiibo

import (
	"github.com/gellel/slice"
)

var (
	_ amiiboSlice = (*AmiiboSlice)(nil)
)

func getAmiiboSlice(content *[]byte) *AmiiboSlice {
	rawPayload, err := unmarshallRawPayload(content)
	if err != nil {
		panic(err)
	}
	amiiboSlice := newAmiiboSlice()
	for _, rawMessage := range rawPayload.AmiiboList {
		r := newRawAmiibo(rawMessage)
		a := newAmiibo(r)
		amiiboSlice.Append(a)
	}
	return amiiboSlice
}

func newAmiiboSlice() *AmiiboSlice {
	return &AmiiboSlice{slice: &slice.Slice{}}
}

type amiiboSlice interface{}

type AmiiboSlice struct {
	slice *slice.Slice
}

func (pointer *AmiiboSlice) Append(amiibo *Amiibo) *AmiiboSlice {
	pointer.slice.Append(amiibo)
	return pointer
}

func (pointer *AmiiboSlice) Each(f func(i int, amiibo *Amiibo)) *AmiiboSlice {
	pointer.slice.Each(func(i int, value interface{}) {
		f(i, value.(*Amiibo))
	})
	return pointer
}
