package amiibo

import (
	"github.com/gellel/amiibo/address"
	"github.com/gellel/amiibo/image"
)

// Game is a representation of a Nintendo video-game that is directly compatible
// with a Nintendo Amiibo figurine.
type Game struct {
	Image *image.Image     `json:"image"`
	Name  string           `json:"name"`
	URL   *address.Address `json:"url"`
}
