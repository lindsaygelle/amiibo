package amiibo

import (
	"strconv"
	"time"

	"github.com/google/uuid"
)

// ENGAmiibo is a formatted ENGChartAmiibo, ENGLineupAmiibo and ENGLineupItem.
type ENGAmiibo struct {

	// Affiliation property is sourced from ENGChartAmiibo.IsRelatedTo OR ENGLineupAmiibo.PresentedBy
	Affiliation string

	// Available property is sourced from ENGChartAmiibo.IsReleased OR ENGLineupAmiibo.IsReleased
	Available bool

	// BoxImage property is sourced from ENGLineupAmiibo.BoxArtURL
	BoxImage string

	// Classification property is sourced from ENGChartAmiibo.Type
	Classification string

	// Description property is sourced from ENGLineupAmiibo.OverviewDescription OR ENGLineupItem.Description
	Description string

	// DetailsPath property is sourced from ENGLineupAmiibo.DetailsPath OR ENGLineupItem.Path
	DetailsPath string

	// DetailsURL property is sourced from ENGLineupAmiibo.AmiiboPage OR ENGLineupAmiibo.DetailsURL OR ENGLineupItem.URL
	DetailsURL string

	// Epoch property is sourced from ENGLineupAmiibo.UnixTimestamp
	Epoch int64

	// Hex property is sourced from ENGLineupAmiibo.HexCode
	Hex string

	// Franchise property is sourced from ENGLineupAmiibo.Franchise
	Franchise string

	// GameID property is sourced from ENGLineupAmiibo.GameCode
	GameID string

	// LastModified property is sourced from ENGLineupItem.LastModified
	LastModified time.Time

	// Name property is sourced from ENGChartAmiibo.Name OR ENGLineupAmiibo.AmiiboName OR ENGLineupItem.Title
	Name string

	// Price property is sourced from ENGLineupAmiibo.Price
	Price float32

	// Producer ENGChartAmiibo.TagID
	Producer string

	// Product property is sourced from ENGLineupAmiibo.Type
	Product string

	// ProductImage property is sourced from ENGChartAmiibo.Image OR ENGLineupAmiibo.Image
	ProductImage string

	// ProductURL property is sourced from ENGChartAmiibo.URL
	ProductURL string

	// ReleaseDate property is sourced from ENGChartAmiibo.ReleaseDateMask OR ENGLineupAmiibo.ReleaseDateMask
	ReleaseDate time.Time

	// Series property is sourced from ENGChartAmiibo.IsRelatedTo OR ENGLineupAmiibo.Series
	Series string

	// Title property is sourced from ENGChartLineupAmiibo.Slug
	Title string

	// UPC property is sourced from ENGLineupAmiibo.UPC
	UPC string

	// UUID property is sourced from ENGChartAmiibo.ID
	UUID uuid.UUID
}

// addENGChartAmiibo adds the content of a ENGChartAmiibo.
func (e ENGAmiibo) addENGChartAmiibo(v ENGChartAmiibo) {
	e.Affiliation = v.IsRelatedTo
	var available, _ = strconv.ParseBool(v.IsReleased)
	e.Available = available
	e.Classification = v.Type
	e.Name = v.Name
	e.Producer = v.TagID
	e.ProductImage = v.Image
	e.ProductURL = v.URL
	e.Series = v.IsRelatedTo
	var releaseDate, _ = time.Parse("2006-01-02", v.ReleaseDateMask)
	e.ReleaseDate = releaseDate
	var UUID, _ = uuid.Parse(v.ID)
	e.UUID = UUID
}
