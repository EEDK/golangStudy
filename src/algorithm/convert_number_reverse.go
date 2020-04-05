package algorithm

import "fmt"

func Convert_number_reverse(n int) {
	 if n != 0 {
		 fmt.Print(n % 10)
		 Convert_number_reverse( n / 10 )
	}
}