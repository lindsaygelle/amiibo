package amiibo

var (
	_ rawPayload = (*RawPayload)(nil)
)

type rawPayload interface{}

// A RawPayload type represents the XHR HTTP response found on the Nintendo Amiibo line-up URI.
type RawPayload struct {
	AmiiboList           []*RawAmiibo `json:"amiiboList"`
	ComponentPath        string       `json:"componentPath"`
	DateFormatString     string       `json:"dateFormatString"`
	Items                []*RawItem   `json:"items"`
	LinkItems            bool         `json:"linkItems"`
	ShowDescription      bool         `json:"showDescription"`
	ShowModificationDate bool         `json:"showModificationDate"`
	Type                 string       `json:":type:"`
}
