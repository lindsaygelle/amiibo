package amiibo

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"

	"github.com/gellel/amiibo/address"
	"github.com/gellel/amiibo/image"
	"github.com/gellel/amiibo/resource"
	t "github.com/gellel/amiibo/text"
)

var (
	errSEmpty   = fmt.Errorf("*s is empty")     // err goquery.Selection no nodes.
	errSNil     = fmt.Errorf("*s is nil")       // err goquery.Selection nil address.
	errSNoHref  = fmt.Errorf("*s has no href")  // err goquery.Selection no href attribute.
	errSNoSrc   = fmt.Errorf("*s has no src")   // err goquery.Selection no src attribute.
	errSNoTitle = fmt.Errorf("*s has no title") // err goquery.Selection no title attribute.
)

// Game is a representation of a Nintendo video-game that is directly compatible
// with a Nintendo Amiibo figurine.
type Game struct {
	Image *image.Image     `json:"image"`
	Name  string           `json:"name"`
	URL   *address.Address `json:"url"`
}

// NewGame creates a new instance of the amiibo.Game.
func NewGame(s *goquery.Selection) (*Game, error) {
	var (
		ok bool
	)
	ok = (s != nil)
	if !ok {
		return nil, errSNil
	}
	ok = (s.Length() != 0)
	if !ok {
		return nil, errSEmpty
	}
	var (
		game = &Game{}
	)
	game.Image, _ = parseGameImage(s)
	game.Name, _ = parseGameName(s)
	game.URL, _ = parseGameURL(s)
	return game, nil
}

// parseGameImage parses the game box image from the goquery.Selection.
func parseGameImage(s *goquery.Selection) (*image.Image, error) {
	const (
		CSS string = "img"
	)
	var (
		ok     bool
		rawurl string
	)
	if s == nil {
		return nil, errSNil
	}
	s = (s.Find(CSS).First())
	ok = (s.Length() != 0)
	if !ok {
		return nil, errSEmpty
	}
	rawurl, ok = s.Attr("src")
	if !ok {
		return nil, errSNoSrc
	}
	rawurl = fmt.Sprintf(tep, resource.Nintendo, rawurl)
	return image.NewImage(rawurl)
}

// parseGameName parses the game name from the goquery.Selection.
func parseGameName(s *goquery.Selection) (string, error) {
	const (
		CSS string = "a[title]"
	)
	var (
		name string
	)
	if s == nil {
		return name, errSNil
	}
	var (
		err error
		ok  bool
	)
	s = (s.Find(CSS).First())
	ok = (s.Length() != 0)
	if !ok {
		return name, errSEmpty
	}
	name, ok = s.Attr("title")
	if !ok {
		return name, errSNoTitle
	}
	return t.Untokenize(name), err
}

// parseGameURL parses the game URL from the goquery.Selection.
func parseGameURL(s *goquery.Selection) (*address.Address, error) {
	const (
		CSS string = "a"
	)
	if s == nil {
		return nil, errSNil
	}
	var (
		ok     bool
		rawurl string
	)
	s = s.Find(CSS)
	ok = (s.Length() != 0)
	if !ok {
		return nil, errSEmpty
	}
	rawurl, ok = s.Attr("href")
	if !ok {
		return nil, errSNoHref
	}
	rawurl = fmt.Sprintf(tep, resource.Nintendo, rawurl)
	return address.NewAddress(rawurl)
}
