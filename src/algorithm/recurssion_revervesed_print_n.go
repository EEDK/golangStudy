package algorithm

import "fmt"

func Recursion_reverse_print_to_n(n int) {
	fmt.Println(n)

	if n > 1 {
		Recursion_reverse_print_to_n(n-1)
	}
}
