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

func superAdd(numbers ...int) int {
	total := 0

	for i := 0 ; i <= len(numbers) ; i += 1{
		total += i
	}
	return total
}

func canIDrink(age int) bool {
	if  koreanAge := age + 2 ; koreanAge < 18{
		return false
	}
	return true
}


func main() {
	fmt.Println(canIDrink(16))
}