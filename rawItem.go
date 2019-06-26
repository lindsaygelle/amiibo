package main

var (
	_ rawAmiiboItem = (*RawAmiiboItem)(nil)
)

type rawAmiiboItem interface{}

type RawAmiiboItem struct {
	Description  string         `json:"description"`
	LastModified *RawAmiiboUnix `json:"lastModified"`
	Path         *RawAmiiboURL  `json:"path"`
	Title        string         `json:"title"`
	URL          *RawAmiiboURL  `json:"url"`
}
