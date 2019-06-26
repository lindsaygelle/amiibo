package main

var (
	_ rawAmiiboDescription = (*RawAmiiboDescription)(nil)
)

type rawAmiiboDescription interface{}

type RawAmiiboDescription string
