package main

import "fmt"

var (
	_ rawAmiiboURL = (*RawAmiiboURL)(nil)
)

type rawAmiiboURL interface{}

type RawAmiiboURL string

func (r *RawAmiiboURL) String() string {
	return fmt.Sprintf("%s%s", "https://www.nintendo.com", *r)
}
