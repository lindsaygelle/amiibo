package amiibo

import (
	"path/filepath"
	"regexp"
	"runtime"
)

var (
	_, file, _, _ = runtime.Caller(0)
	rootpath      = filepath.Dir(file)
)

var (
	// reStripHTML is the regex pattern that matches all valid HTML patterns.
	reStripHTML = regexp.MustCompile(`(<[^>]*>|\n(\s{1,})?)`)
	// reStripName is the regex pattern that matches all unsupported characters in an Amiibo or Item's name.
	reStripName = regexp.MustCompile(`(\&\#[0-9]+\;|™)`)
	// reStripSpaces is the regexp pattern that matches all double or n following whitespace.
	reStripSpaces = regexp.MustCompile(`\s{2,}`)
)
