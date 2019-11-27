package game

import (
	"fmt"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"

	"github.com/gellel/amiibo/address"
	"github.com/gellel/amiibo/image"
	"github.com/gellel/amiibo/resource"
	t "github.com/gellel/amiibo/text"
)

// Amiibo is a representation of a Nintendo Amiibo figurine that is directly compatible
// with a Nintendo video game.
type Amiibo struct {
	Image           *image.Image     `json:"image"`
	IsReleased      bool             `json:"is_released"`
	Name            string           `json:"name"`
	ReleaseDateMask string           `json:"release_date_mask"`
	Series          string           `json:"series"`
	Timestamp       time.Time        `json:"timestamp"`
	URL             *address.Address `json:"url"`
}

// NewAmiibo creates a new instance of the game.Amiibo.
func NewAmiibo(s *goquery.Selection) (*Amiibo, error) {
	var (
		ok bool
	)
	ok = (s != nil)
	if !ok {
		return nil, fmt.Errorf("*s is nil")
	}
	ok = (s.Length() != 0)
	if !ok {
		return nil, fmt.Errorf("*s is empty")
	}
	var (
		amiibo = Amiibo{}
	)
	amiibo.Image, _ = parseAmiiboImage(s)
	amiibo.Name, _ = parseAmiiboName(s)
	amiibo.Series, _ = parseAmiiboSeries(s)
	amiibo.URL, _ = parseAmiiboURL(s)
	amiibo.ReleaseDateMask, _ = parseAmiiboReleaseDateMask(s)
	amiibo.Timestamp, _ = time.Parse("01/02/2006", amiibo.ReleaseDateMask)
	amiibo.Timestamp = amiibo.Timestamp.UTC()
	amiibo.IsReleased = amiibo.Timestamp.Unix() < time.Now().UTC().Unix()
	return &amiibo, nil
}

func parseAmiiboImage(s *goquery.Selection) (*image.Image, error) {
	const (
		CSS string = "img"
	)
	var (
		ok     bool
		rawurl string
	)
	s = (s.Find(CSS).First())
	ok = (s.Length() != 0)
	if !ok {
		return nil, fmt.Errorf("*s is empty")
	}
	rawurl, ok = s.Attr("src")
	if !ok {
		return nil, fmt.Errorf("*s has no src")
	}
	return image.NewImage(fmt.Sprintf("%s%s", resource.Nintendo, rawurl))
}

func parseAmiiboName(s *goquery.Selection) (string, error) {
	const (
		CSS string = ".amiibo-name"
	)
	var (
		err  error
		name string
		ok   bool
	)
	s = (s.Find(CSS).First())
	ok = (s.Length() != 0)
	if !ok {
		return name, fmt.Errorf("*s is empty")
	}
	name = (s.Text())
	ok = (len(name) != 0)
	if !ok {
		return name, fmt.Errorf("*s has no text")
	}
	return t.Name(name), err
}

func parseAmiiboReleaseDateMask(s *goquery.Selection) (string, error) {
	const (
		CSS string = "span[itemprop='releaseDate']"
	)
	var (
		err       error
		ok        bool
		substring = (s.Find(CSS).Text())
	)
	substring = strings.TrimSpace(substring)
	ok = (len(substring) != 0)
	if !ok {
		return substring, fmt.Errorf("*s is empty")
	}
	substring = strings.ToLower(substring)
	substring = strings.ReplaceAll(substring, " ", "")
	substring = strings.ReplaceAll(substring, "\n", "")
	substring = strings.Replace(substring, "available", "", 1)
	return substring, err
}

func parseAmiiboSeries(s *goquery.Selection) (string, error) {
	const (
		CSS string = "span[itemprop='isRelatedTo']"
	)
	var (
		err    error
		series string
		ok     bool
	)
	s = (s.Find(CSS).First())
	ok = (s.Length() != 0)
	if !ok {
		return series, fmt.Errorf("*s is empty")
	}
	series = (s.Text())
	ok = (len(series) != 0)
	if !ok {
		return series, fmt.Errorf("*s has no text")
	}
	series = strings.TrimSpace(series)
	return series, err
}

func parseAmiiboURL(s *goquery.Selection) (*address.Address, error) {
	const (
		CSS string = "a"
	)
	var (
		ok     bool
		rawurl string
	)
	s = s.Find(CSS)
	ok = (s.Length() != 0)
	if !ok {
		return nil, fmt.Errorf("*s is empty")
	}
	rawurl, ok = s.Attr("href")
	if !ok {
		return nil, fmt.Errorf("*s has no href")
	}
	rawurl = fmt.Sprintf("%s%s", resource.Nintendo, rawurl)
	return address.NewAddress(rawurl)
}
