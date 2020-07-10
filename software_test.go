package amiibo_test

import (
	"testing"
	"time"

	"github.com/lindsaygelle/amiibo"
	"golang.org/x/text/language"
)

func TestSoftware(t *testing.T) {
	var name = "name"
	var nameAlt = "nameAlt"
	var releaseDate = time.Now()
	var URL = "https://ninendo.com/"
	var engSoftware amiibo.Software = amiibo.ENGGame{
		Name:        name,
		ReleaseDate: releaseDate,
		URL:         URL}
	var jpnSoftware amiibo.Software = amiibo.JPNSoftware{
		Name:            name,
		NameAlternative: nameAlt,
		ReleaseDate:     releaseDate,
		URL:             URL}

	if getLanguage := engSoftware.GetLanguage(); getLanguage != language.English {
		t.Fatalf("(Software).GetLanguage() %v != %v", getLanguage, language.English)
	}
	if getLanguage := jpnSoftware.GetLanguage(); getLanguage != language.Japanese {
		t.Fatalf("(Software).GetLanguage() %v != %v", getLanguage, language.Japanese)
	}
	for _, v := range []amiibo.Software{engSoftware, jpnSoftware} {
		if getAvailable := v.GetAvailable(); !getAvailable {
			t.Fatalf("(Software).GetAvailable() %t != true", getAvailable)
		}
		if getName := v.GetName(); getName != name {
			t.Fatalf("(Software).GetName() %s != %s", getName, name)
		}
		if getNameAlternative := v.GetNameAlternative(); !(getNameAlternative == nameAlt || getNameAlternative == name) {
			t.Fatalf("(Software).GetNameAlternative() %s != %s || %s", getNameAlternative, nameAlt, name)
		}
		if MD5, err := v.GetMD5(); err != nil {
			t.Fatalf("(Software).GetMD5() (%s, %s)", MD5, err)
		}
		if getReleaseDate := v.GetReleaseDate(); getReleaseDate != releaseDate {
			t.Fatalf("(Software).GetReleaseDate() %v != %v", getReleaseDate, releaseDate)
		}
		if getURL := v.GetURL(); getURL != URL {
			t.Fatalf("(Software).GetURL() %s != %s", getURL, URL)
		}
	}
}
