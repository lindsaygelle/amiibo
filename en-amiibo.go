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

var _ Amiibo = (ENGAmiibo{})

// ENGAmiibo is a formatted ENGChartAmiibo, ENGLineupAmiibo and ENGLineupItem.
type ENGAmiibo struct {
	// Affiliation is the series of Nintendo Amiibo products the item is associated with.
	Affiliation string `json:"affiliation"`
	// Available is the Nintendo Amiibo availability status.
	Availiable bool `json:"availiable"`
	// BoxImageURL is the direct URL to the Nintendo Amiibo product box.
	BoxImageURL string `json:"box_image_url"`
	// Description is the verbose description for the Nintendo Amiibo product.
	Description string `json:"description"`
	// DescriptionAlternative is the alternative description for the Nintendo Amiibo product.
	DescriptionAlternative string `json:"description_alternative"`
	// Path is the relative path to the Nintendo Amiibo product details.
	DetailsPath string `json:"details_path"`
	// DetailsURL is the direct URL to the Nintendo Amiibo product page.
	DetailsURL string `json:"details_url"`
	// Epoch is the release date for the Nintendo Amiibo product in seconds.
	Epoch int64 `json:"epoch"`
	// FigureImageURL is the direct URL to the Nintendo Amiibo figurine image.
	FigureImageURL string `json:"figure_image_url"`
	// Franchise is the name of the series the Nintendo Amiibo product is affiliated with.
	Franchise string `json:"franchise"`
	// GameID is the ID associated with the Nintendo software the Nintendo Amiibo product is associated with.
	GameID string `json:"game_id"`
	// Hex is the hexidecimal code associated with the Nintendo Amiibo product.
	Hex string `json:"hex"`
	// ID is the fully qualified ID for the Nintendo Amiibo product given by Nintendo of America.
	ID string `json:"id"`
	// LastModified is the formatted timestamp when the dataset was modified by Nintendo of America.
	LastModified time.Time `json:"last_modified"`
	// Name is the official name of the Nintendo Amiibo product.
	//
	// Name can contain unicode.
	Name string `json:"name"`
	// NameAlternative is the alternative name given to the Nintendo Amiibo product.
	//
	// NameAlternative can contain unicode.
	NameAlternative string `json:"name_alternative"`
	// Path is the relative path to the Nintendo Amiibo product details.
	Path string `json:"path"`
	// Price is the price of the Nintendo Amiibo product in USD.
	//
	// Price can be empty.
	Price string `json:"price"`
	// Producer is the company affiliation of the Nintendo Amiibo product.
	Producer string `json:"producer"`
	// Product is the product classification of the Nintendo Amiibo item.
	Product string `json:"product"`
	// ProductAlternative is the alternative classification for the Nintendo Amiibo product.
	ProductAlternative string `json:"product_alternative"`
	// ProductImageURL is the direct URL to the Nintendo Amiibo product image.
	ProductImageURL string `json:"product_image_url"`
	// ProductPage is the relative nintendo.com URL to the Nintendo Amiibo product page.
	ProductPage string `json:"product_page"`
	// ReleaseDate is the formatted timestamp of the Nintendo Amiibo products release date.
	ReleaseDate time.Time `json:"release_date"`
	// Series is the defined series of products that the Nintendo Amiibo product is group or affiliated with.
	Series string `json:"series"`
	// Title is the title for the Nintendo Amiibo product.
	Title string `json:"title"`
	// TitleAternative is the alternative title for the Nintendo Amiibo product.
	//
	// TitleAlternative can contain unicode.
	TitleAlternative string `json:"title_alternative"`
	// UPC is the universal product code for the Nintendo Amiibo product.
	UPC string `json:"upc"`
	// URL is the direct URL to the Nintendo Amiibo product page.
	URL string `json:"url"`
	// UUID is the UUID for the product given by Nintendo.
	UUID uuid.UUID `json:"uuid"`
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
	e.Price = v.Price
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

// GetName returns the ENGAmiibo name.
func (e ENGAmiibo) GetName() string {
	return e.Name
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

// ReadENGAmiibo reads a ENGAmiibo from disc.
func ReadENGAmiibo(dir string, filename string) (v ENGAmiibo, err error) {
	err = readJSONFile(dir, filename, &v)
	return v, err
}

// WriteENGAmiibo writes a ENGAmiibo to disc.
func WriteENGAmiibo(dir string, filename string, v *ENGAmiibo) (fullpath string, err error) {
	fullpath, err = writeJSONFile(dir, filename, v)
	return
}
