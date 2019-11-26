package amiibo

import (
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gellel/amiibo/mix"

	"golang.org/x/text/currency"
	"golang.org/x/text/language"

	"github.com/PuerkitoBio/goquery"
	"github.com/gellel/amiibo/address"
	"github.com/gellel/amiibo/compatability"
	"github.com/gellel/amiibo/image"
	"github.com/gellel/amiibo/lineup"
	"github.com/gellel/amiibo/network"
	"github.com/gellel/amiibo/resource"
	t "github.com/gellel/amiibo/text"
)

const (
	noa string = "/content/noa/en_US/" // Nintendo of America prefix.
)

const (
	rep string = ""     // rep string
	sep string = "."    // sep string
	tep string = "%s%s" // tep string for raw url
)

const (
	// Version is the semver of amiibo.Amiibo.
	Version string = "1.0.0"
)

var (
	currencyISO = currency.USD.String() // Currency ISO for all Nintendo Amiibo products.
)

var (
	errNilAll = fmt.Errorf("*c, *l and *i are nil")
)

// Amiibo is a structured representation of a Nintendo Amiibo figuring.
// Amiibo structs are built from a mixture of resources that
// are provided from the amiibo/mix package.
// Amiibos are consumed by amiibo/mux to create a basic HTTP REST API.
type Amiibo struct {
	BoxImage              *image.Image     `json:"box_image"`
	Compatability         []*Game          `json:"compatability"`
	Complete              bool             `json:"complete"`
	Currency              string           `json:"currency"`
	Description           string           `json:"description"`
	DetailsPath           string           `json:"details_path"`
	DetailsURL            *address.Address `json:"details_url"`
	FigureURL             *address.Address `json:"figure_url"`
	Franchise             string           `json:"franchise"`
	FranchiseID           string           `json:"franchise_id"`
	GameCode              string           `json:"game_code"`
	HexCode               string           `json:"hex_code"`
	ID                    string           `json:"id"`
	Image                 *image.Image     `json:"image"`
	IsRelatedTo           string           `json:"is_related_to"`
	IsReleased            bool             `json:"is_released"`
	Language              language.Tag     `json:"language"`
	LastModified          int64            `json:"last_modified"`
	LastModifiedTimestamp time.Time        `json:"last_modified_timestamp"`
	Name                  string           `json:"name"`
	Overview              string           `json:"overview"`
	PageURL               *address.Address `json:"page"`
	Path                  string           `json:"path"`
	PresentedBy           string           `json:"presented_by"`
	Price                 float64          `json:"price"`
	ReleaseDateMask       string           `json:"release_date_mask"`
	Series                string           `json:"series"`
	SeriesID              string           `json:"series_id"`
	Slug                  string           `json:"slug"`
	TagID                 string           `json:"tag_id"`
	Timestamp             time.Time        `json:"timestamp"`
	Title                 string           `json:"title"`
	Type                  string           `json:"type"`
	TypeAlias             string           `json:"type_alias"`
	Unix                  int64            `json:"unix"`
	UPC                   string           `json:"upc"`
	URI                   string           `json:"uri"`
	URL                   *address.Address `json:"url"`
}

// Get gets a field from the Amiibo by its struct name and returns its string value.
func (a *Amiibo) Get(key string) string {
	var r = reflect.ValueOf(a)
	var v = reflect.Indirect(r).FieldByName(key)
	return fmt.Sprintf("%s", v)
}

// NewAmiibo creates a new instance of the amiibo.Amiibo from the aggregation
// of amiibo structs across the amiibo package. Returns an error if all data points are
// not provided to the function.
func NewAmiibo(c *compatability.Amiibo, i *lineup.Item, l *lineup.Amiibo) (*Amiibo, error) {
	var (
		ok bool
	)
	ok = (c != nil) || (l != nil) || (i != nil)
	if !ok {
		return nil, errNilAll
	}
	var (
		a = &Amiibo{
			Currency: currencyISO,
			Language: language.AmericanEnglish}
	)
	if c != nil {
		parseCompatability(a, c)
	}
	if l != nil {
		parseLineup(a, l)
	}
	if i != nil {
		parseItem(a, i)
	}
	a.Compatability, _ = parseAmiiboCompatability(a.URL.URL)
	a.Complete = c != nil && i != nil && l != nil
	a.URI = t.URI(a.Name)
	return a, nil
}

// NewFromMix creates a sequence of amiibo.Amiibo in O(N) time. Omits all mix.Amiibo
// that cannot be instantiated by amiibo.NewAmiibo.
func NewFromMix(m map[string]*mix.Amiibo) []*Amiibo {
	var (
		s  = []*Amiibo{}
		wg sync.WaitGroup
	)
	for _, m := range m {
		wg.Add(1)
		go func(m *mix.Amiibo) {
			defer wg.Done()
			var (
				a, err = NewAmiibo(m.Compatability, m.Item, m.Lineup)
			)
			if err != nil {
				return
			}
			s = append(s, a)
		}(m)
	}
	wg.Wait()
	return s
}

// parseAmiiboBoxImage parses the box art image from the lineup.Amiibo.
func parseAmiiboBoxImage(l *lineup.Amiibo) (*image.Image, error) {
	return image.NewImage(fmt.Sprintf(tep, resource.Nintendo, l.BoxArtURL))
}

// parseAmiiboCompatability parses the HTML content from the Amiibo's detail page.
func parseAmiiboCompatability(rawurl string) ([]*Game, error) {
	const (
		CSS string = "ul#game-set li"
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
		games = []*Game{}
		s     = doc.Find(CSS)
	)
	s.Each(func(i int, s *goquery.Selection) {
		var (
			g, err = NewGame(s)
		)
		if err != nil {
			return
		}
		games = append(games, g)
	})
	return games, err
}

// parseAmiiboDetailsPath parses the details path from the lineup.Amiibo.
func parseAmiiboDetailsPath(l *lineup.Amiibo) string {
	return l.DetailsPath
}

// parseAmiiboDetailsURL parses the details URL from the lineup.Amiibo.
func parseAmiiboDetailsURL(l *lineup.Amiibo) (*address.Address, error) {
	return address.NewAddress(fmt.Sprintf(tep, resource.Nintendo, l.DetailsURL))
}

// parsesAmiiboFigureURL parses the Amiibo figuring image from the lineup.Amiibo.
func parseAmiiboFigureURL(l *lineup.Amiibo) (*address.Address, error) {
	return address.NewAddress(fmt.Sprintf(tep, resource.Nintendo, l.FigureURL))
}

// parseAmiiboFranchise parses the Amiibo's franchise from the lineup.Amiibo.
func parseAmiiboFranchise(l *lineup.Amiibo) string {
	return l.Franchise
}

// parseAmiiboFranchiseID parses the Amiibo's franchise ID using the franchise string from the lineup.Amiibo.
func parseAmiiboFranchiseID(l *lineup.Amiibo) string {
	return t.URI(t.Name(l.Franchise))
}

// parseAmiiboGameCode parses the Amiibo's game ID from the lineup.Amiibo.
func parseAmiiboGameCode(l *lineup.Amiibo) string {
	return l.GameCode
}

// parseAmiiboHexCode parses the Amiibo's hex ID from the lineup.Amiibo.
func parseAmiiboHexCode(l *lineup.Amiibo) string {
	return l.HexCode
}

// parseAmiiboImage parses the Amiibo figuring image from the compatability.Amiibo.
func parseAmiiboImage(c *compatability.Amiibo) (*image.Image, error) {
	return image.NewImage(fmt.Sprintf(tep, resource.Nintendo, c.Image))
}

// parseAmiiboIsReleased parses the release state of the Amiibo product from the compatability.Amiibo.
func parseAmiiboIsReleased(c *compatability.Amiibo) (bool, error) {
	return strconv.ParseBool(c.IsReleased)
}

// parseAmiiboTimestamp parses the unix last modified timestamp from the lineup.Item.
func parseAmiiboLastModifiedTimestamp(i *lineup.Item) time.Time {
	return time.Unix(i.LastModified, 0).UTC()
}

// parseAmiiboName parses the name of the Amiibo using the name field from either compatability.Amiibo or lineup.Amiibo.
func parseAmiiboName(s string) string {
	return t.Name(s)
}

// parseAmiiboOverviewDescription parses the Amiibo description from the lineup.Amiibo.
func parseAmiiboOverviewDescription(l *lineup.Amiibo) string {
	return t.Untokenize(l.OverviewDescription)
}

// parseAmiiboPageURL parses the Amiibo detail page URL from the lineup.Amiibo.
func parseAmiiboPageURL(l *lineup.Amiibo) (*address.Address, error) {
	return address.NewAddress(fmt.Sprintf(tep, resource.Nintendo, l.AmiiboPage))
}

// parseAmiiboPrice parses the Amiibo price from the lineup.Amiibo.
func parseAmiiboPrice(l *lineup.Amiibo) (float64, error) {
	return strconv.ParseFloat(l.Price, 64)
}

// parseAmiiboPresentedBy parses the presenter from the lineup.Amiibo.
func parseAmiiboPresentedBy(l *lineup.Amiibo) string {
	return t.PresentedBy(l.PresentedBy)
}

// parseAmiiboReleaseDateMask parses the release date mask from the lineup.Amiibo.
func parseAmiiboReleaseDateMask(l *lineup.Amiibo) string {
	return l.ReleaseDateMask
}

// parseAmiiboSeriesID parses the Amiibo's series ID using the series string from the lineup.Amiibo.
func parseAmiiboSeriesID(l *lineup.Amiibo) string {
	return t.URI(t.Name(l.Series))
}

// parseAmiiboSlug parses the slug from the lineup.Amiibo.
func parseAmiiboSlug(l *lineup.Amiibo) string {
	return l.Slug
}

// parseAmiiboTimestamp parses the unix timestamp from either compatability.Amiibo or lineup.Amiibo.
func parseAmiiboTimestamp(sec int64) time.Time {
	return time.Unix(sec, 0).UTC()
}

// parseAmiiboTitle parses the title from the lineup.Item.
func parseAmiiboTitle(i *lineup.Item) string {
	return t.Name(i.Title)
}

// parseAmiiboTypeAlias parses the Amiibo type alias from the lineup.Amiibo.
func parseAmiiboTypeAlias(l *lineup.Amiibo) string {
	return strings.ToLower(l.Type)
}

// parseAmiiboURL parses the URL of the Amiibo using the URL field from either compatability.Amiibo or lineup.Amiibo.
func parseAmiiboURL(rawurl string) (*address.Address, error) {
	return address.NewAddress(rawurl)
}

// parseCompatability parses all fields exposed in the compatability.Amiibo and assigns them to the argument amiibo.Amiibo.
func parseCompatability(a *Amiibo, c *compatability.Amiibo) {
	a.ID = c.ID
	a.Image, _ = parseAmiiboImage(c)
	a.IsRelatedTo = c.IsRelatedTo
	a.IsReleased, _ = parseAmiiboIsReleased(c)
	if len(c.Name) == 0 {
		a.Name = parseAmiiboName(c.Name)
	}
	a.TagID = c.TagID
	a.Type = c.Type
	a.URL, _ = parseAmiiboURL(fmt.Sprintf(tep, resource.Nintendo, c.URL))
}

// parseLineup parses all fields exposed in the lineup.Amiibo and assigns them to the argument amiibo.Amiibo.
func parseLineup(a *Amiibo, l *lineup.Amiibo) {
	a.BoxImage, _ = parseAmiiboBoxImage(l)
	a.DetailsPath = parseAmiiboDetailsPath(l)
	a.DetailsURL, _ = parseAmiiboDetailsURL(l)
	a.FigureURL, _ = parseAmiiboFigureURL(l)
	a.Franchise = parseAmiiboFranchise(l)
	a.FranchiseID = parseAmiiboFranchiseID(l)
	a.GameCode = parseAmiiboGameCode(l)
	a.HexCode = parseAmiiboHexCode(l)
	a.IsReleased = l.IsReleased
	if len(a.Name) == 0 {
		a.Name = parseAmiiboName(l.AmiiboName)
	}
	a.Overview = parseAmiiboOverviewDescription(l)
	a.PageURL, _ = parseAmiiboPageURL(l)
	a.PresentedBy = parseAmiiboPresentedBy(l)
	a.Price, _ = parseAmiiboPrice(l)
	a.ReleaseDateMask = parseAmiiboReleaseDateMask(l)
	a.Series = l.Series
	a.SeriesID = parseAmiiboSeriesID(l)
	a.Slug = parseAmiiboSlug(l)
	a.Timestamp = parseAmiiboTimestamp(l.UnixTimestamp)
	a.TypeAlias = parseAmiiboTypeAlias(l)
	a.UPC = l.UPC
	a.Unix = l.UnixTimestamp
}

// parseItem parses all fields exposed in the lineup.Item and assigns them to the argument amiibo.Amiibo.
func parseItem(a *Amiibo, i *lineup.Item) {
	a.Description = i.Description
	a.LastModified = i.LastModified / 1000
	a.LastModifiedTimestamp = parseAmiiboTimestamp(a.LastModified)
	a.Path = i.Path
	a.Title = i.Title
	if a.URL == nil {
		a.URL, _ = parseAmiiboURL(fmt.Sprintf(tep, (resource.Amiibo + "/"), strings.TrimPrefix((i.URL+"/"), noa)))
	}
}
