package amiibo

import (
	"net/http"
	"net/url"
)

type jp struct {
	Cookies    []*http.Cookie
	Status     string
	StatusCode int
	URL        *url.URL
}
