package amiibo_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/lindsaygelle/amiibo"
)

var engLineup amiibo.ENGLineup
var engLineupFileName = "eng-lineup.json"
var engLineupFullpath = filepath.Join(filefolder, engLineupFileName)

func TestGetENGLineup(t *testing.T) {
	var err error
	if _, err := os.Stat(engLineupFullpath); !os.IsNotExist(err) {
		engLineup, err = amiibo.ReadENGLineup(filefolder, engLineupFileName)
		if err != nil {
			t.Fatal("amiibo.ReadENGLineup", err)
		}
	} else {
		_, _, engLineup, err = amiibo.GetENGLineup()
		if err != nil {
			t.Fatal(err)
		}
	}
	s, err := amiibo.WriteENGLineup(filefolder, engLineupFileName, engLineup)
	if err != nil {
		t.Fatal(err)
	}
	if s != engLineupFullpath {
		t.Fatalf("%s != %s", s, engLineupFullpath)
	}
	if l := len(engLineup.AmiiboList); l == 0 {
		t.Fatal("len: enLineup.AmiiboList", l)
	}
	if l := len(engLineup.Items); l == 0 {
		t.Fatal("len: enLineup.Items", l)
	}
}
