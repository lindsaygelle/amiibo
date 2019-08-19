package amiibo

import (
	"testing"
)

func TestAmiiboMap(t *testing.T) {

	amiiboMap := newAmiiboMap()

	if testAmiiboStruct == nil {
		t.Fatalf("amiibo.newAmiiboMap() test cannot run testAmiiboStruct is nil")
	}

	if m := amiiboMap.Add(testAmiiboStruct); m != amiiboMap {
		t.Fatalf("AmiiboMap.Add(a *Amiibo) *AmiiboMap != %v", amiiboMap)
	}

	if n := amiiboMap.Len(); n < 1 {
		t.Fatalf("AmiiboMap.Len() int < 1")
	}

	if _, ok := (*amiiboMap.lexicon)[testAmiiboStruct.ID]; ok != true {
		t.Fatalf("AmiiboMap.Add(a *Amiibo) *AmiiboMap does not have key where ID = %s", testAmiiboStruct.ID)
	}
}
