package main

import (
	"fmt"
	"learngo/dict"
)

func main() {
	dictionary := dict.Dictionary{}
	word := "hello"
	dictionary.Add(word, "first")
	err := dictionary.Update("a", "second")
	if err != nil {
		fmt.Println(err)
	}
	def, _ := dictionary.Search(word)
	dictionary.Delete(word)
	fmt.Println(def)
	fmt.Println(dictionary)
}
