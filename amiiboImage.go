package amiibo

import "fmt"

// NewAmiiboImage returns a new AmiiboImage pointer.
func NewAmiiboImage(boxArtURL, figureURL *RawAmiiboURL) *AmiiboImage {
	return &AmiiboImage{
		Box:    boxArtURL.String(),
		Figure: figureURL.String()}
}

var (
	_ amiiboImage = (*AmiiboImage)(nil)
)

type amiiboImage interface{}

// An AmiiboImage type represents a normalized set of Amiibo image URLs.
type AmiiboImage struct {
	Box    string `json:"box"`    // RawAmiibo.BoxArtURL
	Figure string `json:"figure"` // RawAmiibo.FigureURL
}

func (pointer *AmiiboImage) String() string {
	return fmt.Sprintf("%v", *pointer)
}
