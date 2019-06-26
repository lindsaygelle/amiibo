package main

var (
	_ rawAmiiboPrice = (*RawAmiiboPrice)(nil)
)

type rawAmiiboPrice interface{}

type RawAmiiboPrice string
