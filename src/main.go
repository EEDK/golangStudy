package main

import (
	"fmt"
	"strings"
)

func lenAndUpper(name string)(length int, uppercase string){
	defer fmt.Println("i'm done")
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}

func main() {
	totalLength, up := lenAndUpper("nico")
	fmt.Println(totalLength , up)
}