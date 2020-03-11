package main

import (
	"fmt"
	"strings"
)

func lenAandUpper(name string) (int, string){
	return len(name), strings.ToUpper(name)
}

func main() {
	totalLength, upperName := lenAandUpper("KimDongEon")
	fmt.Println(totalLength, upperName);
}