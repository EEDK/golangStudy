package algorithm

func CalcLCM(n int, m int) int {
	length := n * m
	for i := 2; i <= length; i++ {
		if i%n == 0 && i%m == 0 {
			return i
		}
	}
	return -1
}
