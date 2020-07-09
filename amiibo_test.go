package amiibo_test

import (
	"path/filepath"
	"runtime"
	"testing"
	"time"

	"github.com/lindsaygelle/amiibo"
	"golang.org/x/text/language"
)

var _, caller, _, _ = runtime.Caller(0)
var filefolder = filepath.Dir(caller)

func TestAmiibo(t *testing.T) {
	var name = "name"
	var nameAlt = "nameAlt"
	var price = "1.99"
	var releaseDate = time.Now()
	var series = "series"
	var URL = "https://ninendo.com/"
	var engAmiibo amiibo.Amiibo = amiibo.ENGAmiibo{
		Name:            name,
		NameAlternative: nameAlt,
		Price:           price,
		ReleaseDate:     releaseDate,
		Series:          series,
		URL:             URL}
	var jpnAmiibo amiibo.Amiibo = amiibo.JPNAmiibo{
		Name:            name,
		NameAlternative: nameAlt,
		Price:           price,
		ReleaseDate:     releaseDate,
		Series:          series,
		URL:             URL}

	if getLanguage := engAmiibo.GetLanguage(); getLanguage != language.English {
		t.Fatalf("(Amiibo).GetLanguage() %v != %v", getLanguage, language.English)
	}
	if getLanguage := jpnAmiibo.GetLanguage(); getLanguage != language.Japanese {
		t.Fatalf("(Amiibo).GetLanguage() %v != %v", getLanguage, language.Japanese)
	}
	for _, v := range []amiibo.Amiibo{engAmiibo, jpnAmiibo} {
		if getName := v.GetName(); getName != name {
			t.Fatalf("(Amiibo).GetName() %s != %s", getName, name)
		}
		if getNameAlternative := v.GetNameAlternative(); getNameAlternative != nameAlt {
			t.Fatalf("(Amiibo).GetNameAlternative() %s != %s", getNameAlternative, nameAlt)
		}
		if MD5, err := v.GetMD5(); err != nil {
			t.Fatalf("(Amiibo).GetMD5() (%s, %s)", MD5, err)
		}
		if getPrice := v.GetPrice(); getPrice != price {
			t.Fatalf("(Amiibo).GetPrice() %s != %s", getPrice, price)
		}
		if getReleaseDate := v.GetReleaseDate(); getReleaseDate != releaseDate {
			t.Fatalf("(Amiibo).GetReleaseDate() %v != %v", getReleaseDate, releaseDate)
		}
		if getSeries := v.GetSeries(); getSeries != series {
			t.Fatalf("(Amiibo).GetSeries() %v != %v", getSeries, series)
		}
		if getURL := v.GetURL(); getURL != URL {
			t.Fatalf("(Amiibo).GetURL() %s != %s", getURL, URL)
		}
	}
}
