package amiibo

import "github.com/gellel/slice"

func getRawItemSlice(content *[]byte) {}

func newRawItemSlice() *RawItemSlice {
	return &RawItemSlice{slice: &slice.Slice{}}
}

type rawItemSlice interface{}

type RawItemSlice struct {
	slice *slice.Slice
}

func (pointer *RawItemSlice) Append(rawItem *RawItem) *RawItemSlice {
	pointer.slice.Append(rawItem)
	return pointer
}
