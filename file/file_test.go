package file_test

import (
	"testing"

	"github.com/gellel/amiibo/dir"
	"github.com/gellel/amiibo/file"
)

func Test(t *testing.T) {

	file.Make(dir.Current())
}
