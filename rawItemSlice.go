package main

var (
	_ rawAmiiboItemSlice = (*RawAmiiboItemSlice)(nil)
)

type rawAmiiboItemSlice interface{}

type RawAmiiboItemSlice []*RawAmiiboItem
