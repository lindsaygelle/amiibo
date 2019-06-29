package amiibo

import (
	"fmt"
	"time"
)

var (
	_ rawAmiiboUnix = (*RawAmiiboUnix)(nil)
)

type rawAmiiboUnix interface {
	String() string
	Time() time.Time
}

// A RawAmiiboUnix int64 represents the unix nano integer found in the unixTimestamp property
// that is held by a RawAmiibo within in the Nintendo XHR HTTP response.
type RawAmiiboUnix int64

func (r *RawAmiiboUnix) String() string {
	return fmt.Sprintf("%v", string(*r))
}

// Time returns a parsed time.Time struct from the datetime string value.
func (r *RawAmiiboUnix) Time() time.Time {
	return time.Unix(int64(*r), 0).UTC()
}
