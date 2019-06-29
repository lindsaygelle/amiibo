package amiibo

import "fmt"

var (
	_ rawItem = (*RawItem)(nil)
)

func newRawItem() *RawItem {
	return new(RawItem)
}

type rawItem interface {
	String() string
}

// A RawItem type represents a Item JSON object found in the raw Nintendo XHR HTTP response.
type RawItem struct {
	Description  string         `json:"description"`  // null
	LastModified *RawAmiiboUnix `json:"lastModified"` // 1554418285473
	Path         *RawAmiiboURL  `json:"path"`         // "/content/noa/en_US/amiibo/detail/wolf-link-amiibo"
	Title        string         `json:"title"`        // "Wolf Link"
	URL          *RawAmiiboURL  `json:"url"`          // "/amiibo/detail/wolf-link-amiibo"
}

func (r *RawItem) String() string {
	return fmt.Sprintf("%s", r.Title)
}
