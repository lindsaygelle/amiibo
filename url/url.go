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
	Path() string
	Protocol() string
}

type URL string

func (pointer *URL) Host() string {
	var (
		m = pointer.Map()
	)
	return m[HOSTNAME]
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

func (pointer *URL) Map() map[string]string {
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

func (pointer *URL) Path() string {
	var (
		m = pointer.Map()
	)
	return m[PATH]
}

func (pointer *URL) Unpack() []string {
	var (
		URL = string(*pointer)
	)
	return re.FindStringSubmatch(URL)
}
