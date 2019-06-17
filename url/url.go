package url

import (
	"strings"
)

func New(resource string) URL {
	var (
		ok = (strings.HasPrefix(resource, HTTP) || strings.HasPrefix(resource, HTTPS))
	)
	if ok != true {
		resource = (HTTPS + resource)
	}
	return URL(resource)
}

func NewURL(resource string) *URL {
	var (
		URL = New(resource)
	)
	return &URL
}

type url interface {
	Host() string
	HTTP() string
	HTTPS() string
	Protocol() string
}

type URL string

func (pointer *URL) Host() string {
	var (
		URL = string(*pointer)
	)
	return URL
}

func (pointer *URL) HTTP() bool {
	var (
		URL = string(*pointer)
	)
	return strings.HasPrefix(URL, HTTP)
}

func (pointer *URL) HTTPS() bool {
	var (
		URL = string(*pointer)
	)
	return strings.HasPrefix(URL, HTTPS)
}

func (pointer *URL) Protocol() string {
	var (
		URL = string(*pointer)
	)
	return URL
}

func (pointer *URL) Keyset() map[string]string {
	var (
		content = pointer.Unpack()
	)
	var (
		m = map[string]string{}
	)
	for i, name := range re.SubexpNames() {
		if i > 0 && i < len(content) {
			m[name] = content[i]
		}
	}
	return m
}

func (pointer *URL) Unpack() []string {
	var (
		URL = string(*pointer)
	)
	return re.FindStringSubmatch(URL)
}
