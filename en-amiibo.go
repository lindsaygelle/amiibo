package amiibo

import (
	"fmt"
	"html"
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
	BoxImage               string    `json:"box_image"`
	Description            string    `json:"description"`
	DescriptionAlternative string    `json:"description_alternative"`
	DetailsPath            string    `json:"details_path"`
	DetailsURL             string    `json:"details_url"`
	Epoch                  int64     `json:"epoch"`
	FigureImage            string    `json:"figure_image"`
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
	ProductImage           string    `json:"product_image"`
	ProductPage            string    `json:"product_page"`
	ReleaseDate            time.Time `json:"release_date"`
	ReleaseDateAlternative time.Time `json:"release_date_alternative"`
	Series                 string    `json:"series"`
	Title                  string    `json:"title"`
	TitleAlternative       string    `json:"title_alternative"`
	UPC                    string    `json:"upc"`
	URL                    string    `json:"url"`
	UUID                   uuid.UUID `json:"uuid"`
}

// AddENGChartAmiibo adds the contents of a ENGChartAmiibo to the ENGAmiibo.
func (e *ENGAmiibo) AddENGChartAmiibo(v ENGChartAmiibo) (err error) {
	e.Affiliation = v.IsRelatedTo
	var available bool
	available, err = strconv.ParseBool(v.IsReleased)
	if err != nil {
		return
	}
	e.Availiable = available
	e.FigureImage = v.Image
	e.ID = v.TagID
	e.Name = v.Name
	e.ProductAlternative = v.Type
	var releaseDateAlternative time.Time
	releaseDateAlternative, err = time.Parse("2006-01-02", v.ReleaseDateMask)
	if err == nil {
		e.ReleaseDateAlternative = releaseDateAlternative
	}
	e.ReleaseDateAlternative = releaseDateAlternative
	if reflect.ValueOf(e.URL).IsZero() {
		e.URL = "http://www.nintendo.com" + v.URL
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
func (e *ENGAmiibo) AddENGLineupAmiibo(v ENGLineupAmiibo) (err error) {
	e.BoxImage = v.BoxArtURL
	var description = v.OverviewDescription
	description = regexpSpaces.ReplaceAllString(regexpHTML.ReplaceAllString(description, " "), " ")
	description = html.UnescapeString(strings.TrimSpace(description))
	e.Description = description
	e.DetailsPath = v.DetailsPath
	e.DetailsURL = v.DetailsURL
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
	e.Product = v.Type
	e.Producer = v.PresentedBy
	e.ProductImage = v.FigureURL
	e.ProductPage = v.AmiiboPage
	var releaseDate time.Time
	releaseDate, _ = time.Parse("2006-01-02", v.ReleaseDateMask)
	if err == nil {
		e.ReleaseDate = releaseDate
	}
	e.ReleaseDate = releaseDate
	e.Series = v.Series
	e.Title = v.Slug
	e.UPC = v.UPC
	return
}

// AddENGLineupItem adds the contents of a ENGLineupItem to the ENGAmiibo.
func (e *ENGAmiibo) AddENGLineupItem(v ENGLineupItem) (err error) {
	e.DescriptionAlternative = v.Description
	var lastModified time.Time
	lastModified = time.Unix(0, (v.LastModified * int64(time.Millisecond)))
	e.LastModified = lastModified
	e.Path = v.Path
	e.TitleAlternative = v.Title
	if reflect.ValueOf(e.URL).IsZero() {
		e.URL = "http://www.nintendo.com" + v.URL
	}
	return
}

// NewENGAmiibo returns a ENGAmiibo.
func NewENGAmiibo(ENGChartAmiibo ENGChartAmiibo, ENGLineupAmiibo ENGLineupAmiibo, ENGLineupItem ENGLineupItem) (v ENGAmiibo, err error) {
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
