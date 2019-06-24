package amiibo

import (
	"fmt"
	"time"
)

var (
	_ amiibo = (*Amiibo)(nil)
)

func newAmiibo() *Amiibo {
	return &Amiibo{Release: newRelease()}
}

func NewAmiibo(amiibo, character, game, head, image, name, series, tail, t, URL string, AU, EU, JP, NA time.Time) *Amiibo {
	return &Amiibo{
		Amiibo:    amiibo,
		Character: character,
		Game:      game,
		Head:      head,
		Image:     image,
		Name:      name,
		Release:   NewRelease(AU, EU, JP, NA),
		Series:    series,
		Tail:      tail,
		Type:      t,
		URL:       URL}
}

type amiibo interface{}

type Amiibo struct {
	Amiibo    string   `json:"amiiboSeries"`
	Character string   `json:"character"`
	Game      string   `json:"gameSeries"`
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
