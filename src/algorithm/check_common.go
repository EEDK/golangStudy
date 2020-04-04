package algorithm

import "fmt"


func Check_common(n int){
	for i := 1 ; i <= n ; i++ {
		if (i % 3 == 0) && (i % 5 == 0) {
			fmt.Println(i)
		}
	}
}

