package amiibo_test

import (
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/lindsaygelle/amiibo"
)

var engAmiiboMapFileName = "eng-amiibo-map.json"

func testENGAmiiboMap(t *testing.T) {
	var v, err = amiibo.NewENGAmiiboMap(engChart, engLineup)
	if err != nil {
		t.Fatal(err)
	}
	if len(v) == 0 {
		t.Fatalf("len(engAmiiboMap) == 0")
	}
	if l := len(v); l != ((len(engChart.AmiiboList) + len(engLineup.AmiiboList) + len(engLineup.Items)) / 3) {
		t.Logf("engAmiiboMap %d engChart.AmiiboList %d engLineup.AmiiboList %d engLineup.Items %d", l, len(engChart.AmiiboList), len(engLineup.AmiiboList), len(engLineup.Items))
	}
	_, err = amiibo.WriteENGAmiiboMap(filefolder, engAmiiboMapFileName, &v)
	if err != nil {
		t.Fatal(err)
	}
	_, err = amiibo.ReadENGAmiiboMap(filefolder, engAmiiboMapFileName)
	if err != nil {
		t.Fatal(err)
	}
	(&v).Each(func(k string, v amiibo.ENGAmiibo) {
		fmt.Println(v.GetNamespace(), "-", strings.ToLower(strings.ReplaceAll(regexp.MustCompile(`(\.|\,|\:)`).ReplaceAllString(v.Series, ""), " ", "-")))
		fmt.Println("---")
	})
	engAmiibo := &amiibo.ENGAmiibo{URL: "/1"}
	if ok := v.Add(engAmiibo); !ok {
		t.Fatal("(ENGAmiiboMap).Add(*ENGAmiibo) bool != true")
	}
	a, ok := v.Get("1")
	if !ok {
		t.Fatal("(ENGAmiiboMap).Get(string) (ENGAmiibo, bool) != _, true")
	}
	if a.GetID() != "1" {
		t.Fatalf("ENGAmiiboMap provides incorrect keys; %s != %s", a.GetID(), strings.TrimPrefix(engAmiibo.URL, "/"))
	}
}
