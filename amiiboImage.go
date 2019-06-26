package main

type AmiiboImage struct {
	Box    string `json:"box"`    // RawAmiibo.BoxArtURL
	Figure string `json:"figure"` // RawAmiibo.FigureURL
}
