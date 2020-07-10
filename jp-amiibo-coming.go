package amiibo

import (
	"strings"
	"time"
)

type JPNAmiiboComing struct {
	Description            string
	Name                   string
	NameAlternative        string
	Price                  string
	ReleaseDate            time.Time
	ReleaseDateAlternative string
	Series                 string
	Title                  string
}

func (j *JPNAmiiboComing) AddJPNLineupComingItem(v *JPNLineupComingItem) (err error) {
	// v.AmiiboLabel
	// v.AmiiboLink
	// v.Link
	// v.LinkTarget
	j.Description = v.Memo

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
	return
}
