package amiibo

import (
	"fmt"
	"time"
)

func NewRawAmiiboUnix(unix int64) *RawAmiiboUnix {
	r := RawAmiiboUnix(unix)
	return &r
}

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

func (r *RawAmiiboUnix) Time() time.Time {
	return time.Unix(int64(*r), 0).UTC()
}
