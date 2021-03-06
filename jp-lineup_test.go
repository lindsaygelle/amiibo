package amiibo_test

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/lindsaygelle/amiibo"
)

var jpnLineup amiibo.JPNLineup
var jpnLineupFileName = "jpn-lineup.xml"
var jpnLineupFullpath = filepath.Join(filefolder, jpnLineupFileName)

func TestGetJPNLineup(t *testing.T) {
	var err error
	if _, err := os.Stat(jpnLineupFullpath); !os.IsNotExist(err) {
		jpnLineup, err = amiibo.ReadJPNLineup(filefolder, jpnLineupFileName)
		if err != nil {
			t.Fatal("amiibo.ReadJPNLineup", err)
		}
	} else {
		_, _, jpnLineup, err = amiibo.GetJPNLineup()
		if err != nil {
			t.Fatal(err)
		}
	}
	s, err := amiibo.WriteJPNLineup(filefolder, jpnLineupFileName, &jpnLineup)
	if err != nil {
		t.Fatal(err)
	}
	if s != jpnLineupFullpath {
		t.Fatalf("%s != %s", s, jpnLineupFullpath)
	}
	if l := len(jpnLineup.Items); l == 0 {
		t.Fatal("len: jpnLineup.Items", l)
	}
	if l := len(jpnLineup.SeriesItems); l == 0 {
		t.Fatal("len: jpnLineup.SeriesItems", l)
	}
	_, err = amiibo.ReadJPNLineup(filefolder, jpnLineupFileName)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.ValueOf(jpnChart).IsZero() && !reflect.ValueOf(jpnLineup).IsZero() {
		testJPNAmiiboMap(t)
	}
}
