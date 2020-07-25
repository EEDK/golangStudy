package main

import "fmt"

type person struct{
	name string
	age int
	job string
	hobbies []string
}

func main() {
	hobbies := []string{"reading book" , "go learn"}
	kim := person{
		"kimdongeon", 25, "programmer", hobbies,
	}
	fmt.Println(kim.name)
}
