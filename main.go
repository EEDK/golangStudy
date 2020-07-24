package main

import "fmt"

func main() {
	names := []string{"kim", "dong","eon"}
	names = append(names, "flynn")

	fmt.Println(names)
}
