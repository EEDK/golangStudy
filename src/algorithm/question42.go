package algorithm


var table [101]int
var n = 0

func Question42(k int) int{
	if k == n + 1{
		return 1
	}
	table[k] = k + Question42(k-1)
	return table[k]
}
