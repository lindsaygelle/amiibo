package amiibo

import (
	"fmt"
	"time"
)

type JPNSoftware struct {
	ID              string
	Name            string
	NameAlternative string
	Platform        string
	Price           string
	Priority        string
	ReleaseDate     time.Time
}

func (j *JPNSoftware) Add(v *JPNChartSoftwareItem) (err error) {
	j.ID = v.Code
	j.Name = v.Name
	j.NameAlternative = v.NameKana
	j.Platform = v.Series
	j.Price = v.Price
	j.Priority = v.Priority
	var l = len(v.Date)
	var date = fmt.Sprintf("%s-%s-%s", v.Date[:4], v.Date[l-4:l-2], v.Date[l-2:])
	var releaseDate time.Time
	releaseDate, err = time.Parse("2006-01-02", date)
	if err == nil {
		j.ReleaseDate = releaseDate
	}
	return
}
