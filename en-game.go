package amiibo

import (
	"fmt"
	"path/filepath"
	"reflect"
	"strconv"
	"time"

	"github.com/google/uuid"
)

// ENGGame is a formatted ENGChartGame and ENGChartItem.
type ENGGame struct {
	Available    bool      `json:"available"`
	Description  string    `json:"description"`
	LastModified time.Time `json:"last_modified"`
	Name         string    `json:"name"`
	Path         string    `json:"path"`
	Product      string    `json:"product"`
	ProductImage string    `json:"product_image"`
	ReleaseDate  time.Time `json:"release_date"`
	Title        string    `json:"title"`
	URI          string    `json:"uri"`
	URL          string    `json:"url"`
	UUID         uuid.UUID `json:"uuid"`
}

// AddENGChartGame adds the contents of a ENGChartGame to the ENGGame.
func (e *ENGGame) AddENGChartGame(v ENGChartGame) (err error) {
	var available bool
	available, err = strconv.ParseBool(v.IsReleased)
	if err != nil {
		return
	}
	e.Available = available
	e.Name = v.Name
	if !reflect.ValueOf(v.Path).IsZero() {
		e.Path = v.Path
	}
	var releaseDate time.Time
	releaseDate, _ = time.Parse("2006-01-02", v.ReleaseDateMask)
	if err == nil {
		e.ReleaseDate = releaseDate
	}
	e.Product = v.Type
	e.ProductImage = v.Image
	var UUID uuid.UUID
	UUID, err = uuid.Parse(v.ID)
	if err != nil {
		return
	}
	if !reflect.ValueOf(e.URL).IsZero() {
		e.URL = v.URL
	}
	e.UUID = UUID
	return
}

// AddENGChartItem adds the contents of a ENGChartItem to the ENGGame.
func (e *ENGGame) AddENGChartItem(v ENGChartItem) (err error) {
	e.Description = v.Description
	var lastModified time.Time
	lastModified = time.Unix(0, (v.LastModified * int64(time.Millisecond)))
	e.LastModified = lastModified
	if !reflect.ValueOf(v.Path).IsZero() {
		e.Path = v.Path
	}
	e.Title = v.Title
	if !reflect.ValueOf(e.Path).IsZero() {
		e.URI = filepath.Dir(v.Path)
	}
	if !reflect.ValueOf(e.URL).IsZero() {
		e.URL = v.URL
	}
	return
}

// NewENGGame returns a ENGGame.
func NewENGGame(ENGChartGame ENGChartGame, ENGChartItem ENGChartItem) (v ENGGame, err error) {
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
