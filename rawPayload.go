package amiibo

import "fmt"

var (
	_ rawPayload = (*RawPayload)(nil)
)

type rawPayload interface{}

// A RawPayload type represents the XHR HTTP response found on the Nintendo Amiibo line-up URI.
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

func (r *RawPayload) String() string {
	return fmt.Sprintf("{&[%v] &[%v]}", r.AmiiboList.Len(), r.Items.Len())
}
