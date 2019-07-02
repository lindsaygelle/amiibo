package amiibo

import "github.com/gellel/slice"

func getItemSlice(content *[]byte) {}

func newItemSlice() *ItemSlice {
	return &ItemSlice{slice: &slice.Slice{}}
}

type ItemSlice struct {
	slice *slice.Slice
}
