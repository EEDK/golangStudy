package main

import (
	"fmt"
	"strings"
)

func lenAndUpper(name string) (length int , uppercase string) {
	length = len(name)
	uppercase = strings.ToUpper(name)

	return
}

func main(){
	length , _ := lenAndUpper("kimdongeon")
	fmt.Println(length)
}

