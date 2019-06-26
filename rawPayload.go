package amiibo

var (
	_ rawPayload = (*RawPayload)(nil)
)

type rawPayload interface{}

type RawPayload struct {
	AmiiboList           *RawAmiiboSlice     `json:"amiiboList"`
	ComponentPath        string              `json:"componentPath"`
	DateFormatString     string              `json:"dateFormatString"`
	Items                *RawAmiiboItemSlice `json:"items"`
	LinkItems            bool                `json:"linkItems"`
	ShowDescription      bool                `json:"showDescription"`
	ShowModificationDate bool                `json:"showModificationDate"`
	Type                 string              `json:":type:"`
}
