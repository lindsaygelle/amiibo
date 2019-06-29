package amiibo

import (
	"fmt"
	"time"

	"golang.org/x/text/currency"
)

func newAmiibo() *Amiibo {
	return &Amiibo{}
}

// NewAmiibo returns a new Amiibo struct from a RawAmiibo.
func NewAmiibo(r *RawAmiibo) *Amiibo {
	return &Amiibo{
		Available:   r.IsReleased,
		Code:        r.GameCode,
		Description: r.OverviewDescription.String(),
		Franchise:   r.Franchise,
		Hex:         r.HexCode,
		Images:      newAmiiboImage(r.BoxArtURL, r.FigureURL),
		Item:        nil,
		Name:        r.AmiiboName.String(),
		Page:        r.DetailsURL.String(),
		Path:        r.DetailsPath.String(),
		Presenter:   r.PresentedBy.String(),
		Price:       r.Price.Currency(),
		Release:     r.ReleaseDateMask.Time(),
		Series:      r.Series,
		Slug:        r.Slug,
		Timestamp:   r.UnixTimestamp.Time(),
		Type:        r.Type,
		UPC:         r.UPC,
		URL:         r.AmiiboPage.String()}
}

var (
	_ amiibo = (*Amiibo)(nil)
)

type amiibo interface {
	String() string
}

// An Amiibo type representeds a normalized RawAmiibo.
type Amiibo struct {
	Available   bool             `json:"available"`   // RawAmiibo.IsReleased
	Code        string           `json:"code"`        // RawAmiibo.GameCode
	Description string           `json:"description"` // RawAmiibo.OverviewDescription
	Franchise   string           `json:"franchise"`   // RawAmiibo.Franchise
	Hex         string           `json:"hex"`         // RawAmiibo.HexCode
	Images      *AmiiboImage     `json:"images"`      // RawAmiibo.BoxArtURL && RawAmiibo.FigureURL
	Item        *Item            `json:"item"`        //
	Name        string           `json:"name"`        // RawAmiibo.Name
	Page        string           `json:"page"`        // RawAmiibo.DetailsURL
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

func (a *Amiibo) String() string {
	return fmt.Sprintf("%s", a.Name)
}
