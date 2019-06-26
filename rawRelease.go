package main

var (
	_ rawAmiiboReleaseDate = (*RawAmiiboReleaseDate)(nil)
)

type rawAmiiboReleaseDate interface{}

type RawAmiiboReleaseDate string
