package amiibo

import (
	"fmt"
	"html"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

// ENGAmiibo is a formatted ENGChartAmiibo, ENGLineupAmiibo and ENGLineupItem.
type ENGAmiibo struct {
	Affiliation            string    `json:"affiliation"`
	Availiable             bool      `json:"availiable"`
	BoxImageURL            string    `json:"box_image_url"`
	Description            string    `json:"description"`
	DescriptionAlternative string    `json:"description_alternative"`
	DetailsPath            string    `json:"details_path"`
	DetailsURL             string    `json:"details_url"`
	Epoch                  int64     `json:"epoch"`
	FigureImageURL         string    `json:"figure_image_url"`
	Franchise              string    `json:"franchise"`
	GameID                 string    `json:"game_id"`
	Hex                    string    `json:"hex"`
	ID                     string    `json:"id"`
	LastModified           time.Time `json:"last_modified"`
	Name                   string    `json:"name"`
	NameAlternative        string    `json:"name_alternative"`
	Path                   string    `json:"path"`
	Price                  float64   `json:"price"`
	Producer               string    `json:"producer"`
	Product                string    `json:"product"`
	ProductAlternative     string    `json:"product_alternative"`
	ProductImageURL        string    `json:"product_image_url"`
	ProductPage            string    `json:"product_page"`
	ReleaseDate            time.Time `json:"release_date"`
	Series                 string    `json:"series"`
	Title                  string    `json:"title"`
	TitleAlternative       string    `json:"title_alternative"`
	UPC                    string    `json:"upc"`
	URL                    string    `json:"url"`
	UUID                   uuid.UUID `json:"uuid"`
}

// AddENGChartAmiibo adds the contents of a ENGChartAmiibo to the ENGAmiibo.
func (e *ENGAmiibo) AddENGChartAmiibo(v *ENGChartAmiibo) (err error) {
	e.Affiliation = v.IsRelatedTo
	var available bool
	available, err = strconv.ParseBool(v.IsReleased)
	if err != nil {
		return
	}
	e.Availiable = available
	e.ID = v.TagID
	e.Name = v.Name
	e.ProductAlternative = strings.ToLower(v.Type)
	e.ProductImageURL = v.Image
	var releaseDate time.Time
	if reflect.ValueOf(e.ReleaseDate).IsZero() {
		releaseDate, err = time.Parse("2006-01-02", v.ReleaseDateMask)
		if err == nil {
			e.ReleaseDate = releaseDate
		}
		err = nil
	}
	if reflect.ValueOf(e.URL).IsZero() {
		e.URL = strings.ReplaceAll(("https://nintendo.com" + v.URL), " ", "%20")
	}
	var UUID uuid.UUID
	UUID, err = uuid.Parse(v.ID)
	if err != nil {
		return
	}
	e.UUID = UUID
	return
}

// AddENGLineupAmiibo adds the contents of a ENGLineupAmiibo to the ENGAmiibo.
func (e *ENGAmiibo) AddENGLineupAmiibo(v *ENGLineupAmiibo) (err error) {
	e.BoxImageURL = strings.ReplaceAll(("https://nintendo.com" + v.BoxArtURL), " ", "%20")
	var description = v.OverviewDescription
	description = regexpSpaces.ReplaceAllString(regexpHTML.ReplaceAllString(description, " "), " ")
	description = html.UnescapeString(strings.TrimSpace(description))
	e.Description = description
	e.DetailsPath = v.DetailsPath
	e.DetailsURL = strings.ReplaceAll(("https://nintendo.com" + v.DetailsURL), " ", "%20")
	e.Epoch = v.UnixTimestamp
	e.Franchise = v.Franchise
	e.GameID = v.GameCode
	e.Hex = v.HexCode
	e.NameAlternative = v.AmiiboName
	var price float64
	if reflect.ValueOf(v.Price).IsZero() {
		if len(v.Price) != 0 {
			price, err = strconv.ParseFloat(v.Price, 32)
			if err != nil {
				return
			}
		}
	}
	e.Price = price
	e.Product = strings.ToLower(v.Type)
	e.Producer = v.PresentedBy
	e.ProductImageURL = strings.ReplaceAll(("https://nintendo.com" + v.FigureURL), " ", "%20")
	e.ProductPage = v.AmiiboPage
	var releaseDate time.Time
	if reflect.ValueOf(e.ReleaseDate).IsZero() {
		releaseDate, err = time.Parse("01-02-2006", strings.ReplaceAll(v.ReleaseDateMask, "/", "-"))
		if err == nil {
			e.ReleaseDate = releaseDate
		}
		err = nil
	}
	e.Series = v.Series
	e.Title = v.Slug
	e.UPC = v.UPC
	return
}

// AddENGLineupItem adds the contents of a ENGLineupItem to the ENGAmiibo.
func (e *ENGAmiibo) AddENGLineupItem(v *ENGLineupItem) (err error) {
	e.DescriptionAlternative = v.Description
	var lastModified time.Time
	lastModified = time.Unix(0, (v.LastModified * int64(time.Millisecond)))
	e.LastModified = lastModified
	e.Path = v.Path
	e.TitleAlternative = v.Title
	if reflect.ValueOf(e.URL).IsZero() {
		e.URL = strings.ReplaceAll(("https://nintendo.com" + v.URL), " ", "%20")
	}
	return
}

// GetID returns the ENGAmiibo ID.
func (e ENGAmiibo) GetID() string {
	return strings.TrimSuffix(filepath.Base(e.URL), ".html")
}

// NewENGAmiibo returns a ENGAmiibo.
func NewENGAmiibo(ENGChartAmiibo *ENGChartAmiibo, ENGLineupAmiibo *ENGLineupAmiibo, ENGLineupItem *ENGLineupItem) (v ENGAmiibo, err error) {
	var ok bool
	ok = ENGChartAmiibo.GetID() == ENGLineupAmiibo.GetID()
	if !ok {
		err = fmt.Errorf("ENGChartAmiibo != ENGLineupAmiibo")
	}
	if err != nil {
		return
	}
	ok = ENGLineupAmiibo.GetID() == ENGLineupItem.GetID()
	if !ok {
		err = fmt.Errorf("ENGLineupAmiibo != ENGLineupItem")
	}
	if err != nil {
		return
	}
	err = v.AddENGChartAmiibo(ENGChartAmiibo)
	if err != nil {
		return
	}
	err = v.AddENGLineupAmiibo(ENGLineupAmiibo)
	if err != nil {
		return
	}
	err = v.AddENGLineupItem(ENGLineupItem)
	return
}
