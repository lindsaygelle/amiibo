package amiibo

// JPNAmiiboSoftware is a formatted JPNChartItemSoftware.
type JPNAmiiboSoftware struct {
	Description string `json:"description"`
	Digital     bool   `json:"digital"`
	ID          string `json:"id"`
	Name        string `json:"name"`
}

// NewJPNAmiiboSoftware returns a new JPNAmiiboSoftware
func NewJPNAmiiboSoftware(JPNChartItemSoftware JPNChartItemSoftware) (v JPNAmiiboSoftware, err error) {
	v.Description = JPNChartItemSoftware.More
	v.Digital = JPNChartItemSoftware.Pickup != 0
	v.ID = JPNChartItemSoftware.Code
	v.Name = JPNChartItemSoftware.Name
	return
}
