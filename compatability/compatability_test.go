package compatability_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/gellel/amiibo/lineup"

	"github.com/gellel/amiibo/compatability"
)

func Test(t *testing.T) {
	var (
		c, err = compatability.Get()
	)
	var (
		l, _ = lineup.Get()
	)
	if err != nil {
		t.Fatalf(err.Error())
	}
	x := map[string]int{}
	y := map[string]int{}
	for _, a := range c.Amiibo {
		k := a.Key()
		if _, ok := x[k]; !ok {
			x[k] = 0
		}
		x[k]++
	}
	for _, a := range l.Amiibo {
		k := a.Key()

		if _, ok := x[k]; !ok {
			x[k] = 0
		}
		x[k]++
	}
	for _, a := range l.Items {
		k := a.Key()
		if _, ok := x[k]; !ok {
			x[k] = 0
		}
		x[k]++
	}
	for k, v := range x {
		fmt.Println(k, "->", v)
	}
	time.Sleep(time.Second * 3)
	for _, a := range c.Games {
		k := a.Key()
		if _, ok := y[k]; !ok {
			y[k] = 0
		}
		y[k]++
	}
	for _, a := range c.Items {
		k := a.Key()
		if _, ok := y[k]; !ok {
			y[k] = 0
		}
		y[k]++
	}
	for k, v := range y {
		fmt.Println(k, "->", v)
	}
}
