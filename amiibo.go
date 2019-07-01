package amiibo

import (
	"fmt"
	"html"
	"strings"
	"time"

	"golang.org/x/text/currency"
)

var (
	_ amiibo = (*Amiibo)(nil)
)

func getAmiibo() []*Amiibo {
	r, err := net()
	if err != nil {
		return nil
	}
	p, err := unmarshall(r)
	if err != nil {
		return nil
	}
	amiibos := make([]*Amiibo, len(p.AmiiboList))
	for i, amiibo := range p.AmiiboList {
		amiibos[i] = newAmiibo(amiibo)
	}
	return amiibos
}

func newAmiibo(r *RawAmiibo) *Amiibo {
	var (
		t, _ = time.Parse(timeLayoutRelease, r.ReleaseDateMask)
		desc = reStripSpaces.ReplaceAllString(reStripHTML.ReplaceAllString(r.OverviewDescription, " "), " ")
	)
	return &Amiibo{
		Available:   r.IsReleased,
		Box:         (nintendo + r.BoxArtURL),
		Code:        r.GameCode,
		Description: html.UnescapeString(strings.TrimSpace(desc)),
		Franchise:   r.Franchise,
		Figure:      (nintendo + r.FigureURL),
		Hex:         strings.ToUpper(r.HexCode),
		Name:        (reStripName.ReplaceAllString(r.AmiiboName, "")),
		Page:        (nintendo + r.DetailsURL),
		Path:        r.DetailsPath,
		Presenter:   (strings.Replace(r.PresentedBy, "noa:publisher/", "", -1)),
		Price:       new(currency.Amount),
		Release:     t,
		Series:      r.Series,
		Slug:        r.Slug,
		Timestamp:   (time.Unix(r.UnixTimestamp, 0).UTC()),
		Type:        r.Type,
		UPC:         r.UPC,
		URL:         (nintendo + r.AmiiboPage)}
}

type amiibo interface {
	String() string
}

type Amiibo struct {
	Available   bool             `json:"available"`   // RawAmiibo.IsReleased
	Box         string           `json:"box"`         // RawAmiibo.BoxArtURL
	Code        string           `json:"code"`        // RawAmiibo.GameCode
	Description string           `json:"description"` // RawAmiibo.OverviewDescription
	Figure      string           `json:"figure"`      // RawAmiibo.FigureURL
	Franchise   string           `json:"franchise"`   // RawAmiibo.Franchise
	Hex         string           `json:"hex"`         // RawAmiibo.HexCode
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

func (pointer *Amiibo) String() string {
	return fmt.Sprintf("%s", pointer.Name)
}
