package algorithm


func Question31(n int) int{
	answer := -1
	for i := 0 ; n >= 0 ; i++ {
		n = n - 2 * i + 1

		if n < 0 {
			answer = i-1
		}
	}
	return answer
}