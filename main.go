package main

import "fmt"

func main() {
	names := [5]string{"kim", "dong","eon"}
	names[4] = "one"
	names[3] = "good"

	fmt.Println(names)
}
