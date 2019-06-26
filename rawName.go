package main

var (
	_ rawAmiiboName = (*RawAmiiboName)(nil)
)

type rawAmiiboName interface{}

type RawAmiiboName string
