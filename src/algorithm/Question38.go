package algorithm

func Question38(a , b int) int{
	if b == 0 {
		return a
	} else {
		return Question38(b, a % b)
	}
}
