package amiibo

import "github.com/gellel/slice"

var (
	_ amiiboSlice = (*AmiiboSlice)(nil)
)

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
