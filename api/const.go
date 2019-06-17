package api

import "time"

const (
	timeout time.Duration = 10
)

const (
	URL string = "https://www.amiiboapi.com/api/amiibo/?"
)

const (
	ID string = URL + "id=%s"
)

const (
	NAME string = URL + "name=%s"
)
