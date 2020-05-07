package algorithm

func Question39(n int) int{
	a := 1
	b := 1


	for i := 3; i < n; i++{
		temp := b
		b = a + b
		a = temp
	}
	return b
}