package main

import (
	"fmt"
	"strings"
)

func lenAndUpper(name string) (int , string) {
	return len(name), strings.ToUpper(name)
}

func main(){
	totalLength , _ := lenAndUpper("king dong eon")
	fmt.Println(totalLength )
}

