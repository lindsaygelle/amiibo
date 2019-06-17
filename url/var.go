package url

import (
	"fmt"
	"regexp"
)

var (
	expression = fmt.Sprintf(pattern, scheme, username, password, hostname, port, path, filename, query, fragment)
)

var (
	re = regexp.MustCompile(expression)
)
