package algorithm

import "fmt"

var N int = 12

func Question37(){
	for n := 0 ; n <= N ; n++{
		for r := 0 ; r <= n; r++{
			fmt.Printf("%d ", Qustion19(n , r))
		}
		fmt.Printf("\n")
	}
}
