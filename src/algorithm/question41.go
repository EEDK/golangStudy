package algorithm


var table [100]int

func Question41(n int) {
	if n == 0 || n == 1{

	} else {
		table[0] = 0
		table[1] = 1
	}

	for i := 2 ; i < n+1 ; i++{
		table[i] = table[i-1] + table[i-2]
	}
}
