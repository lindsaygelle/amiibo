package amiibo

import (
	"fmt"
)

var (
	_ amiibo = (*Amiibo)(nil)
)

func NewAmiibo(character, game, head, image, name, series, tail, t, AU, EU, JP, NA string) *Amiibo {
	return &Amiibo{
		Character: character,
		Game:      game,
		Head:      head,
		ID:        head + tail,
		Image:     image,
		Name:      name,
		Release:   NewRelease(AU, EU, JP, NA),
		Series:    series,
		Tail:      tail,
		Type:      t,
		URL:       "https://www.amiiboapi.com/api/amiibo/?id=" + head + tail}
}

type amiibo interface {
	String() string
}

type Amiibo struct {
	Character string   `json:"character"`
	Game      string   `json:"game"`
	Head      string   `json:"head"`
	ID        string   `json:"ID"`
	Image     string   `json:"image"`
	Name      string   `json:"name"`
	Release   *Release `json:"release"`
	Series    string   `json:"series"`
	Tail      string   `json:"tail"`
	Type      string   `json:"type"`
	URL       string   `json:"URL"`
}

func (pointer *Amiibo) String() string {
	return fmt.Sprintf("%s", pointer.Name)
}
