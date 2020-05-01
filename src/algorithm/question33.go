package algorithm

import "fmt"

func Question33(n int32){
	if n < 'A' || n > 'Z' {

	} else {
		for c := 'A'; c <= n ; c++{
			fmt.Printf("%2c " , c)
		}
		fmt.Printf("\n")
	}
	next := n - 1
	Question33(next)
}
