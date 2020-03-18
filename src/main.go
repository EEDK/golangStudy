package main

import (
	"fmt"
	"strings"
)

type person struct {
	name string
	age int
	favFood []string
}

func main() {
	favFood := []string{"비빔밥" , "김치"}
	dongeon := person{name : "dongeon" , age : 25 , favFood: favFood};
	fmt.Println(dongeon)
}


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
	switch koreanAge := age + 2; koreanAge {
	case 10 :
		return false
	case 18 :
		return true
	}
	return false
}
