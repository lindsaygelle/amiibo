package amiibo_test

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os/user"
	"path/filepath"
	"testing"

	"github.com/gellel/amiibo/amiibo"
	"github.com/gellel/amiibo/mix"
)

func Test(t *testing.T) {
	/*var (
		a, err = amiibo.NewAmiibo(&compatability.Amiibo{}, &lineup.Item{}, &lineup.Amiibo{})
	)
	if err != nil {
		t.Fatalf(err.Error())
	}
	b, _ := json.Marshal(a)
	fmt.Println(string(b))
	*/
	u, err := user.Current()
	if err != nil {
		panic(err)
	}
	m, err := mix.Get()
	if err != nil {
		panic(err)
	}
	a := amiibo.NewFromMix(m.Amiibo)
	b, err := json.Marshal(a)
	if err != nil {
		panic(err)
	}
	filename := fmt.Sprintf("%s.json", filepath.Join(u.HomeDir, "Desktop", "Amiibo"))
	err = ioutil.WriteFile(filename, b, 0777)
	if err != nil {
		panic(err)
	}
}
