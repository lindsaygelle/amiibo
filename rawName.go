package amiibo

import (
	"fmt"
	"regexp"
)

var (
	_ rawAmiiboName = (*RawAmiiboName)(nil)
)

type rawAmiiboName interface {
	String() string
}

// A RawAmiiboName string represents the unparsed Amiibo name found in the amiiboName property
// that is held by a RawAmiibo within in the Nintendo XHR HTTP response.
type RawAmiiboName string

func (r *RawAmiiboName) String() string {
	return fmt.Sprintf("%s", regexp.MustCompile(`(&#\d+;|™)`).ReplaceAllString(string(*r), ""))
}
