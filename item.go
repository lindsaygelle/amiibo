package amiibo

import "time"

func NewAmiiboItem(rawAmiiboItem *RawAmiiboItem) *AmiiboItem {
	return &AmiiboItem{}
}

var (
	_ amiiboItem = (*AmiiboItem)(nil)
)

type amiiboItem interface{}

type AmiiboItem struct {
	Description string    `json:"description"` // RawAmiiboItem.Description
	Path        string    `json:"path"`        // RawAmiiboItem.Path
	Name        string    `json:"name"`        // RawAmiiboItem.Title
	Timestamp   time.Time `json:"timestamp"`   // RawAmiiboItem.LastModified
	URL         string    `json:"url"`         // RawAmiiboItem.URL
}
