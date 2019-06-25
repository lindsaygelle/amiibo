package amiibo

import "fmt"

func newRawAmiibo() *RawAmiibo {
	return &RawAmiibo{}
}

// NewRawAmiibo returns a RawAmiibo.
func NewRawAmiibo(amiibo, character, game, head, image, name, tail, t, AU, EU, JP, NA string) *RawAmiibo {
	return &RawAmiibo{
		Amiibo:    amiibo,
		Character: character,
		Game:      game,
		Head:      head,
		Image:     image,
		Name:      name,
		Release:   NewRawRelease(AU, EU, JP, NA),
		Tail:      tail,
		Type:      t}
}

type rawAmiibo interface {
	String() string
}

// RawAmiibo is the struct representation of the Amiibo JSON returned from the Amiibo API endpoint.
type RawAmiibo struct {
	Amiibo    string      `json:"amiiboSeries"`
	Character string      `json:"character"`
	Game      string      `json:"gameSeries"`
	Head      string      `json:"head"`
	Image     string      `json:"image"`
	Name      string      `json:"name"`
	Release   *RawRelease `json:"release"`
	Tail      string      `json:"tail"`
	Type      string      `json:"type"`
}

func (pointer *RawAmiibo) String() string {
	return fmt.Sprintf("%s%s", pointer.Head, pointer.Tail)
}
