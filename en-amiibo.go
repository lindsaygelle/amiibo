package amiibo

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"html"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"golang.org/x/text/language"
)

var _ Amiibo = (ENGAmiibo{})

// ENGAmiibo is a formatted ENGChartAmiibo, ENGLineupAmiibo and ENGLineupItem.
type ENGAmiibo struct {

	// Affiliation is the series of Nintendo Amiibo products the item is associated with.
	Affiliation string `json:"affiliation"`

	// Available is the Nintendo Amiibo availability status.
	Availiable bool `json:"availiable"`

	// BoxImage is the image data for the Nintendo Amiibo product box.
	BoxImage *Image `json:"box_image,omitempty"`

	// BoxImageURL is the direct URL to the Nintendo Amiibo product box.
	BoxImageURL string `json:"box_image_url"`

	// CompatibilityURL is the direct URL to the Nintendo Amiibo product software compatibility.
	CompatibilityURL string `json:"compatibility_url"`

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

	// ProductImage is the image data for the Nintend Amiibo product.
	ProductImage *Image `json:"product_image,omitempty"`

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
	e.CompatibilityURL = (NintendoURL + "/amiibo/compatibility/#compatible/amiibo/" + v.ID)
	e.ID = v.TagID
	e.Name = v.Name
	e.ProductAlternative = strings.ToLower(v.Type)
	if reflect.ValueOf(e.ProductImageURL).IsZero() {
		e.ProductImageURL = strings.ReplaceAll((NintendoURL + v.Image), " ", "%20")
	}
	var releaseDate time.Time
	if reflect.ValueOf(e.ReleaseDate).IsZero() {
		releaseDate, err = time.Parse("2006-01-02", v.ReleaseDateMask)
		if err == nil {
			e.ReleaseDate = releaseDate
		}
		err = nil
	}
	if reflect.ValueOf(e.URL).IsZero() {
		e.URL = strings.ReplaceAll((NintendoURL + v.URL), " ", "%20")
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
	e.BoxImageURL = strings.ReplaceAll((NintendoURL + v.BoxArtURL), " ", "%20")
	var description = v.OverviewDescription
	description = regexpSpaces.ReplaceAllString(regexpHTML.ReplaceAllString(description, " "), " ")
	description = html.UnescapeString(strings.TrimSpace(description))
	e.Description = description
	e.DetailsPath = v.DetailsPath
	e.DetailsURL = strings.ReplaceAll((NintendoURL + v.DetailsURL), " ", "%20")
	e.Epoch = v.UnixTimestamp
	e.Franchise = v.Franchise
	e.GameID = v.GameCode
	e.Hex = v.HexCode
	e.NameAlternative = v.AmiiboName
	e.Price = v.Price
	e.Product = strings.ToLower(v.Type)
	e.Producer = v.PresentedBy
	e.ProductImageURL = strings.ReplaceAll((NintendoURL + v.FigureURL), " ", "%20")
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
		e.URL = strings.ReplaceAll((NintendoURL + v.URL), " ", "%20")
	}
	return
}

// GetAvailable returns the ENGAmiibo availability.
func (e ENGAmiibo) GetAvailable() bool {
	return time.Now().After(e.ReleaseDate)
}

// GetID returns the ENGAmiibo ID.
func (e ENGAmiibo) GetID() string {
	return strings.TrimSuffix(filepath.Base(e.URL), ".html")
}

// GetLanguage returns the ENGAmiibo language.
func (e ENGAmiibo) GetLanguage() language.Tag {
	return language.English
}

// GetName returns the ENGAmiibo name.
func (e ENGAmiibo) GetName() string {
	return e.Name
}

// GetNamespace returns the ENGAmiibo formatted name.
func (e ENGAmiibo) GetNamespace() (s string) {
	var franchiseOK, seriesOK, productAlternativeOK = len(e.Franchise) != 0, len(e.Series) != 0, len(e.ProductAlternative) != 0
	s = e.Title
	if !franchiseOK && !seriesOK && !productAlternativeOK {
		return s
	}
	var substrings = []string{}
	if franchiseOK {
		substrings = append(substrings, e.Franchise)
	}
	if seriesOK {
		substrings = append(substrings, e.Series)
	}
	for _, substring := range substrings {
		var v = strings.ToLower(strings.ReplaceAll(regexPunctuation.ReplaceAllString(substring, ""), " ", "-"))
		s = strings.ReplaceAll(s, (v + "-" + "series"), "")
	}
	if productAlternativeOK {
		s = strings.ReplaceAll(s, e.ProductAlternative, "")
	}
	s = regexpHyphens.ReplaceAllString(s, "-")
	s = strings.Trim(s, "-")
	return
}

// GetNameAlternative returns the ENGAmiibo name alternative.
func (e ENGAmiibo) GetNameAlternative() string {
	return e.NameAlternative
}

// GetMD5 returns the ENGAmiibo MD5.
func (e ENGAmiibo) GetMD5() (MD5 string, b []byte, err error) {
	b, err = marshal(&e, json.Marshal)
	if err != nil {
		return
	}
	MD5 = fmt.Sprintf("%x", md5.Sum(b))
	return
}

// GetPrice returns the ENGAmiibo price.
func (e ENGAmiibo) GetPrice() string {
	return e.Price
}

// GetReleaseDate returns the ENGAmiibo release date.
func (e ENGAmiibo) GetReleaseDate() time.Time {
	return e.ReleaseDate
}

// GetSeries returns the ENGAmiibo series.
func (e ENGAmiibo) GetSeries() string {
	return e.Series
}

// GetURL returns the ENGAmiibo URL.
func (e ENGAmiibo) GetURL() string {
	return e.URL
}

// GetENGAmiiboBoxImage gets the box art image for the ENGAmiibo.
func GetENGAmiiboBoxImage(ENGAmiibo *ENGAmiibo) (err error) {
	var v Image
	v, err = GetImage(ENGAmiibo.BoxImageURL)
	if err != nil {
		return
	}
	ENGAmiibo.BoxImage = &v
	return
}

// GetENGAmiiboProductImage gets the product image for the ENGAmiibo.
func GetENGAmiiboProductImage(ENGAmiibo *ENGAmiibo) (err error) {
	var v Image
	v, err = GetImage(ENGAmiibo.ProductImageURL)
	if err != nil {
		return
	}
	ENGAmiibo.ProductImage = &v
	return
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
