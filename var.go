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
	reStripHTML   = regexp.MustCompile(`(<[^>]*>|\n(\s{1,})?)`)
	reStripName   = regexp.MustCompile(`(\&\#[0-9]+\;|™)`)
	reStripSpaces = regexp.MustCompile(`\s{2,}`)
)
