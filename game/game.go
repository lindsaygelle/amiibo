package game

import (
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"sync"
	"time"

	"golang.org/x/text/language"

	"github.com/PuerkitoBio/goquery"

	"github.com/gellel/amiibo/address"
	"github.com/gellel/amiibo/compatability"
	"github.com/gellel/amiibo/errors"
	"github.com/gellel/amiibo/image"
	"github.com/gellel/amiibo/mix"
	"github.com/gellel/amiibo/network"
	"github.com/gellel/amiibo/resource"
	t "github.com/gellel/amiibo/text"
)

const (
	tep string = "%s%s" // tep string for raw url
)

const (
	// Version is the semver of game.Game.
	Version string = "1.0.0"
)

// Game is a structured representation of a Nintendo video-game that is compatible with a
// Nintendo Amiibo figurine product. Game structs are built from a mixture of resource that
// are provided from the amiibo/mix package.
// Games are consumed by amiibo/mux to create a basic HTTP REST API.
type Game struct {
	Compatability         []*Amiibo        `json:"compatability"`
	Complete              bool             `json:"complete"`
	Description           string           `json:"description"`
	GamePath              string           `json:"game_path"`
	GameURL               *address.Address `json:"game_url"`
	ID                    string           `json:"id"`
	Image                 *image.Image     `json:"image"`
	IsReleased            bool             `json:"is_released"`
	Language              language.Tag     `json:"language"`
	LastModified          int64            `json:"last_modified"`
	LastModifiedTimestamp time.Time        `json:"last_modified_timestamp"`
	Path                  string           `json:"path"`
	Name                  string           `json:"name"`
	ReleaseDateMask       string           `json:"release_date_mask"`
	Timestamp             time.Time        `json:"timestamp"`
	Title                 string           `json:"title"`
	Type                  string           `json:"type"`
	Unix                  int64            `json:"unix"`
	URI                   string           `json:"uri"`
	URL                   *address.Address `json:"url"`
	Version               string           `json:"version"`
}

// Get gets a field from the Game by its struct name and returns its string value.
func (g *Game) Get(key string) string {
	var r = reflect.ValueOf(g)
	var v = reflect.Indirect(r).FieldByName(key)
	return fmt.Sprintf("%s", v)
}

// NewGame creates a new instance of the game.Game from the aggregation
// of game structs across the amiibo package. Returns an error if all data points are
// not provided to the function.
func NewGame(c *compatability.Game, i *compatability.Item) (*Game, error) {
	var (
		ok bool
	)
	ok = (c != nil) || (i != nil)
	if !ok {
		return nil, errors.ErrArgsNil
	}
	var (
		g = &Game{
			Language: language.AmericanEnglish}
	)
	if c != nil {
		parseCompatability(g, c)
	}
	if i != nil {
		parseItem(g, i)
	}
	g.Compatability, _ = parseGameCompatability(g.URL.URL)
	g.Complete = c != nil && i != nil
	g.URI = t.URI(g.Name)
	g.Version = Version
	return g, nil
}

// NewFromMix creates a sequence of game.Game in O(N) time. Omits all mix.Game
// that cannot be instantiated by game.NewGame.
func NewFromMix(m map[string]*mix.Game) []*Game {
	var (
		s  = []*Game{}
		wg sync.WaitGroup
	)
	for _, m := range m {
		wg.Add(1)
		go func(m *mix.Game) {
			defer wg.Done()
			var (
				g, err = NewGame(m.Game, m.Item)
			)
			if err != nil {
				return
			}
			s = append(s, g)
		}(m)
	}
	wg.Wait()
	return s
}

// parseGameCompatability parses the HTML content from the Game's detail page.
func parseGameCompatability(rawurl string) ([]*Amiibo, error) {
	const (
		CSS string = "ul.figures li"
	)
	var (
		req, _ = http.NewRequest(http.MethodGet, rawurl, nil)
		res, _ = network.Client.Do(req)
	)
	var (
		doc, err = goquery.NewDocumentFromResponse(res)
	)
	if err != nil {
		return nil, err
	}
	var (
		amiibo = []*Amiibo{}
		s      = doc.Find(CSS)
	)
	s.Each(func(i int, s *goquery.Selection) {
		var (
			a, err = NewAmiibo(s)
		)
		if err != nil {
			return
		}
		amiibo = append(amiibo, a)
	})
	return amiibo, err
}

// parseGameIsReleased parses the release state of the Nintendo game from the compatability.Game.
func parseGameIsReleased(c *compatability.Game) (bool, error) {
	return strconv.ParseBool(c.IsReleased)
}

// parseGameName parses the name of the game using the name field from either compatability.Game or compatability.Item.
func parseGameName(s string) string {
	return t.Name(s)
}

// parseGameReleaseDateMask parses the release date mask from the compatability.Game.
func parseGameReleaseDateMask(c *compatability.Game) string {
	return c.ReleaseDateMask
}

// parsGameTimestamp parses the unix timestamp from either compatability.Game or compatability.Item.
func parsGameTimestamp(sec int64) time.Time {
	return time.Unix(sec, 0).UTC()
}

// parseGameURL parses the URL of the game using the URL field from either compatability.Game or compatability.Item.
func parseGameURL(rawurl string) (*address.Address, error) {
	return address.NewAddress(rawurl)
}

// parseCompatability parses all fields exposed in the compatability.Game and assigns them to the argument game.Game.
func parseCompatability(g *Game, c *compatability.Game) {
	g.GamePath = c.Path
	g.ID = c.ID
	g.Image, _ = image.NewImage(fmt.Sprintf(tep, resource.Nintendo, c.Image))
	g.IsReleased, _ = parseGameIsReleased(c)
	g.Name = parseGameName(c.Name)
	g.Path = c.Path
	g.ReleaseDateMask = parseGameReleaseDateMask(c)
	g.Timestamp, _ = time.Parse("2006-01-02", c.ReleaseDateMask)
	g.Type = c.Type
	g.Unix = g.Timestamp.Unix()
	g.URL, _ = parseGameURL(fmt.Sprintf(tep, resource.Nintendo, c.URL))
}

// parseItem parses all fields exposed in the compatability.Item and assigns them to the argument game.Game.
func parseItem(g *Game, i *compatability.Item) {
	g.Description = i.Description
	g.GameURL, _ = parseGameURL(fmt.Sprintf(tep, resource.Game, i.URL))
	g.LastModified = i.LastModified / 1000
	g.LastModifiedTimestamp = parsGameTimestamp(g.LastModified)
	g.Path = i.Path
	g.Title = i.Title
}
