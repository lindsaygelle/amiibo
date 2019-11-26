package amiibo_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/gellel/amiibo/amiibo"
	"github.com/gellel/amiibo/compatability"
	"github.com/gellel/amiibo/lineup"
)

func Test(t *testing.T) {
	var (
		a, err = amiibo.NewAmiibo(&compatability.Amiibo{}, &lineup.Amiibo{}, &lineup.Item{})
	)
	if err != nil {
		t.Fatalf(err.Error())
	}
	b, _ := json.Marshal(a)
	fmt.Println(string(b))
}
