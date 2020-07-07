package amiibo_test

import (
	"path/filepath"
	"runtime"
)

var _, caller, _, _ = runtime.Caller(0)
var filefolder = filepath.Dir(caller)
