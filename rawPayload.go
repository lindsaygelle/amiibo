package amiibo

import "encoding/json"

var (
	_ rawPayload = (*RawPayload)(nil)
)

func unmarshallRawPayload(content *[]byte) (*RawPayload, error) {
	r := &RawPayload{}
	err := json.Unmarshal(*content, r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

type rawPayload interface{}

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
