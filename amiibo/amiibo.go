package amiibo

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"golang.org/x/text/currency"
	"golang.org/x/text/language"

	"github.com/gellel/amiibo/address"
	"github.com/gellel/amiibo/compatability"
	"github.com/gellel/amiibo/image"
	"github.com/gellel/amiibo/lineup"
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
	BoxImage        *image.Image     `json:"box_image"`
	Compatability   []*Game          `json:"compatability"`
	Complete        bool             `json:"complete"`
	Currency        string           `json:"currency"`
	Description     string           `json:"description"`
	DetailsPath     string           `json:"details_path"`
	DetailsURL      *address.Address `json:"details_url"`
	FigureURL       *address.Address `json:"figure_url"`
	Franchise       string           `json:"franchise"`
	FranchiseID     string           `json:"franchise_id"`
	GameCode        string           `json:"game_code"`
	HexCode         string           `json:"hex_code"`
	ID              string           `json:"id"`
	Image           *image.Image     `json:"image"`
	IsRelatedTo     string           `json:"is_related_to"`
	IsReleased      bool             `json:"is_released"`
	Language        language.Tag     `json:"language"`
	LastModified    int64            `json:"last_modified"`
	Name            string           `json:"name"`
	Overview        string           `json:"overview"`
	PageURL         *address.Address `json:"page"`
	Path            string           `json:"path"`
	PresentedBy     string           `json:"presented_by"`
	Price           float64          `json:"price"`
	ReleaseDateMask string           `json:"release_date_mask"`
	Series          string           `json:"series"`
	SeriesID        string           `json:"series_id"`
	Slug            string           `json:"slug"`
	TagID           string           `json:"tag_id"`
	Timestamp       time.Time        `json:"timestamp"`
	Title           string           `json:"title"`
	Type            string           `json:"type"`
	TypeAlias       string           `json:"type_alias"`
	Unix            int64            `json:"unix"`
	UPC             string           `json:"upc"`
	URI             string           `json:"uri"`
	URL             *address.Address `json:"url"`
}

// Get gets a field from the Amiibo by its struct name and returns its string value.
func (a *Amiibo) Get(key string) string {
	var r = reflect.ValueOf(a)
	var v = reflect.Indirect(r).FieldByName(key)
	return fmt.Sprintf("%s", v)
}

func NewAmiibo(c *compatability.Amiibo, l *lineup.Amiibo, i *lineup.Item) (*Amiibo, error) {
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
	a.Compatability, _ = GetGames(a.URL.URL)
	a.URI = t.URI(a.Name)
	return a, nil
}

func parseAmiiboBoxImage(l *lineup.Amiibo) (*image.Image, error) {
	return image.NewImage(fmt.Sprintf(tep, resource.Nintendo, l.BoxArtURL))
}

func parseAmiiboDetailsPath(l *lineup.Amiibo) string {
	return l.DetailsPath
}

func parseAmiiboDetailsURL(l *lineup.Amiibo) (*address.Address, error) {
	return address.NewAddress(fmt.Sprintf(tep, resource.Nintendo, l.DetailsURL))
}

func parseAmiiboFigureURL(l *lineup.Amiibo) (*address.Address, error) {
	return address.NewAddress(fmt.Sprintf(tep, resource.Nintendo, l.FigureURL))
}

func parseAmiiboFranchise(l *lineup.Amiibo) string {
	return l.Franchise
}

func parseAmiiboFranchiseID(l *lineup.Amiibo) string {
	return t.URI(t.Name(l.Franchise))
}

func parseAmiiboGameCode(l *lineup.Amiibo) string {
	return l.GameCode
}

func parseAmiiboHexCode(l *lineup.Amiibo) string {
	return l.HexCode
}

func parseAmiiboImage(c *compatability.Amiibo) (*image.Image, error) {
	return image.NewImage(fmt.Sprintf(tep, resource.Nintendo, c.Image))
}

func parseAmiiboIsReleased(c *compatability.Amiibo) (bool, error) {
	return strconv.ParseBool(c.IsReleased)
}

func parseAmiiboName(s string) string {
	return t.Name(s)
}

func parseAmiiboOverviewDescription(l *lineup.Amiibo) string {
	return t.Untokenize(l.OverviewDescription)
}

func parseAmiiboPageURL(l *lineup.Amiibo) (*address.Address, error) {
	return address.NewAddress(fmt.Sprintf(tep, resource.Nintendo, l.AmiiboPage))
}

func parseAmiiboPrice(l *lineup.Amiibo) (float64, error) {
	return strconv.ParseFloat(strings.TrimPrefix(l.Price, "$"), 32)
}

func parseAmiiboPresentedBy(l *lineup.Amiibo) string {
	return t.PresentedBy(l.PresentedBy)
}

func parseAmiiboReleaseDateMask(l *lineup.Amiibo) string {
	return l.ReleaseDateMask
}

func parseAmiiboSeriesID(l *lineup.Amiibo) string {
	return t.URI(t.Name(l.Series))
}

func parseAmiiboSlug(l *lineup.Amiibo) string {
	return l.Slug
}

func parseAmiiboTimestamp(l *lineup.Amiibo) time.Time {
	return time.Unix(l.UnixTimestamp, 0).UTC()
}

func parseAmiiboTitle(i *lineup.Item) string {
	return t.Name(i.Title)
}

func parseAmiiboTypeAlias(l *lineup.Amiibo) string {
	return strings.ToLower(l.Type)
}

func parseAmiiboURL(rawurl string) (*address.Address, error) {
	return address.NewAddress(rawurl)
}

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
	a.Timestamp = parseAmiiboTimestamp(l)
	a.TypeAlias = parseAmiiboTypeAlias(l)
	a.UPC = l.UPC
	a.Unix = l.UnixTimestamp
}

func parseItem(a *Amiibo, i *lineup.Item) {
	a.Description = i.Description
	a.LastModified = i.LastModified
	a.Path = i.Path
	a.Title = i.Title
	if a.URL == nil {
		a.URL, _ = parseAmiiboURL(fmt.Sprintf(tep, (resource.Amiibo + "/"), strings.TrimPrefix((i.URL+"/"), noa)))
	}
}
