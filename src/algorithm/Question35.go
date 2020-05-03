package algorithm

import (
	"fmt"
	"unicode/utf8"
)

func Question35(n string){
	length := utf8.RuneCountInString(n) // 2: 문자열의 실제 길이를 구함

	isRight := true
	for i := 0 ; i < length / 2 ; i++{
		if n[i] != n[length -i - 1]{
			isRight = false
		}
	}
	if isRight{
		fmt.Print("it is palindrome")
	} else {
		fmt.Print("it's not palindrome")
	}
}