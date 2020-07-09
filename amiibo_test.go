package amiibo_test

import (
	"path/filepath"
	"runtime"
	"testing"

	"github.com/lindsaygelle/amiibo"
)

var _, caller, _, _ = runtime.Caller(0)
var filefolder = filepath.Dir(caller)

func TestAmiibo(t *testing.T) {
	var name = "name"
	var nameAlt = "nameAlt"
	var engAmiibo amiibo.Amiibo = amiibo.ENGAmiibo{
		Name:            name,
		NameAlternative: nameAlt}
	if getName := engAmiibo.GetName(); getName != name {
		t.Fatalf("(Amiibo).GetName() %s != %s", getName, name)
	}
	if getNameAlternative := engAmiibo.GetNameAlternative(); getNameAlternative != nameAlt {
		t.Fatalf("(Amiibo).GetNameAlternative() %d != %s", getNameAlternative, nameAlt)
	}
}
