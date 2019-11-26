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
	Price           string           `json:"price"`
	ReleaseDateMask string           `json:"release_date_mask"`
	Series          string           `json:"series"`
	SeriesID        string           `json:"series_id"`
	Slug            string           `json:"slug"`
	TagID           string           `json:"tag_id"`
	Timestamp       time.Time        `json:"timestamp"`
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
		a               *Amiibo
		boxImage        *image.Image
		compatability   []*Game
		complete        bool
		description     string
		detailsPath     string
		detailsURL      *address.Address
		figureURL       *address.Address
		franchise       string
		franchiseID     string
		gameCode        string
		hex             string
		ID              string
		img             *image.Image
		isRelatedTo     string
		isReleased      bool
		language        = language.AmericanEnglish
		lastModified    int64
		name            string
		overview        string
		pageURL         *address.Address
		path            string
		presentedBy     string
		price           string
		rawurl          string
		releaseDateMask string
		series          string
		seriesID        string
		slug            string
		tagID           string
		timestamp       time.Time
		typeAlias       string
		typeOf          string
		unix            int64
		UPC             string
		URI             string
		URL             *address.Address
	)
	complete = (c != nil) && (l != nil) && (i != nil)
	if c != nil {
		ID = c.ID
		img, _ = image.NewImage(fmt.Sprintf(tep, resource.Nintendo, c.Image))
		isRelatedTo = c.IsRelatedTo
		isReleased, _ = strconv.ParseBool(c.IsReleased)
		name = t.Name(c.Name)
		tagID = c.TagID
		typeOf = c.Type
		URL, _ = address.NewAddress(fmt.Sprintf(tep, resource.Nintendo, c.URL))
	}
	if l != nil {
		boxImage, _ = parseAmiiboBoxImage(l)
		detailsPath = parseAmiiboDetailsPath(l)
		detailsURL, _ = parseAmiiboDetailsURL(l)
		figureURL, _ = parseAmiiboFigureURL(l)
		franchise = parseAmiiboFranchise(l)
		franchiseID = parseAmiiboFranchiseID(l)
		gameCode = parseAmiiboGameCode(l)
		hex = parseAmiiboHexCode(l)
		isReleased = l.IsReleased
		ok = (len(name) != 0)
		if !ok {
			name = parseAmiiboName(l)
		}
		overview = parseAmiiboOverviewDescription(l)
		pageURL, _ = parseAmiiboPageURL(l)
		presentedBy = t.PresentedBy(l.PresentedBy)
		price = l.Price
		releaseDateMask = l.ReleaseDateMask
		series = l.Series
		seriesID = t.URI(t.Name(series))
		slug = l.Slug
		timestamp = time.Unix(l.UnixTimestamp, 0).UTC()
		typeAlias = strings.ToLower(l.Type)
		UPC = l.UPC
		unix = l.UnixTimestamp
	}
	if i != nil {
		description = i.Description
		lastModified = i.LastModified
		path = i.Path
		ok = (len(name) != 0)
		if !ok {
			name = t.Name(i.Title)
		}
		ok = (URL != nil)
		if !ok {
			rawurl = strings.TrimPrefix((i.URL + "/"), noa)
			rawurl = fmt.Sprintf(tep, (resource.Amiibo + "/"), rawurl)
			URL, _ = address.NewAddress(rawurl)
		}
	}
	URI = t.URI(name)
	compatability, _ = GetGames(URL.URL)
	a = &Amiibo{
		BoxImage:        boxImage,
		Compatability:   compatability,
		Complete:        complete,
		Currency:        currencyISO,
		Description:     description,
		DetailsPath:     detailsPath,
		DetailsURL:      detailsURL,
		FigureURL:       figureURL,
		Franchise:       franchise,
		FranchiseID:     franchiseID,
		GameCode:        gameCode,
		HexCode:         hex,
		ID:              ID,
		Image:           img,
		IsRelatedTo:     isRelatedTo,
		IsReleased:      isReleased,
		Language:        language,
		LastModified:    lastModified,
		Name:            name,
		Overview:        overview,
		Path:            path,
		PageURL:         pageURL,
		PresentedBy:     presentedBy,
		Price:           price,
		ReleaseDateMask: releaseDateMask,
		Series:          series,
		SeriesID:        seriesID,
		Slug:            slug,
		TagID:           tagID,
		Timestamp:       timestamp,
		Type:            typeOf,
		TypeAlias:       typeAlias,
		Unix:            unix,
		UPC:             UPC,
		URI:             URI,
		URL:             URL}
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

func parseAmiiboName(l *lineup.Amiibo) string {
	return t.Name(l.AmiiboName)
}

func parseAmiiboOverviewDescription(l *lineup.Amiibo) string {
	return t.Untokenize(l.OverviewDescription)
}

func parseAmiiboPageURL(l *lineup.Amiibo) (*address.Address, error) {
	return address.NewAddress(fmt.Sprintf(tep, resource.Nintendo, l.AmiiboPage))
}
