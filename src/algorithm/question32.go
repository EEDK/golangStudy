package algorithm

import "fmt"

func Question32(){
	for i := 'A' ; i <= 'Z' ; i++{
		for j := 'A' ; j <= 'Z' - (i - 65) ; j++{
			fmt.Printf("%c " , j)
		}
		fmt.Printf("\n")
	}
}