package main

import (
	"fmt"
	"golangStudy/myDict"
)

func main() {
	dictionary := myDict.Dictionary{"first": "First word"}
	word := "hello"
	definition := "Greeting"
	err := dictionary.Add(word, definition)
	if err != nil{
		fmt.Println(err)
	}
	hello , err := dictionary.Search(word)
	
	fmt.Println(hello)
}
