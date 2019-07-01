package amiibo

var (
	_ rawPayloadChecker = (*rawPayload)(nil)
)

type rawPayloadChecker interface{}

type rawPayload struct {
	AmiiboList           []*RawAmiibo `json:"amiiboList"`
	ComponentPath        string       `json:"componentPath"`
	DateFormatString     string       `json:"dateFormatString"`
	Items                []*RawItem   `json:"items"`
	LinkItems            bool         `json:"linkItems"`
	ShowDescription      bool         `json:"showDescription"`
	ShowModificationDate bool         `json:"showModificationDate"`
	Type                 string       `json:":type:"`
}
