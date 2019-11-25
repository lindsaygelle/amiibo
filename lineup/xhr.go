package lineup

type XHR struct {
	Amiibo           []*Amiibo `json:"amiiboList"`
	ComponentPath    string    `json:"componentPath"`
	DateFormatString string    `json:"dataFormatString"`
	Items            []*Item   `json:"items"`
}
