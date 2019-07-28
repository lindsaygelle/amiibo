package amiibo

import (
	"encoding/json"
	"fmt"
)

var (
	_ rawPayload = (*RawPayload)(nil)
)

// NewRawPayload returns a new Raw Payload pointer. A Raw Payload pointer can be built
// from a cached XHR payload or directly from the Nintendo Amiibo source. To create from source
// parse in the optional byte code pointer, otherwise leave empty and it will be collected from
// the Nintendo XHR HTTP response.
func NewRawPayload(b ...byte) {}

// unmarshallRawPayload returns a new Raw Payload pointer using the argument bytes as the initial entries.
func unmarshallRawPayload(content *[]byte) (*RawPayload, error) {
	r := &RawPayload{}
	err := json.Unmarshal(*content, r)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// rawPayload describes the Raw Payload struct.
type rawPayload interface {
	String() string
}

// RawPayload is the raw HTTP response JSON that is fetched from the Nintendo Amiibo API.
// Contains a series of potentially optional fields that help describe the contents
// of the returned struct. There may be mechanisms that can be used to add or remove
// potential fields from this structure, but will require experimentation with the
// HTTP request headers that are passed into the HTTP GET request.
type RawPayload struct {
	AmiiboList           []*json.RawMessage `json:"amiiboList"`
	ComponentPath        string             `json:"componentPath"`
	DateFormatString     string             `json:"dateFormatString"`
	Items                []*json.RawMessage `json:"items"`
	LinkItems            bool               `json:"linkItems"`
	ShowDescription      bool               `json:"showDescription"`
	ShowModificationDate bool               `json:"showModificationDate"`
	Type                 string             `json:":type:"`
}

// String returns the string value of the raw Payload.
func (pointer *RawPayload) String() string {
	return fmt.Sprintf("%v", *pointer)
}
