package amiibo

import (
	"fmt"
	"time"
)

var (
	_ item = (*Item)(nil)
)

func getItems() []*Item {
	r, err := net()
	if err != nil {
		return nil
	}
	p, err := unmarshall(r)
	if err != nil {
		return nil
	}
	items := make([]*Item, len(p.Items))
	for i, item := range p.Items {
		items[i] = newItem(item)
	}
	return items
}

func newItem(r *RawItem) *Item {
	return &Item{
		Description: r.Description,
		Name:        (reStripName.ReplaceAllString(r.Title, "")),
		Path:        r.Path,
		Timestamp:   time.Unix(r.LastModified, 0),
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
