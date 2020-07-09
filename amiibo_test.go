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
	if getLanguage := engAmiibo.GetLanguage(); getLanguage != language.English {
		t.Fatalf("(Amiibo).GetLanguage() %v != %v", getLanguage, language.English)
	}
	if getName := engAmiibo.GetName(); getName != name {
		t.Fatalf("(Amiibo).GetName() %s != %s", getName, name)
	}
	if getNameAlternative := engAmiibo.GetNameAlternative(); getNameAlternative != nameAlt {
		t.Fatalf("(Amiibo).GetNameAlternative() %s != %s", getNameAlternative, nameAlt)
	}
	if getPrice := engAmiibo.GetPrice(); getPrice != price {
		t.Fatalf("(Amiibo).GetPrice() %s != %s", getPrice, price)
	}
	if getReleaseDate := engAmiibo.GetReleaseDate(); getReleaseDate != releaseDate {
		t.Fatalf("(Amiibo).GetReleaseDate() %v != %v", getReleaseDate, releaseDate)
	}
	if getSeries := engAmiibo.GetSeries(); getSeries != series {
		t.Fatalf("(Amiibo).GetSeries() %v != %v", getSeries, series)
	}
	if getURL := engAmiibo.GetURL(); getURL != URL {
		t.Fatalf("(Amiibo).GetURL() %s != %s", getURL, URL)
	}
}
