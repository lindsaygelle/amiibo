package amiibo

import "fmt"

var (
	_ rawAmiiboItem = (*RawAmiiboItem)(nil)
)

func newRawAmiiboItem() {}

type rawAmiiboItem interface {
	String() string
}

// A RawAmiiboItem type represents a Item JSON object found in the raw Nintendo XHR HTTP response.
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
