package main

import (
	"fmt"
	"testing"
)

func TestAPINewAPI(t *testing.T) {

	api := NewAPI()

	fmt.Println(api.HTTP())
}

func TestAPINot404(t *testing.T) {

	api := NewAPI()

	response := api.HTTP()

	fmt.Println(response)
}

func TestAPITimeout(t *testing.T) {

}
