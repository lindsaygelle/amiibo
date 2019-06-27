package amiibo

import (
	"fmt"

	"github.com/gellel/slice"
)

func newAmiiboSlice() *AmiiboSlice {
	return &AmiiboSlice{slice: &slice.Slice{}}
}

func NewAmiiboSlice(r ...*RawAmiibo) *AmiiboSlice {
	slice := newAmiiboSlice()
	for _, r := range r {
		slice.Append(NewAmiibo(r))
	}
	return slice
}

func NewAmiiboSliceFromRawSlice(r *RawAmiiboSlice) *AmiiboSlice {
	slice := newAmiiboSlice()
	r.Each(func(i int, r *RawAmiibo) {
		slice.Append(NewAmiibo(r))
	})
	return slice
}

var (
	_ amiiboSlice = (*AmiiboSlice)(nil)
)

type amiiboSlice interface {
	Append(amiibo *Amiibo) *AmiiboSlice
	Assign(amiibo ...*Amiibo) *AmiiboSlice
	Bounds(i int) bool
	Concatenate(amiiboSlice *AmiiboSlice) *AmiiboSlice
	Each(f func(i int, amiibo *Amiibo)) *AmiiboSlice
	Empty() bool
	Fetch(i int) *Amiibo
	Get(i int) (*Amiibo, bool)
	Len() int
	Map(func(i int, amiibo *Amiibo) *Amiibo) *AmiiboSlice
	Poll() *Amiibo
	Pop() *Amiibo
	Preassign(amiibo ...*Amiibo) *AmiiboSlice
	Precatenate(amiiboSlice *AmiiboSlice) *AmiiboSlice
	Prepend(amiibo *Amiibo) *AmiiboSlice
	Push(amiibo *Amiibo) int
	Replace(i int, amiibo *Amiibo) bool
	Slice(start, end int) *AmiiboSlice
	Splice(start, end int) *AmiiboSlice
	String() string
}

type AmiiboSlice struct {
	slice *slice.Slice
}

func (pointer *AmiiboSlice) Append(amiibo *Amiibo) *AmiiboSlice {
	pointer.slice.Append(amiibo)
	return pointer
}

func (pointer *AmiiboSlice) Assign(amiibo ...*Amiibo) *AmiiboSlice {
	for _, amiibo := range amiibo {
		pointer.Append(amiibo)
	}
	return pointer
}

func (pointer *AmiiboSlice) Bounds(i int) bool {
	return pointer.slice.Bounds(i)
}

func (pointer *AmiiboSlice) Concatenate(amiiboSlice *AmiiboSlice) *AmiiboSlice {
	pointer.slice.Concatenate(amiiboSlice.slice)
	return pointer
}

func (pointer *AmiiboSlice) Each(f func(i int, amiibo *Amiibo)) *AmiiboSlice {
	pointer.slice.Each(func(i int, value interface{}) {
		f(i, value.(*Amiibo))
	})
	return pointer
}

func (pointer *AmiiboSlice) Empty() bool {
	return pointer.slice.Empty()
}

func (pointer *AmiiboSlice) Fetch(i int) *Amiibo {
	amiibo, _ := pointer.Get(i)
	return amiibo
}

func (pointer *AmiiboSlice) Get(i int) (*Amiibo, bool) {
	value, ok := pointer.slice.Get(i)
	if ok {
		return value.(*Amiibo), ok
	}
	return nil, ok
}

func (pointer *AmiiboSlice) Len() int {
	return pointer.slice.Len()
}

func (pointer *AmiiboSlice) Map(f func(i int, amiibo *Amiibo) *Amiibo) *AmiiboSlice {
	pointer.slice.Map(func(i int, value interface{}) interface{} {
		return f(i, value.(*Amiibo))
	})
	return pointer
}

func (pointer *AmiiboSlice) Poll() *Amiibo {
	if value := pointer.slice.Poll(); value != nil {
		return value.(*Amiibo)
	}
	return nil
}

func (pointer *AmiiboSlice) Pop() *Amiibo {
	if value := pointer.slice.Pop(); value != nil {
		return value.(*Amiibo)
	}
	return nil
}

func (pointer *AmiiboSlice) Preassign(amiibo ...*Amiibo) *AmiiboSlice {
	for _, amiibo := range amiibo {
		pointer.Prepend(amiibo)
	}
	return pointer
}

func (pointer *AmiiboSlice) Precatenate(amiiboSlice *AmiiboSlice) *AmiiboSlice {
	pointer.slice.Precatenate(amiiboSlice.slice)
	return pointer
}

func (pointer *AmiiboSlice) Prepend(amiibo *Amiibo) *AmiiboSlice {
	pointer.slice.Prepend(amiibo)
	return pointer
}

func (pointer *AmiiboSlice) Push(amiibo *Amiibo) int {
	return pointer.slice.Push(amiibo)
}

func (pointer *AmiiboSlice) Replace(i int, amiibo *Amiibo) bool {
	return pointer.slice.Replace(i, amiibo)
}

func (pointer *AmiiboSlice) Slice(start, end int) *AmiiboSlice {
	return &AmiiboSlice{slice: pointer.slice.Slice(start, end)}
}

func (pointer *AmiiboSlice) Splice(start, end int) *AmiiboSlice {
	return &AmiiboSlice{slice: pointer.slice.Splice(start, end)}
}

func (pointer *AmiiboSlice) String() string {
	return fmt.Sprintf("%v", pointer.slice)
}
