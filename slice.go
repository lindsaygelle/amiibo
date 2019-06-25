package amiibo

import (
	"fmt"

	"github.com/gellel/slice"
)

var (
	_ s = (*Slice)(nil)
)

func newSlice() *Slice {
	return &Slice{slice: &slice.Slice{}}
}

func NewSlice(amiibo ...*Amiibo) *Slice {
	return newSlice().Assign(amiibo...)
}

func NewSliceFromResponse(r *RawResponse) *Slice {
	slice := newSlice()
	for _, r := range *r.Amiibo {
		slice.Append(NewAmiiboFromRawAmiibo(r))
	}
	return slice
}

type s interface {
	Append(amiibo *Amiibo) *Slice
	Assign(amiibo ...*Amiibo) *Slice
	Bounds(i int) bool
	Concatenate(slice *Slice) *Slice
	Each(f func(i int, amiibo *Amiibo)) *Slice
	Empty() bool
	Fetch(i int) *Amiibo
	Get(i int) (*Amiibo, bool)
	Len() int
	Map(func(i int, amiibo *Amiibo) *Amiibo) *Slice
	Poll() *Amiibo
	Pop() *Amiibo
	Preassign(amiibo ...*Amiibo) *Slice
	Precatenate(slice *Slice) *Slice
	Prepend(amiibo *Amiibo) *Slice
	Push(amiibo *Amiibo) int
	Replace(i int, amiibo *Amiibo) bool
	Slice(start, end int) *Slice
	Splice(start, end int) *Slice
	String() string
}

type Slice struct {
	slice *slice.Slice
}

func (pointer *Slice) Append(amiibo *Amiibo) *Slice {
	pointer.slice.Append(amiibo)
	return pointer
}

func (pointer *Slice) Assign(amiibo ...*Amiibo) *Slice {
	for _, amiibo := range amiibo {
		pointer.Append(amiibo)
	}
	return pointer
}

func (pointer *Slice) Bounds(i int) bool {
	return pointer.slice.Bounds(i)
}

func (pointer *Slice) Concatenate(slice *Slice) *Slice {
	pointer.slice.Concatenate(slice.slice)
	return pointer
}

func (pointer *Slice) Each(f func(i int, amiibo *Amiibo)) *Slice {
	pointer.slice.Each(func(i int, value interface{}) {
		f(i, value.(*Amiibo))
	})
	return pointer
}

func (pointer *Slice) Empty() bool {
	return pointer.slice.Empty()
}

func (pointer *Slice) Fetch(i int) *Amiibo {
	amiibo, _ := pointer.Get(i)
	return amiibo
}

func (pointer *Slice) Get(i int) (*Amiibo, bool) {
	value, ok := pointer.slice.Get(i)
	if ok {
		return value.(*Amiibo), ok
	}
	return nil, ok
}

func (pointer *Slice) Len() int {
	return pointer.slice.Len()
}

func (pointer *Slice) Map(f func(i int, amiibo *Amiibo) *Amiibo) *Slice {
	pointer.slice.Map(func(i int, value interface{}) interface{} {
		return f(i, value.(*Amiibo))
	})
	return pointer
}

func (pointer *Slice) Poll() *Amiibo {
	if value := pointer.slice.Poll(); value != nil {
		return value.(*Amiibo)
	}
	return nil
}

func (pointer *Slice) Pop() *Amiibo {
	if value := pointer.slice.Pop(); value != nil {
		return value.(*Amiibo)
	}
	return nil
}

func (pointer *Slice) Preassign(amiibo ...*Amiibo) *Slice {
	for _, amiibo := range amiibo {
		pointer.Prepend(amiibo)
	}
	return pointer
}

func (pointer *Slice) Precatenate(slice *Slice) *Slice {
	pointer.slice.Precatenate(slice.slice)
	return pointer
}

func (pointer *Slice) Prepend(amiibo *Amiibo) *Slice {
	pointer.slice.Prepend(amiibo)
	return pointer
}

func (pointer *Slice) Push(amiibo *Amiibo) int {
	return pointer.slice.Push(amiibo)
}

func (pointer *Slice) Replace(i int, amiibo *Amiibo) bool {
	return pointer.slice.Replace(i, amiibo)
}

func (pointer *Slice) Slice(start, end int) *Slice {
	return &Slice{slice: pointer.slice.Slice(start, end)}
}

func (pointer *Slice) Splice(start, end int) *Slice {
	return &Slice{slice: pointer.slice.Splice(start, end)}
}

func (pointer *Slice) String() string {
	return fmt.Sprintf("%v", pointer.slice)
}
