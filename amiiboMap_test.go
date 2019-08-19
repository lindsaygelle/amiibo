package amiibo

import (
	"testing"
)

func TestAmiiboMap(t *testing.T) {

	amiiboMap := newAmiiboMap()

	if testAmiiboStruct == nil {
		t.Fatalf("amiibo.newAmiiboMap() test cannot run testAmiiboStruct is nil")
	}

	amiiboMap.Add(testAmiiboStruct)

	if _, ok := (*amiiboMap.lexicon)[testAmiiboStruct.ID]; ok != true {
		t.Fatalf("AmiiboMap.Add(a *Amiibo) *AmiiboMap does not have key where ID = %s", testAmiiboStruct.ID)
	}
}
