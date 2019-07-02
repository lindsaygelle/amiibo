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

type amiiboSlice interface {
	/*Append(amiibo *Amiibo) *AmiiboSlice
	Assign(amiibo ...*Amiibo) *AmiiboSlice
	Bounds(i int) bool
	Concatenate(s *AmiiboSlice) *AmiiboSlice
	Each(f func(i int, amiibo *Amiibo)) *AmiiboSlice
	Empty() bool
	Fetch(i int) *Amiibo
	Get(i int) (*Amiibo, bool)
	Len() int
	Map(func(i int, amiibo *Amiibo) *Amiibo) *AmiiboSlice
	Poll() *Amiibo
	Pop() *Amiibo
	Preassign(amiibo ...*Amiibo) *AmiiboSlice
	Precatenate(s *AmiiboSlice) *AmiiboSlice
	Prepend(amiibo *Amiibo) *AmiiboSlice
	Push(amiibo *Amiibo) int
	Replace(i int, amiibo *Amiibo) bool
	*AmiiboSlice() string*/
}

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
