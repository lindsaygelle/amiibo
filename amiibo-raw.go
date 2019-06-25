package amiibo

import "fmt"

type RawAmiibo struct {
	Amiibo    string      `json:"amiiboSeries"`
	Character string      `json:"character"`
	Game      string      `json:"gameSeries"`
	Head      string      `json:"head"`
	Image     string      `json:"image"`
	Name      string      `json:"name"`
	Relase    *RawRelease `json:"release"`
	Tail      string      `json:"tail"`
	Type      string      `json:"type"`
}

func (pointer *RawAmiibo) String() string {
	return fmt.Sprintf("%s%s", pointer.Head, pointer.Tail)
}
