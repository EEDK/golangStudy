package algorithm

func Question30(a int , b int) int{
	if a == b {
		return 0
	} else if b > a {
		return Question30(a, b / 2) + 1
	} else {
		return Question30(a / 2 , b) + 1
	}
}