package algorithm

import "fmt"

func Recursion_print_to_n(n int) {
	if n > 0 {
		fmt.Println(n)
		Recursion_print_to_n(n-1)
	}
}