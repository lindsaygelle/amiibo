package amiibo

import (
	"strings"
	"time"
)

// JPNAmiiboComing is a formatted JPNLineupComingItem.
//
// JPNAmiiboComing is currently considered a work in progress and may change as new fields are discovered.
type JPNAmiiboComing struct {
	Description            string
	DetailsPath            string
	DetailsURL             string
	Label                  string
	Name                   string
	NameAlternative        string
	Price                  string
	ReleaseDate            time.Time
	ReleaseDateAlternative string
	Series                 string
	Title                  string
	URL                    string
}

// AddJPNLineupComingItem adds a JPNLineupComingItem to the JPNAmiiboComing.
func (j *JPNAmiiboComing) AddJPNLineupComingItem(v *JPNLineupComingItem) (err error) {
	j.Description = v.Memo
	j.DetailsPath = v.LinkTarget
	j.DetailsURL = v.Link
	j.Label = v.AmiiboLabel
	j.Name = v.Title
	j.NameAlternative = v.TitleRuby
	j.Price = v.Price
	var releaseDate time.Time
	releaseDate, err = time.Parse("2006-02-01", strings.ReplaceAll(v.D, ".", "-"))
	if err == nil {
		j.ReleaseDate = releaseDate
	}
	err = nil
	j.ReleaseDateAlternative = v.ReleaseDateStr
	j.Series = v.AmiiboSeries
	j.Title = v.ThumbVariation
	j.URL = v.AmiiboLink
	return
}
