package amiibo

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"golang.org/x/text/language"
)

var _ Software = (&ENGGame{})

// ENGGame is a formatted ENGChartGame and ENGChartItem.
type ENGGame struct {
	// Available is the Nintendo game availability status.
	Available bool `json:"available"`
	// Description is the verbose description for the Nintendo game product.
	Description string `json:"description"`
	// LastModified is the formatted timestamp when the dataset was modified by Nintendo of America.
	LastModified time.Time `json:"last_modified"`
	// Name is the official name of the Nintendo game product.
	//
	// Name can contain unicode.
	Name string `json:"name"`
	// Path is the relative path to the Nintendo game product details.
	Path string `json:"path"`
	// Product is the product classification of the Nintendo software item.
	Product string `json:"product"`
	// ProductImageURL is the direct URL to the Nintendo software product image.
	ProductImageURL string `json:"product_image_url"`
	// ReleaseDate is the formatted timestamp of the Nintendo software release date.
	ReleaseDate time.Time `json:"release_date"`
	// Title is the title for the Nintendo Amiibo product.
	Title string `json:"title"`
	// URL is the direct URL to the Nintendo Amiibo software page.
	URL string `json:"url"`
	// UUID is the UUID for the product given by Nintendo.
	UUID uuid.UUID `json:"uuid"`
}

// AddENGChartGame adds the contents of a ENGChartGame to the ENGGame.
func (e *ENGGame) AddENGChartGame(v *ENGChartGame) (err error) {
	var available bool
	available, err = strconv.ParseBool(v.IsReleased)
	if err != nil {
		return
	}
	e.Available = available
	e.Name = v.Name
	if reflect.ValueOf(v.Path).IsZero() {
		e.Path = v.Path
	}
	var releaseDate time.Time
	releaseDate, _ = time.Parse("2006-01-02", v.ReleaseDateMask)
	if err == nil {
		e.ReleaseDate = releaseDate
	}
	e.Product = strings.ToLower(v.Type)
	e.ProductImageURL = strings.ReplaceAll((NintendoURL + v.Image), " ", "%20")
	var UUID uuid.UUID
	UUID, err = uuid.Parse(v.ID)
	if err != nil {
		return
	}
	if reflect.ValueOf(e.URL).IsZero() {
		e.URL = strings.ReplaceAll((NintendoURL + v.URL), " ", "%20")
	}
	e.UUID = UUID
	return
}

// AddENGChartItem adds the contents of a ENGChartItem to the ENGGame.
func (e *ENGGame) AddENGChartItem(v *ENGChartItem) (err error) {
	e.Description = v.Description
	var lastModified time.Time
	lastModified = time.Unix(0, (v.LastModified * int64(time.Millisecond)))
	e.LastModified = lastModified
	if reflect.ValueOf(v.Path).IsZero() {
		e.Path = v.Path
	}
	e.Title = v.Title
	if reflect.ValueOf(e.URL).IsZero() {
		e.URL = strings.ReplaceAll((NintendoURL + v.URL), " ", "%20")
	}
	return
}

// GetAvailable returns the ENGGame availability.
func (e ENGGame) GetAvailable() bool {
	return time.Now().After(e.ReleaseDate)
}

// GetID returns the ENGGame ID.
func (e ENGGame) GetID() string {
	return strings.TrimSuffix(filepath.Base(e.URL), ".html")
}

// GetLanguage returns the ENGGame language.
func (e ENGGame) GetLanguage() language.Tag {
	return language.English
}

// GetName returns the ENGGame name.
func (e ENGGame) GetName() string {
	return e.Name
}

// GetNameAlternative returns the ENGGame name alternative.
func (e ENGGame) GetNameAlternative() string {
	return e.Name
}

// GetMD5 returns the ENGGame MD5.
func (e ENGGame) GetMD5() (MD5 string, err error) {
	var b ([]byte)
	b, err = marshal(&e, json.Marshal)
	if err != nil {
		return
	}
	MD5 = fmt.Sprintf("%x", md5.Sum(b))
	return
}

// GetReleaseDate returns the ENGGame release date.
func (e ENGGame) GetReleaseDate() time.Time {
	return e.ReleaseDate
}

// GetURL returns the ENGGame URL.
func (e ENGGame) GetURL() string {
	return e.URL
}

// NewENGGame returns a ENGGame.
func NewENGGame(ENGChartGame *ENGChartGame, ENGChartItem *ENGChartItem) (v ENGGame, err error) {
	var ok bool
	ok = ENGChartGame.GetID() == ENGChartItem.GetID()
	if !ok {
		err = fmt.Errorf("ENGChartGame != ENGChartItem")
	}
	if err != nil {
		return
	}
	err = v.AddENGChartGame(ENGChartGame)
	if err != nil {
		return
	}
	err = v.AddENGChartItem(ENGChartItem)
	if err != nil {
		return
	}
	return
}

// ReadENGGame reads a ENGGame from disc.
func ReadENGGame(dir string, filename string) (v ENGGame, err error) {
	err = readJSONFile(dir, filename, &v)
	return v, err
}

// WriteENGGame writes a ENGGame to disc.
func WriteENGGame(dir string, filename string, v *ENGGame) (fullpath string, err error) {
	fullpath, err = writeJSONFile(dir, filename, v)
	return
}
