package amiibo

import (
	"path/filepath"
	"runtime"
)

var (
	_, file, _, _ = runtime.Caller(0)
	rootpath      = filepath.Dir(filepath.Dir(file))
)
