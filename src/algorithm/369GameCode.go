package algorithm

import "fmt"

func ThreeSixNine(n int){
	for i := 1 ; i <= n ; i ++{
		if i % 3 == 0 {
			fmt.Print("x ")
		} else {
			fmt.Printf("%d " , i)
		}
	}
}
