package amiibo

import "fmt"

var (
	_ rawItem = (*RawItem)(nil)
)

func deleteRawItem() bool {
	return false
}

func getRawItem() *RawItem {
	return nil
}

func writeRawItem(rawItem *RawItem) bool {
	return false
}

type rawItem interface {
	String() string
}

type RawItem struct {
	Description  string `json:"description"`  // null
	LastModified int64  `json:"lastModified"` // 1554418285473
	Path         string `json:"path"`         // "/content/noa/en_US/amiibo/detail/wolf-link-amiibo"
	Title        string `json:"title"`        // "Wolf Link"
	URL          string `json:"url"`          // "/amiibo/detail/wolf-link-amiibo"
}

func (pointer *RawItem) String() string {
	return fmt.Sprintf("%s", pointer.Title)
}
