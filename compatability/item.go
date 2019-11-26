package compatability

import (
	"path/filepath"
	"strings"
)

// Item is a snapshot of a Nintendo video-game product provided from resource.Compatability.
// Item contains data provided as-is from Nintendo with a mixture of content describing
// a Nintendo video-game that is compatabile with an Nintendo Amiibo product.
// Items contain less verbose details than the compatability.Game struct
// but contains details not captured in the aforementioned.
// Items collected from the compatability resource are consumed by the amiibo/game
// package to construct a normalized aggregation of a Nintendo video-game across all resources.
type Item struct {
	Description  string `json:"description"`
	LastModified int64  `json:"lastModified"`
	Path         string `json:"path"`
	Title        string `json:"title"`
	URL          string `json:"url"`
}

// Key returns a reliable ID.
func (i *Item) Key() string {
	return strings.TrimSuffix(filepath.Base(i.URL), ".html")
}
