package main

import "time"

var (
	_ rawAmiiboUnix = (*RawAmiiboUnix)(nil)
)

type rawAmiiboUnix interface{}

type RawAmiiboUnix int64

func (r RawAmiiboUnix) Time() time.Time {
	return time.Unix(int64(r), 0)
}
