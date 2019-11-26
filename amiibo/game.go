package amiibo

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"

	"github.com/gellel/amiibo/address"
	"github.com/gellel/amiibo/image"
	"github.com/gellel/amiibo/resource"
	t "github.com/gellel/amiibo/text"
)

// Game is a representation of a Nintendo video-game that is directly compatible
// with a Nintendo Amiibo figurine.
type Game struct {
	Image *image.Image     `json:"image"`
	Name  string           `json:"name"`
	URL   *address.Address `json:"url"`
}

func NewGame(s *goquery.Selection) (*Game, error) {
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
		game     Game
		image, _ = parseGameImage(s)
		name, _  = parseGameName(s)
		URL, _   = parseGameURL(s)
	)
	game = Game{
		Image: image,
		Name:  name,
		URL:   URL}
	return &game, nil
}

func parseGameImage(s *goquery.Selection) (*image.Image, error) {
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
	rawurl = fmt.Sprintf(tep, resource.Nintendo, rawurl)
	return image.NewImage(rawurl)
}

func parseGameName(s *goquery.Selection) (string, error) {
	const (
		CSS string = "a[title]"
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
	name, ok = s.Attr("title")
	if !ok {
		return name, fmt.Errorf("*s has no title")
	}
	return t.Untokenize(name), err
}

func parseGameURL(s *goquery.Selection) (*address.Address, error) {
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
	rawurl = fmt.Sprintf(tep, resource.Nintendo, rawurl)
	return address.NewAddress(rawurl)
}
