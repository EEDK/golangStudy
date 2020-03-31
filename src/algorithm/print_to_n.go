package algorithm

func Print_to_n(n int) int {
	sum := 0
	for i := 0 ; i <= n ; i ++ {
		sum += i
	}

	return sum
}