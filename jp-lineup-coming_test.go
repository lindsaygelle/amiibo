package amiibo_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/lindsaygelle/amiibo"
)

var jpnLineupComing amiibo.JPNLineupComing
var jpnLineupComingFileName = "jpn-lineup-coming.xml"
var jpnLineupComingFullpath = filepath.Join(filefolder, jpnLineupComingFileName)

func TestGetJPNLineupComing(t *testing.T) {
	var err error
	if _, err := os.Stat(jpnLineupComingFullpath); !os.IsNotExist(err) {
		jpnLineupComing, err = amiibo.ReadJPNLineupComing(filefolder, "jpn-lineup-coming.xml")
		if err != nil {
			t.Fatal("amiibo.ReadJPNLineupComing", err)
		}
	} else {
		_, _, jpnLineupComing, err = amiibo.GetJPNLineupComing()
		if err != nil {
			t.Fatal(err)
		}
	}
	s, err := amiibo.WriteJPNLineupComing(filefolder, jpnLineupComingFileName, &jpnLineupComing)
	if err != nil {
		t.Fatal(err)
	}
	if s != jpnLineupComingFullpath {
		t.Fatalf("%s != %s", s, jpnLineupComingFullpath)
	}
	if l := len(jpnLineupComing.Items); l == 0 {
		t.Fatal("len: jpnLineupComing.Items", l)
	}
}
