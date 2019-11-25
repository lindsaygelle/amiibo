package compatability

type XHR struct {
	Amiibo           []*Amiibo `json:"amiiboList"`
	AuthorMode       bool      `json:"authorMode"`
	ComponentPath    string    `json:"componentPath"`
	DateFormatString string    `json:"dataFormatString"`
	Games            []*Game   `json:"gameList"`
	Items            []*Item   `json:"items"`
	Language         string    `json:"language"`
	Mode             string    `json:"mode"`
}
