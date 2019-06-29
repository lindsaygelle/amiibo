package amiibo

import (
	"fmt"
	"time"
)

func newItem() *Item {
	return new(Item)
}

var (
	_ item = (*Item)(nil)
)

type item interface {
	String() string
}

type Item struct {
	Description string    `json:"description"` // RawItem.Description
	Path        string    `json:"path"`        // RawItem.Path
	Name        string    `json:"name"`        // RawItem.Title
	Timestamp   time.Time `json:"timestamp"`   // RawItem.LastModified
	URL         string    `json:"url"`         // RawItem.URL
}

func (pointer *Item) String() string {
	return fmt.Sprintf("%s", pointer.Name)
}
