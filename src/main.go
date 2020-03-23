package main

import (
	"awesomeProject/src/myDict"
	"fmt"
)

func main() {
	dictionary := myDict.Dictionary{"first" : "First word"}

	word := "hello"
	defintion := "Greeting"

	err := dictionary.Add(word, defintion);
	if err != nil{
		fmt.Println(err)
	}
	hello , err2 := dictionary.Search(word)
	fmt.Println(hello)
	if err2 != nil {
		fmt.Println(err2)
	}
}
