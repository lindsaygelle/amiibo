package amiibo

// JPNAmiiboSoftware is a formatted JPNChartItemSoftware.
type JPNAmiiboSoftware struct {
	// Description is the verbose description for the Nintendo software product.
	Description string `json:"description"`
	// Digital indicates the whether the Nintendo software product is a digital only release.
	Digital bool `json:"digital"`
	// ID is the fully qualified ID for the Nintendo software product given by Nintendo Japan.
	ID string `json:"id"`
	// Name is the official name of the Nintendo software product.
	//
	// Name contains Japanese Hiragana.
	Name string `json:"name"`
	// URL is the direct URL to the Nintendo software product page.
	URL string `json:"url"`
}

// AddJPNChartItemSoftware adds a JPNChartItemSoftware to the JPNAmiiboSoftware
func (j *JPNAmiiboSoftware) AddJPNChartItemSoftware(v *JPNChartItemSoftware) (err error) {
	j.Description = v.More
	j.Digital = v.Pickup != 0
	j.ID = v.Code
	j.Name = v.Name
	j.URL = "https://www.nintendo.co.jp/hardware/amiibo/game/" + j.ID
	return
}

// GetID returns the JPNAmiiboSoftware ID.
func (j JPNAmiiboSoftware) GetID() string {
	return j.ID
}

// NewJPNAmiiboSoftware returns a new JPNAmiiboSoftware
func NewJPNAmiiboSoftware(JPNChartItemSoftware *JPNChartItemSoftware) (v JPNAmiiboSoftware, err error) {
	err = (&v).AddJPNChartItemSoftware(JPNChartItemSoftware)
	return
}

// ReadJPNAmiiboSoftware reads a JPNAmiiboSoftware from disc.
func ReadJPNAmiiboSoftware(dir string, filename string) (v JPNAmiiboSoftware, err error) {
	err = readJSONFile(dir, filename, &v)
	return v, err
}

// WriteJPNAmiiboSoftware writes a JPNAmiiboSoftware to disc.
func WriteJPNAmiiboSoftware(dir string, filename string, v *JPNAmiiboSoftware) (fullpath string, err error) {
	fullpath, err = writeJSONFile(dir, filename, v)
	return
}
