package amiibo

import "fmt"

var (
	_ rawAmiiboItem = (*RawAmiiboItem)(nil)
)

type rawAmiiboItem interface {
	String() string
}

type RawAmiiboItem struct {
	Description  string         `json:"description"`
	LastModified *RawAmiiboUnix `json:"lastModified"`
	Path         *RawAmiiboURL  `json:"path"`
	Title        string         `json:"title"`
	URL          *RawAmiiboURL  `json:"url"`
}

func (r *RawAmiiboItem) String() string {
	return fmt.Sprintf("%s", r.Title)
}
