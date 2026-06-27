package main

import (
	"fmt"
	"learngo/dict"
)

func main() {
	dictionary := dict.Dictionary{"first": "First word"}
	dictionary["name"] = "jihun"
	definition, err := dictionary.Search("first")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}
	fmt.Println(dictionary)
}
