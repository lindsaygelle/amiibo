package main

import (
	"time"

	"golang.org/x/text/currency"
)

var (
	_ amiibo = (*Amiibo)(nil)
)

type amiibo interface{}

// An Amiibo type representeds a normalized RawAmiibo.
type Amiibo struct {
	Available   bool             `json:"available"`   // RawAmiibo.IsReleased
	Code        string           `json:"code"`        // RawAmiibo.GameCode
	Description string           `json:"description"` // RawAmiibo.OverviewDescription
	Franchise   string           `json:"franchise"`   // RawAmiibo.Franchise
	Hex         string           `json:"hex"`         // RawAmiibo.HexCode
	Images      *AmiiboImage     `json:"images"`      // RawAmiibo.BoxArtURL && RawAmiibo.FigurURL
	Item        *AmiiboItem      `json:"item"`        //
	Name        string           `json:"name"`        // RawAmiibo.Name
	Page        string           `json:"page"`        // RawAmiibo.DetailsPage
	Path        string           `json:"path"`        // RawAmiibo.DetailsPath
	Presenter   string           `json:"presenter"`   // RawAmiibo.PresentedBy
	Price       *currency.Amount `json:"price"`       // RawAmiibo.Price
	Release     time.Time        `json:"release"`     // RawAmiibo.ReleaseDateMask
	Series      string           `json:"series"`      // RawAmiibo.Series
	Slug        string           `json:"slug"`        // RawAmiibo.Slug
	Timestamp   time.Time        `json:"timestamp"`   // RawAmiibo.UnixTimestamp
	Type        string           `json:"type"`        // RawAmiibo.Type
	UPC         string           `json:"upc"`         // RawAmiibo.UPC
	URL         string           `json:"url"`         // RawAmiibo.AmiiboPage
}
