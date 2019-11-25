package amiibo

import (
	"github.com/gellel/amiibo/address"
	"github.com/gellel/amiibo/image"
)

type Game struct {
	Image *image.Image     `json:"image"`
	Name  string           `json:"name"`
	URL   *address.Address `json:"url"`
}
