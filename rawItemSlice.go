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
