package amiibo

// JPNAmiiboSoftware is a formatted JPNChartItemSoftware.
type JPNAmiiboSoftware struct {
	Description string `json:"description"`
	Digital     bool   `json:"digital"`
	ID          string `json:"id"`
	Name        string `json:"name"`
}

// AddJPNChartItemSoftware adds a JPNChartItemSoftware to the JPNAmiiboSoftware
func (j *JPNAmiiboSoftware) AddJPNChartItemSoftware(v *JPNChartItemSoftware) (err error) {
	j.Description = v.More
	j.Digital = v.Pickup != 0
	j.ID = v.Code
	j.Name = v.Name
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
