package amiibo

import (
	"encoding/json"
	"fmt"
	"time"
)

var (
	_ item = (*Item)(nil)
)

func deleteItem() bool {
	return false
}

func getItem() *Item {
	return nil
}

func writeItem(item *Item) bool {
	return false
}

func unmarshallItem(content *[]byte) (*Item, error) {
	r := &Item{}
	err := json.Unmarshal(*content, r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func newItem(r *RawItem) *Item {
	return &Item{
		Description: r.Description,
		Name:        (reStripName.ReplaceAllString(r.Title, "")),
		Path:        r.Path,
		Timestamp:   (time.Unix(r.LastModified, 0).UTC()),
		URL:         (nintendo + r.URL)}
}

type item interface {
	String() string
}

type Item struct {
	Description string    `json:"description"` // RawItem.Description
	Name        string    `json:"name"`        // RawItem.Title
	Path        string    `json:"path"`        // RawItem.Path
	Timestamp   time.Time `json:"timestamp"`   // RawItem.LastModified
	URL         string    `json:"url"`         // RawItem.URL
}

func (pointer *Item) String() string {
	return fmt.Sprintf("%s", pointer.Name)
}
