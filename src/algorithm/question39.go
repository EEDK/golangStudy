package algorithm

func Question39(n int) int {
	a := 1
	b := 1

	for i := 3; i < n; i++ {
		temp := b
		b = a + b
		a = temp
	}
	return b
}

func RecurssionQuestion40(n int) int{
	if n <= 2{
		return 1
	}	else {
		return RecurssionQuestion40(n-1) + RecurssionQuestion40(n-2)
	}
}