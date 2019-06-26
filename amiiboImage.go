package amiibo

import "fmt"

func NewAmiiboImage(boxArtURL, figureURL *RawAmiiboURL) *AmiiboImage {
	return &AmiiboImage{
		Box:    boxArtURL.String(),
		Figure: figureURL.String()}
}

var (
	_ amiiboImage = (*AmiiboImage)(nil)
)

type amiiboImage interface{}

type AmiiboImage struct {
	Box    string `json:"box"`    // RawAmiibo.BoxArtURL
	Figure string `json:"figure"` // RawAmiibo.FigureURL
}

func (pointer *AmiiboImage) String() string {
	return fmt.Sprintf("%s %s", pointer.Box, pointer.Figure)
}
