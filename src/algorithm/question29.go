package algorithm

import (
	"fmt"
)

func Question29(n int , a string , b string , c string){
	if n > 0 {
		Question29(n-1, a , c , b)
		fmt.Printf("%d번 원판을 %s에서 %s 로 옮김\n", n , a , b)
		Question29(n-1, c , b , a)
	}
}